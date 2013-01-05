package room

import (
	"sync"
)

const (
	North = iota
	South
	East
	West
	Dennis
	directions
)

var (
	lock    sync.Mutex
	rooms   = make(map[Coord]*Room)
	current Coord
)

type Room struct {
	Desc   string
	Pass   [directions]bool
	Direct [directions]string
	Object map[string]*Object
	mtx    sync.Mutex
}

func (r *Room) Release() {
	if r != nil {
		r.mtx.Unlock()
	}
}

func Current() *Room {
	lock.Lock()
	r := rooms[current]
	lock.Unlock()

	if r != nil {
		r.mtx.Lock()
	}
	return r
}
