package main

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// Equality ////////////////////////////////////////////////////////////////////

func eqLobby(lhs Lobby, rhs Lobby) bool {
	return lhs.GID == rhs.GID && lhs.Leader == rhs.Leader
}

func eqCIDSlice(lhs []CID, rhs []CID) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}
	if len(lhs) != len(rhs) {
		return false
	}
	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func eqLobbySlice(lhs []Lobby, rhs []Lobby) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}
	if len(lhs) != len(rhs) {
		return false
	}
	for i := 0; i < len(lhs); i++ {
		if eqLobby(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

// ignore isSpy
func eqCurrentGame(lhs CurrentGame, rhs CurrentGame) bool {
	return lhs.ResPts == rhs.ResPts &&
		lhs.SpyPts == rhs.SpyPts &&
		lhs.Captain == rhs.Captain &&
		lhs.NumMembers == rhs.NumMembers &&
		eqCIDSlice(lhs.Members, rhs.Members) &&
		lhs.Active == rhs.Active
}

func eqEndGame(lhs EndGame, rhs EndGame) bool {
	return lhs.ResPts == rhs.ResPts &&
		lhs.SpyPts == rhs.SpyPts &&
		eqLobbySlice(lhs.Lobbies, rhs.Lobbies)
}

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
		t.Fatal("Members and NumMembers differ", len(y.Members), y.NumMembers)
	}
	if y.Active && y.Members == nil {
		t.Fatal("Members nil when active")
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

// Getters /////////////////////////////////////////////////////////////////////

func getCurrentGame(t *testing.T, cs []*testClient) CurrentGame {
	lhs := cs[0].recvCurrentGame(t)
	for i := 1; i < len(cs); i++ {
		rhs := cs[i].recvCurrentGame(t)
		if !eqCurrentGame(lhs, rhs) {
			t.Fatal("bad CurrentGame", lhs, rhs)
		}
	}
	return lhs
}

func getEndGame(t *testing.T, cs []*testClient) EndGame {
	lhs := cs[0].recvEndGame(t)
	for i := 1; i < len(cs); i++ {
		rhs := cs[i].recvEndGame(t)
		if !eqEndGame(lhs, rhs) {
			t.Fatal("bad EndGame", lhs, rhs)
		}
	}
	return lhs
}

// Extra methods on server /////////////////////////////////////////////////////

func (s *Server) addClient(t *testing.T) *testClient {
	tc := newTestClient()
	tc.send(Connect{})
	s.C <- tc.Client
	tc.CID = tc.recvSetMe(t).Me
	return tc
}

// Helpers /////////////////////////////////////////////////////////////////////

func mkMembers(cs []*testClient, n int) []CID {
	membersMap := make(map[CID]bool)
	for i := n; i > 0; /*  */ {
		cid := cs[rand.Intn(len(cs))].CID
		if membersMap[cid] {
			continue
		}
		membersMap[cid] = true
		i--
	}
	members := make([]CID, 0, n)
	for cid, add := range membersMap {
		if add {
			members = append(members, cid)
		}
	}
	return members
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

func TestGame(t *testing.T) {
	const n = 5
	s := NewServer()
	defer s.Close()
	cs := make([]*testClient, n)
	toIdx := make(map[CID]int)
	for i := 0; i < n; i++ {
		cs[i] = s.addClient(t)
		cs[i].send(NameChoose{string(i)})
		cs[i].recvLobbyChoices(t, 0)
		toIdx[cs[i].CID] = i
	}
	cs[0].send(LobbyCreate{})
	cs[0].recvCurrentLobby(t, 1)
	for i := 1; i < n; i++ {
		lb := cs[i].recvLobbyChoices(t, 1).Lobbies[0]
		cs[i].send(LobbyChoose{lb.GID})
		for j := 0; j <= i; j++ {
			cs[j].recvCurrentLobby(t, i+1)
		}
	}
	cs[0].send(GameStart{})
	var cg CurrentGame
	for i := 0; i < MaxPts; i++ {
		cg = getCurrentGame(t, cs)
		if cg.SpyPts != 0 {
			t.Fatal("bad SpyPts", cg.SpyPts)
		}
		if cg.ResPts != i {
			t.Fatal("bad ResPts", cg.ResPts)
		}
		ms := mkMembers(cs, cg.NumMembers)
		cs[toIdx[cg.Captain]].send(MemberChoose{ms})
		cg = getCurrentGame(t, cs)
		if !eqCIDSlice(ms, cg.Members) {
			t.Fatal("bad members", ms, cg.Members)
		}
		for _, cl := range cs {
			cl.send(MemberVote{true})
		}
		cg = getCurrentGame(t, cs)
		for _, cid := range ms {
			cs[toIdx[cid]].send(MissionVote{true})
		}
	}
	eg := getEndGame(t, cs)
	if !eqEndGame(eg, EndGame{ResPts: 3, SpyPts: 0, Lobbies: []Lobby{}}) {
		t.Fatal("bad EndGame", eg)
	}
}
