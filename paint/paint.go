package paint

import (
	"image"
	"image/color"
	"image/draw"
)

var (
	frame      = 0
	background *image.RGBA
	scanline   = image.NewUniform(color.NRGBA{0, 255, 0, 64})
)

func paint(dst *image.RGBA) {
	if background == nil || background.Bounds() != dst.Bounds() {
		background = image.NewRGBA(dst.Bounds())
	}

	paintText(background, con)

	fade(background)

	bounds := background.Bounds()
	bounds.Min.Y = bounds.Max.Y * (frame % (FramesPerSecond * 3)) / (FramesPerSecond * 2)
	bounds.Max.Y = bounds.Min.Y + (bounds.Max.Y / FramesPerSecond / 4)
	draw.Draw(background, bounds, scanline, image.ZP, draw.Over)

	bounds = background.Bounds()
	bounds.Min.Y = bounds.Max.Y * (frame % (FramesPerSecond * 5)) / (FramesPerSecond * 2)
	bounds.Max.Y = bounds.Min.Y + (bounds.Max.Y / FramesPerSecond / 4)
	draw.Draw(background, bounds, scanline, image.ZP, draw.Over)

	draw.Draw(dst, dst.Bounds(), background, image.ZP, draw.Src)
	paintClearText(dst)

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
