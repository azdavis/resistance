package main

// Only one goroutine may invoke add/rm at a time.
type clientMap struct {
	c     chan IDAction
	m     map[ID]*client
	quits map[ID]chan struct{}
}

func newClientMap() *clientMap {
	cm := &clientMap{
		c:     make(chan IDAction),
		m:     make(map[ID]*client),
		quits: make(map[ID]chan struct{}),
	}
	return cm
}

func (cm *clientMap) add(cl *client) {
	_, ok := cm.quits[cl.id]
	if ok {
		panic("already present")
	}
	quit := make(chan struct{})
	cm.m[cl.id] = cl
	cm.quits[cl.id] = quit
	go cm.pipe(cl.id, cl.recv, quit)
}

func (cm *clientMap) rm(id ID) {
	_, ok := cm.quits[id]
	if !ok {
		panic("not present")
	}
	close(cm.quits[id])
	delete(cm.m, id)
	delete(cm.quits, id)
}

func (cm *clientMap) pipe(id ID, ch chan Action, quit chan struct{}) {
	for {
		select {
		case <-quit:
			return
		case ac := <-ch:
			cm.c <- IDAction{id, ac}
		}
	}
}
