package paint

import (
	"image"
	"image/color"
	"image/draw"
)

var (
	frame    = 0
	fade     = image.NewUniform(color.NRGBA{0, 0, 0, 8})
	scanline = image.NewUniform(color.NRGBA{0, 255, 0, 64})
)

func paint(dst *image.RGBA) {
	bounds := dst.Bounds()
	draw.Draw(dst, bounds, fade, image.ZP, draw.Over)

	bounds.Min.Y = bounds.Max.Y * (frame % (FramesPerSecond * 2)) / FramesPerSecond
	bounds.Max.Y = bounds.Min.Y + (bounds.Max.Y / FramesPerSecond / 2)
	draw.Draw(dst, bounds, scanline, image.ZP, draw.Over)

	frame++
}
