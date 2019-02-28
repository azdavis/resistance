package main

import (
	"log"
)

// Only one goroutine may call Add or Rm or access M at a time.
type ClientMap struct {
	C     chan IDAction
	M     map[ID]*Client
	quits map[ID]chan struct{}
}

func NewClientMap() *ClientMap {
	cm := &ClientMap{
		C:     make(chan IDAction),
		M:     make(map[ID]*Client),
		quits: make(map[ID]chan struct{}),
	}
	return cm
}

func (cm *ClientMap) Add(cl *Client) {
	_, ok := cm.quits[cl.id]
	if ok {
		panic("already present")
	}
	log.Println("Add", cl.id)
	quit := make(chan struct{})
	cm.M[cl.id] = cl
	cm.quits[cl.id] = quit
	go cm.pipe(cl.id, cl.recv, quit)
}

func (cm *ClientMap) Rm(id ID) {
	_, ok := cm.quits[id]
	if !ok {
		panic("not present")
	}
	log.Println("Rm", id)
	close(cm.M[id].send)
	close(cm.quits[id])
	delete(cm.M, id)
	delete(cm.quits, id)
}

func (cm *ClientMap) pipe(id ID, ch chan Action, quit chan struct{}) {
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
