package main

import (
	"github.com/Nightgunner5/g/room"
)

func init() {
	commands["look"] = func(argv []string) {
		r := room.Current()
		defer r.Release()

		r.Look(argv)
	}

	help["look"] = `look v8.27: usage:
look
look <object>
look <direction>`
}
