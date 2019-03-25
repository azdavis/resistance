package main

import (
	"runtime"
	"testing"
	"time"
)

func TestBasicNumGoroutine(t *testing.T) {
	before := runtime.NumGoroutine()
	s := NewServer()
	during := runtime.NumGoroutine()
	s.Close()
	time.Sleep(time.Second)
	after := runtime.NumGoroutine()
	if before > during {
		t.Fatal("before > during")
	}
	if before != after {
		t.Fatal("before != after")
	}
}
