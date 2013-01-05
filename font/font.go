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
	dpi   = 72
	size  = 16
	inset = raster.Fix32(8 << 8)
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
	c.SetClip(dst.Bounds())
	c.SetDst(dst)
	c.SetSrc(image.Opaque)
	return c
}

func Start(c *freetype.Context, height int) raster.Point {
	return raster.Point{inset, raster.Fix32(height<<8) - inset}
}

func NextLine(c *freetype.Context, pt raster.Point) raster.Point {
	return raster.Point{inset, pt.Y - c.PointToFix32(size*1.5)}
}
