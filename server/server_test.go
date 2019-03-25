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

func checkClients(t *testing.T, xs []*Client) {
	if xs == nil {
		t.Fatal("Clients was nil")
	}
	for _, x := range xs {
		if x == nil {
			t.Fatal("client was nil")
		}
		checkCID(t, x.CID)
	}
}

func checkLobbies(t *testing.T, xs []Lobby) {
	if xs == nil {
		t.Fatal("Lobbies was nil")
	}
	for _, x := range xs {
		checkGID(t, x.GID)
	}
}

// testClient //////////////////////////////////////////////////////////////////

type testClient struct {
	*Client
	req chan struct{}
	res chan ToClient
}

func newTestClient() *testClient {
	tc := &testClient{NewClient(nil), make(chan struct{}), make(chan ToClient)}
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

func (tc *testClient) recvLobbyChoices(t *testing.T) LobbyChoices {
	x := tc.recv()
	y, ok := x.(LobbyChoices)
	if !ok {
		t.Fatal("response was not LobbyChoices")
	}
	checkLobbies(t, y.Lobbies)
	return y
}

func (tc *testClient) recvCurrentLobby(t *testing.T) CurrentLobby {
	x := tc.recv()
	y, ok := x.(CurrentLobby)
	if !ok {
		t.Fatal("response was not CurrentLobby")
	}
	checkGID(t, y.GID)
	checkCID(t, y.Leader)
	checkClients(t, y.Clients)
	return y
}

func (tc *testClient) recvCurrentGame(t *testing.T) CurrentGame {
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
		t.Fatal("number of members differ", len(y.Members), y.NumMembers)
	}
	if y.Active && y.Members == nil {
		t.Fatal("Members nil when active")
	}
	for _, x := range y.Members {
		checkCID(t, x)
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
	sm := tc.recvSetMe(t)
	if sm.Me != tc.CID {
		t.Fatal("SetMe CIDs differ", sm.Me, tc.CID)
	}
	return tc
}

// Tests ///////////////////////////////////////////////////////////////////////

func TestNumGoroutine(t *testing.T) {
	pre := runtime.NumGoroutine()
	s := NewServer()
	now := runtime.NumGoroutine()
	if now != pre+2 {
		t.Fatal("bad number of goroutines", pre, now)
	}
	for i := 1; i < 10; i++ {
		s.addClient(t)
		now = runtime.NumGoroutine()
		if now != pre+2+(i*2) {
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
	lc := c.recvLobbyChoices(t)
	if len(lc.Lobbies) != 0 {
		t.Fatal("lobbies not len 0")
	}
}

func TestTwoClients(t *testing.T) {
	s := NewServer()
	defer s.Close()
	c1 := s.addClient(t)
	c2 := s.addClient(t)
	c1.send(NameChoose{"c1"})
	c2.send(NameChoose{"c2"})
	c1.recvLobbyChoices(t)
	c2.recvLobbyChoices(t)
	c1.send(LobbyCreate{})
	cl := c1.recvCurrentLobby(t)
	if c1.CID != cl.Leader {
		t.Fatal("leader differ", c1.CID, cl.Leader)
	}
	if len(cl.Clients) != 1 {
		t.Fatal("Clients not len 1")
	}
	lc := c2.recvLobbyChoices(t)
	if len(lc.Lobbies) != 1 {
		t.Fatal("Lobbies not len 1")
	}
	l1 := lc.Lobbies[0]
	if c1.Name != l1.Leader {
		t.Fatal("leader differ", c1.Name, l1.Leader)
	}
	c2.send(LobbyChoose{l1.GID})
	cl = c1.recvCurrentLobby(t)
	if len(cl.Clients) != 2 {
		t.Fatal("Clients not len 2")
	}
	cl = c2.recvCurrentLobby(t)
	if len(cl.Clients) != 2 {
		t.Fatal("Clients not len 2")
	}
}
