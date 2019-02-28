package main

type bigChan struct {
	c     chan IDAction
	quits map[ID]chan struct{}
}

func newBigChan() *bigChan {
	bc := &bigChan{
		c:     make(chan IDAction),
		quits: make(map[ID]chan struct{}),
	}
	return bc
}

func (bc *bigChan) add(id ID, ch chan Action) {
	_, ok := bc.quits[id]
	if ok {
		panic("already present")
	}
	quit := make(chan struct{})
	bc.quits[id] = quit
	go bc.pipe(id, ch, quit)
}

func (bc *bigChan) rm(id ID) {
	_, ok := bc.quits[id]
	if !ok {
		panic("not present")
	}
	close(bc.quits[id])
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
