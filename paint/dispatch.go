package paint

import (
	"image"
	"image/draw"
	"time"
)

const FramesPerSecond = 30

type signal struct {
	tick  bool
	reset *image.Rectangle
}

var (
	flush    func(*image.RGBA)
	viewport *image.RGBA
	comm     chan signal
)

func Start(f func(*image.RGBA)) {
	flush = f
	comm = make(chan signal)
	go dispatch()
}

func dispatch() {
	go tick()

	for s := range comm {
		switch {
		case s.tick:
			if viewport != nil {
				paint(viewport)
				flush(viewport)
			}

		case s.reset != nil:
			viewport = image.NewRGBA(*s.reset)
			draw.Draw(viewport, viewport.Bounds(), image.Black, image.ZP, draw.Src)
		}
	}
}

func tick() {
	defer func() {
		// don't crash if the comm channel gets closed
		recover()
	}()
	for {
		time.Sleep(time.Second / FramesPerSecond)
		comm <- signal{
			tick: true,
		}
	}
}

func ResetViewport(bounds image.Rectangle) {
	comm <- signal{
		reset: &bounds,
	}
}
