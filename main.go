package main

import (
	"github.com/Nightgunner5/g/console"
	"github.com/Nightgunner5/g/paint"
	"github.com/skelterjohn/go.wde"
	"image"
	"strings"
)

var cleanup []func()

var commands = make(map[string]func([]string))

func main() {
	go ui()

	wde.Run()

	for _, f := range cleanup {
		f()
	}
}

func ui() {
	defer wde.Stop()

	w, err := wde.NewWindow(640, 480)
	handle(err)

	w.SetTitle("g")

	w.Show()

	paint.Start(func(img *image.RGBA) {
		screen := w.Screen()
		screen.CopyRGBA(img, screen.Bounds())
		w.FlushImage(screen.Bounds())
	}, func(argv []string) {
		if len(argv) == 0 {
			return
		}
		paint.WithConsole(func(c *console.Console) {
			c.Bprint(">")
			c.Println(strings.Join(argv, " "))
		})
		if cmd, ok := commands[argv[0]]; ok {
			go cmd(argv[1:])
		} else {
			paint.WithConsole(func(c *console.Console) {
				c.Errorln("Unknown command or filename.")
				c.Print("Type")
				c.Bprint("help")
				c.Println("at the prompt for a listing of basic commands.")
			})
		}
	})
	paint.ResetViewport(w.Screen().Bounds())

	for event := range w.EventChan() {
		switch e := event.(type) {
		case wde.CloseEvent:
			_ = e
			return

		case wde.KeyTypedEvent:
			paint.Typed(e.Key, e.Glyph)

		case wde.ResizeEvent:
			paint.ResetViewport(w.Screen().Bounds())
		}
	}
}
