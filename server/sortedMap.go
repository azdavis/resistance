package main

import (
	"sort"
)

// Entry is an entry in a SortedMap.
type Entry struct {
	K uint64
	V string
}

// SortedMap is a sorted map of Entries. M will always be sorted with respect to
// the K of the entries, in increasing order. Only one goroutine at a time may
// operate on a SortedMap.
type SortedMap struct {
	M []Entry
}

// NewSortedMap returns a new SortedMap with the given cap.
func NewSortedMap(cap int) *SortedMap {
	sm := &SortedMap{M: make([]Entry, 0, cap)}
	return sm
}

// Add and Rm impl: https://github.com/golang/go/wiki/SliceTricks

// Add adds the (k, v) pair to the map. k must not already be in the map.
func (sm *SortedMap) Add(k uint64, v string) {
	newEntry := Entry{K: k, V: v}
	i := sort.Search(len(sm.M), func(i int) bool { return sm.M[i].K > k })
	sm.M = append(sm.M, newEntry)
	copy(sm.M[i+1:], sm.M[i:])
	sm.M[i] = newEntry
}

// Rm removes the given k and its corresponding v from the map, preserving
// sortedness. k must already be in the map.
func (sm *SortedMap) Rm(k uint64) {
	i := sort.Search(len(sm.M), func(i int) bool { return sm.M[i].K == k })
	copy(sm.M[i:], sm.M[i+1:])
	sm.M[len(sm.M)-1] = Entry{}
	sm.M = sm.M[:len(sm.M)-1]
}
