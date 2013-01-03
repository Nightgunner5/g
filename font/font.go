package font

import (
	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/raster"
	"code.google.com/p/freetype-go/freetype/truetype"
	"image"
)

var (
	reg = must(freetype.ParseFont(MondaRegularTtf))
	bld = must(freetype.ParseFont(MondaBoldTtf))
)

func must(font *truetype.Font, err error) *truetype.Font {
	if err != nil {
		panic(err)
	}
	return font
}

const (
	dpi  = 72
	size = 16
	left = raster.Fix32(8 << 8)
	top  = raster.Fix32(8 << 8)
)

func Context(dst *image.Alpha, bold bool) *freetype.Context {
	c := freetype.NewContext()
	c.SetDPI(dpi)
	if bold {
		c.SetFont(bld)
	} else {
		c.SetFont(reg)
	}
	c.SetFontSize(size)
	c.SetClip(dst.Bounds().Inset(8))
	c.SetDst(dst)
	c.SetSrc(image.Opaque)
	return c
}

func Start(c *freetype.Context) raster.Point {
	return raster.Point{left, top + c.PointToFix32(size)}
}

func NextLine(c *freetype.Context, pt raster.Point) raster.Point {
	return raster.Point{left, pt.Y + c.PointToFix32(size*1.5)}
}
