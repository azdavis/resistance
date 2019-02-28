package main

import (
	"log"
)

// Only one goroutine may call add or rm or access M at a time.
type clientMap struct {
	C     chan IDAction
	M     map[ID]*client
	quits map[ID]chan struct{}
}

func newClientMap() *clientMap {
	cm := &clientMap{
		C:     make(chan IDAction),
		M:     make(map[ID]*client),
		quits: make(map[ID]chan struct{}),
	}
	return cm
}

func (cm *clientMap) add(cl *client) {
	_, ok := cm.quits[cl.id]
	if ok {
		panic("already present")
	}
	log.Println("add", cl.id)
	quit := make(chan struct{})
	cm.M[cl.id] = cl
	cm.quits[cl.id] = quit
	go cm.pipe(cl.id, cl.recv, quit)
}

func (cm *clientMap) rm(id ID) {
	_, ok := cm.quits[id]
	if !ok {
		panic("not present")
	}
	log.Println("rm", id)
	close(cm.M[id].send)
	close(cm.quits[id])
	delete(cm.M, id)
	delete(cm.quits, id)
}

func (cm *clientMap) pipe(id ID, ch chan Action, quit chan struct{}) {
	defer log.Println("exit pipe", id)
	for {
		select {
		case <-quit:
			return
		case ac := <-ch:
			cm.C <- IDAction{id, ac}
		}
	}
}
