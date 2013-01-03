package paint

import (
	"code.google.com/p/freetype-go/freetype"
	"github.com/Nightgunner5/g/font"
	"image"
	"image/color"
	"image/draw"
)

var (
	textColor     = image.NewUniform(color.NRGBA{0, 255, 0, 255})
	textBuf       *image.Alpha
	textBufPrompt *image.Alpha
	textCtxReg    *freetype.Context
	textCtxBld    *freetype.Context
	textDirty     = true
)

func paintText(dst *image.RGBA) {
	if textBuf == nil || textBuf.Bounds() != dst.Bounds() {
		textBuf = image.NewAlpha(dst.Bounds())
		textBufPrompt = image.NewAlpha(dst.Bounds())
		textCtxReg = font.Context(textBuf, false)
		textCtxBld = font.Context(textBuf, true)
		textDirty = true
	}

	if textDirty {
		draw.Draw(textBuf, textBuf.Bounds(), image.Transparent, image.ZP, draw.Src)
		draw.Draw(textBufPrompt, textBufPrompt.Bounds(), image.Transparent, image.ZP, draw.Src)
		pt := font.Start(textCtxReg)
		pt, _ = textCtxReg.DrawString("> g", pt)

		c := font.Context(textBufPrompt, false)
		c.DrawString("_", pt)

		textDirty = false
	}

	draw.DrawMask(dst, dst.Bounds(), textColor, image.ZP, textBuf, image.ZP, draw.Over)
	if (frame/FramesPerSecond)%4 == 0 {
		draw.DrawMask(dst, dst.Bounds(), textColor, image.ZP, textBufPrompt, image.ZP, draw.Over)
	}
}
