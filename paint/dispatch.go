package paint

import (
	"github.com/Nightgunner5/g/console"
	"image"
	"image/draw"
	"time"
)

const FramesPerSecond = 30

type signal struct {
	tick  bool
	reset *image.Rectangle
	typed *[2]string
}

var (
	flush    func(*image.RGBA)
	viewport *image.RGBA
	comm     chan signal
	con      = console.New()
)

func Start(f func(*image.RGBA), e func([]string)) {
	flush = f
	con.Exec = e
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

		case s.typed != nil:
			con.Typed((*s.typed)[0], (*s.typed)[1])

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

func Typed(key, glyph string) {
	var keyglyph = [2]string{key, glyph}
	comm <- signal{
		typed: &keyglyph,
	}
}
