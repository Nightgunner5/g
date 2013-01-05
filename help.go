package main

import (
	"github.com/Nightgunner5/g/console"
	"github.com/Nightgunner5/g/paint"
	"strings"
)

var help = map[string]string{
	"":     "usage: help <command name>",
	"help": "HELP! I'm trapped in a documentation factory!",
}

func init() {
	commands["help"] = func(argv []string) {
		key := strings.Join(argv, " ")
		paint.WithConsole(func(c *console.Console) {
			if h, ok := help[key]; ok {
				c.Println(h)
			} else {
				c.Error("No help for topic")
				c.Berrorln(key)
			}
		})
	}
}
