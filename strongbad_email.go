package main

import (
	"github.com/Nightgunner5/g/console"
	"github.com/Nightgunner5/g/paint"
)

func init() {
	commands["strongbad_email.exe"] = func([]string) {
		paint.WithConsole(func(c *console.Console) {
			c.Println("Dear Stong Bag,")
			c.Println("Do you like crap? Do you like the word crap?")
			c.Println("Sometimes I like to write about crap. Isn't")
			c.Println("crap craptastic? Crap is the crap.")
			c.Println()
			c.Println("Crapfully crapfully,")
			c.Println("Dennis")
		})
	}
}
