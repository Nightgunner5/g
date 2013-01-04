package paint

import (
	"image"
	"image/color"
	"image/draw"
)

var (
	frame    = 0
	scanline = image.NewUniform(color.NRGBA{0, 255, 0, 64})
)

func paint(dst *image.RGBA) {
	paintText(dst, con)

	fade(dst)

	bounds := dst.Bounds()
	bounds.Min.Y = bounds.Max.Y * (frame % (FramesPerSecond * 3)) / (FramesPerSecond * 2)
	bounds.Max.Y = bounds.Min.Y + (bounds.Max.Y / FramesPerSecond / 4)
	draw.Draw(dst, bounds, scanline, image.ZP, draw.Over)

	bounds = dst.Bounds()
	bounds.Min.Y = bounds.Max.Y * (frame % (FramesPerSecond * 5)) / (FramesPerSecond * 2)
	bounds.Max.Y = bounds.Min.Y + (bounds.Max.Y / FramesPerSecond / 4)
	draw.Draw(dst, bounds, scanline, image.ZP, draw.Over)

	frame++
}

// A more efficient way of obtaining the same result as drawing
// color.RGBA{0, 0, 0, 8} over every pixel.
func fade(dst *image.RGBA) {
	const (
		a = 256 - 8
	)

	for i, b := range dst.Pix {
		dst.Pix[i] = uint8((uint(b) * a) >> 8)
	}
}
