package main

type bigChan struct {
	c     chan IDAction
	m     map[ID]*client
	quits map[ID]chan struct{}
}

func newBigChan() *bigChan {
	bc := &bigChan{
		c:     make(chan IDAction),
		quits: make(map[ID]chan struct{}),
	}
	return bc
}

func (bc *bigChan) add(cl *client) {
	_, ok := bc.quits[cl.id]
	if ok {
		panic("already present")
	}
	quit := make(chan struct{})
	bc.m[cl.id] = cl
	bc.quits[cl.id] = quit
	go bc.pipe(cl.id, cl.recv, quit)
}

func (bc *bigChan) rm(id ID) {
	_, ok := bc.quits[id]
	if !ok {
		panic("not present")
	}
	close(bc.quits[id])
	delete(bc.m, id)
	delete(bc.quits, id)
}

func (bc *bigChan) pipe(id ID, ch chan Action, quit chan struct{}) {
	for {
		select {
		case <-quit:
			return
		case ac := <-ch:
			bc.c <- IDAction{id, ac}
		}
	}
}
