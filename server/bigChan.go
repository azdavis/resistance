package main

type addMsg struct {
	id ID
	ch chan Action
}

type bigChan struct {
	c    chan IDAction
	adds chan addMsg
	rms  chan ID
}

func newBigChan() *bigChan {
	bc := &bigChan{
		c:    make(chan IDAction),
		adds: make(chan addMsg),
		rms:  make(chan ID),
	}
	go bc.run()
	return bc
}

func (bc *bigChan) add(id ID, ch chan Action) {
	bc.adds <- addMsg{id, ch}
}

func (bc *bigChan) rm(id ID) {
	bc.rms <- id
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

func (bc *bigChan) run() {
	quits := make(map[ID]chan struct{})
	for {
		select {
		case a := <-bc.adds:
			quit := make(chan struct{})
			quits[a.id] = quit
			go bc.pipe(a.id, a.ch, quit)
		case id := <-bc.rms:
			close(quits[id])
			delete(quits, id)
		}
	}
}
