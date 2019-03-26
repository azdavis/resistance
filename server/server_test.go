package main

import (
	"runtime"
	"testing"
	"time"
)

// Checkers ////////////////////////////////////////////////////////////////////

func checkCID(t *testing.T, cid CID) {
	if cid == 0 {
		t.Fatal("cid was 0")
	}
}

func checkGID(t *testing.T, gid GID) {
	if gid == 0 {
		t.Fatal("gid was 0")
	}
}

func checkClients(t *testing.T, xs []ClientInfo, n int) {
	if xs == nil {
		t.Fatal("Clients was nil")
	}
	if n != len(xs) {
		t.Fatal("bad Clients len", n, len(xs))
	}
	for _, x := range xs {
		checkCID(t, x.CID)
	}
}

func checkLobbies(t *testing.T, xs []Lobby, n int) {
	if xs == nil {
		t.Fatal("Lobbies was nil")
	}
	if n != len(xs) {
		t.Fatal("bad Lobbies len", n, len(xs))
	}
	for _, x := range xs {
		checkGID(t, x.GID)
	}
}

// testClient //////////////////////////////////////////////////////////////////

type testClient struct {
	CID
	Client
	req chan struct{}
	res chan ToClient
}

func newTestClient() *testClient {
	tc := &testClient{
		CID:    0,
		Client: NewClient(nil),
		req:    make(chan struct{}),
		res:    make(chan ToClient),
	}
	go tc.doTestTx()
	return tc
}

func (tc *testClient) doTestTx() {
	// invariant: for all i, next <= i < have, ms[i] exists
	next := uint(0)
	have := uint(0)
	// invariant: for all i, i < want, ms[i] was requested
	want := uint(0)
	ms := make(map[uint]ToClient)
	for {
		select {
		case <-tc.q:
			return
		case m := <-tc.tx:
			ms[have] = m
			have++
		case <-tc.req:
			want++
		}
		for next < have && next < want {
			tc.res <- ms[next]
			delete(ms, next)
			next++
		}
	}
}

func (tc *testClient) send(m ToServer) {
	tc.rx <- m
}

func (tc *testClient) recv() ToClient {
	tc.req <- struct{}{}
	return <-tc.res
}

func (tc *testClient) recvSetMe(t *testing.T) SetMe {
	x := tc.recv()
	y, ok := x.(SetMe)
	if !ok {
		t.Fatal("response was not SetMe")
	}
	checkCID(t, y.Me)
	return y
}

func (tc *testClient) recvNameReject(t *testing.T) NameReject {
	x := tc.recv()
	y, ok := x.(NameReject)
	if !ok {
		t.Fatal("response was not NameReject")
	}
	return y
}

func (tc *testClient) recvLobbyChoices(t *testing.T, n int) LobbyChoices {
	x := tc.recv()
	y, ok := x.(LobbyChoices)
	if !ok {
		t.Fatal("response was not LobbyChoices")
	}
	checkLobbies(t, y.Lobbies, n)
	return y
}

func (tc *testClient) recvCurrentLobby(t *testing.T, n int) CurrentLobby {
	x := tc.recv()
	y, ok := x.(CurrentLobby)
	if !ok {
		t.Fatal("response was not CurrentLobby")
	}
	checkGID(t, y.GID)
	checkCID(t, y.Leader)
	checkClients(t, y.Clients, n)
	return y
}

func (tc *testClient) recvCurrentGame(t *testing.T, n int) CurrentGame {
	x := tc.recv()
	y, ok := x.(CurrentGame)
	if !ok {
		t.Fatal("response was not CurrentGame")
	}
	if !(0 <= y.ResPts && y.ResPts < MaxPts) {
		t.Fatal("ResPts not in range", y.ResPts)
	}
	if !(0 <= y.SpyPts && y.SpyPts < MaxPts) {
		t.Fatal("SpyPts not in range", y.SpyPts)
	}
	if y.Members != nil && len(y.Members) != y.NumMembers {
		t.Fatal("Members and NumMembers differ", len(y.Members), y.NumMembers)
	}
	if y.Active && y.Members == nil {
		t.Fatal("Members nil when active")
	}
	if n != len(y.Members) {
		t.Fatal("bad Members len", n, len(y.Members))
	}
	if y.Members != nil {
		for _, x := range y.Members {
			checkCID(t, x)
		}
	}
	return y
}

func (tc *testClient) recvEndGame(t *testing.T) EndGame {
	x := tc.recv()
	y, ok := x.(EndGame)
	if !ok {
		t.Fatal("response was not EndGame")
	}
	if !(0 <= y.ResPts && y.ResPts <= MaxPts) {
		t.Fatal("ResPts not in range", y.ResPts)
	}
	if !(0 <= y.SpyPts && y.SpyPts <= MaxPts) {
		t.Fatal("SpyPts not in range", y.SpyPts)
	}
	if y.ResPts == MaxPts && y.SpyPts == MaxPts {
		t.Fatal("both ResPts and SpyPts are MaxPts")
	}
	if y.ResPts != MaxPts && y.SpyPts != MaxPts {
		t.Fatal("neither ResPts nor SpyPts are MaxPts")
	}
	return y
}

// Extra methods on server /////////////////////////////////////////////////////

func (s *Server) addClient(t *testing.T) *testClient {
	tc := newTestClient()
	tc.send(Connect{})
	s.C <- tc.Client
	tc.CID = tc.recvSetMe(t).Me
	return tc
}

// Tests ///////////////////////////////////////////////////////////////////////

func TestNumGoroutine(t *testing.T) {
	pre := runtime.NumGoroutine()
	s := NewServer()
	now := runtime.NumGoroutine()
	if now != pre+1 {
		t.Fatal("bad number of goroutines", pre, now)
	}
	for i := 1; i < 10; i++ {
		s.addClient(t)
		now = runtime.NumGoroutine()
		if now != pre+1+(i*2) {
			t.Fatal("bad number of goroutines", pre, now, i)
		}
	}
	s.Close()
	time.Sleep(500 * time.Millisecond)
	now = runtime.NumGoroutine()
	if pre != now {
		t.Fatal("pre != now", pre, now)
	}
}

func TestOneClient(t *testing.T) {
	s := NewServer()
	defer s.Close()
	c := s.addClient(t)
	c.send(NameChoose{""})
	c.recvNameReject(t)
	c.send(NameChoose{"        "})
	c.recvNameReject(t)
	c.send(NameChoose{"   \t\t   \t  "})
	c.recvNameReject(t)
	c.send(NameChoose{"asdfasdfahsfasfhkaslkdjfhasfkajshfaslkfjhadlkfjahsflkas"})
	c.recvNameReject(t)
	c.send(NameChoose{"fella"})
	c.recvLobbyChoices(t, 0)
}

func TestTwoClients(t *testing.T) {
	s := NewServer()
	defer s.Close()
	c1 := s.addClient(t)
	c2 := s.addClient(t)
	c1.send(NameChoose{"c1"})
	c2.send(NameChoose{"c2"})
	c1.recvLobbyChoices(t, 0)
	c2.recvLobbyChoices(t, 0)
	c1.send(LobbyCreate{})
	c1.recvCurrentLobby(t, 1)
	lc := c2.recvLobbyChoices(t, 1)
	l1 := lc.Lobbies[0]
	c2.send(LobbyChoose{l1.GID})
	c1.recvCurrentLobby(t, 2)
	c2.recvCurrentLobby(t, 2)
}
