package main

import (
	"github.com/Nightgunner5/g/paint"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image"
)

func main() {
	go ui()

	wde.Run()
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
	})
	paint.ResetViewport(w.Screen().Bounds())

	for event := range w.EventChan() {
		switch e := event.(type) {
		case wde.CloseEvent:
			_ = e
			return

		case wde.ResizeEvent:
			paint.ResetViewport(w.Screen().Bounds())
		}
	}
}
