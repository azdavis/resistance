package main

import (
	"runtime"
	"testing"
	"time"
)

// testClient //////////////////////////////////////////////////////////////////

type testClient struct {
	*Client
	req chan struct{}
	res chan ToClient
}

func newTestClient() *testClient {
	cl := &Client{
		CID:  0,
		Name: "",
		tx:   make(chan ToClient, ClientChLen),
		rx:   make(chan ToServer, ClientChLen),
		conn: nil,
	}
	tc := &testClient{cl, make(chan struct{}), make(chan ToClient)}
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
	if y.Lobbies == nil {
		t.Fatal("Lobbies was nil")
	}
	return y
}

func (tc *testClient) recvCurrentLobby(t *testing.T) CurrentLobby {
	x := tc.recv()
	y, ok := x.(CurrentLobby)
	if !ok {
		t.Fatal("response was not CurrentLobby")
	}
	if y.Clients == nil {
		t.Fatal("Clients was nil")
	}
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
	return y
}

func (tc *testClient) recvEndGame(t *testing.T) EndGame {
	x := tc.recv()
	y, ok := x.(EndGame)
	if !ok {
		t.Fatal("response was not EndGame")
	}
	if !(0 <= y.ResPts && y.ResPts < MaxPts) {
		t.Fatal("ResPts not in range", y.ResPts)
	}
	if !(0 <= y.SpyPts && y.SpyPts < MaxPts) {
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
	m := tc.recv()
	sm, ok := m.(SetMe)
	if !ok {
		t.Fatal("response was not SetMe")
	}
	if sm.Me != tc.CID {
		t.Fatal("SetMe CIDs differ", sm.Me, tc.CID)
	}
	return tc
}

func (s *Server) closeAndWait() {
	s.Close()
	time.Sleep(500 * time.Millisecond)
}

// Tests ///////////////////////////////////////////////////////////////////////

func TestBasicNumGoroutine(t *testing.T) {
	before := runtime.NumGoroutine()
	s := NewServer()
	during := runtime.NumGoroutine()
	s.closeAndWait()
	after := runtime.NumGoroutine()
	if before > during {
		t.Fatal("before > during")
	}
	if before != after {
		t.Fatal("before != after")
	}
}

func TestOneClient(t *testing.T) {
	s := NewServer()
	defer s.closeAndWait()
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
