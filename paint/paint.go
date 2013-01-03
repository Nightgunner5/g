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
	paintText(dst)

	fade(dst)

	bounds := dst.Bounds()
	bounds.Min.Y = bounds.Max.Y * (frame % (FramesPerSecond * 2)) / FramesPerSecond
	bounds.Max.Y = bounds.Min.Y + (bounds.Max.Y / FramesPerSecond / 2)
	draw.Draw(dst, bounds, scanline, image.ZP, draw.Over)

	frame++
}

var fadeMap [256]uint8

func init() {
	const (
		m = 255
		a = m - 8
	)

	for i := range fadeMap {
		fadeMap[i] = uint8(uint16(i) * a / m)
	}
}

// A more efficient way of obtaining the same result as drawing
// color.RGBA{0, 0, 0, 8} over every pixel.
func fade(dst *image.RGBA) {
	i0 := dst.PixOffset(0, 0)
	i1 := i0 + dst.Bounds().Dx()*4

	dy := dst.Bounds().Dy()

	for ; dy > 0; dy-- {
		for i := i0; i < i1; i += 4 {
			dst.Pix[i+0] = fadeMap[dst.Pix[i+0]]
			dst.Pix[i+1] = fadeMap[dst.Pix[i+1]]
			dst.Pix[i+2] = fadeMap[dst.Pix[i+2]]
		}

		i0 += dst.Stride
		i1 += dst.Stride
	}
}
