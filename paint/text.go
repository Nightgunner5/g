package paint

import (
	"code.google.com/p/freetype-go/freetype/raster"
	"github.com/Nightgunner5/g/console"
	"github.com/Nightgunner5/g/font"
	"image"
	"image/color"
	"image/draw"
)

var (
	textColorGreen   = image.NewUniform(color.NRGBA{64, 225, 64, 255})
	textColorRed     = image.NewUniform(color.NRGBA{225, 64, 64, 255})
	textColorCursor  = image.NewUniform(color.NRGBA{64, 225, 64, 128})
	textBufGreen     *image.Alpha
	textBufRed       *image.Alpha
	textBufInput     *image.Alpha
	textBufCursor    *image.Alpha
	textInputDirty   = true
	textDirty        = true
	textInputVersion uint64
	textVersion      uint64
	textInputStart   raster.Point
)

func paintText(dst *image.RGBA, con *console.Console) {
	if textBufCursor == nil || textBufCursor.Bounds() != dst.Bounds() {
		textBufGreen = image.NewAlpha(dst.Bounds())
		textBufRed = image.NewAlpha(dst.Bounds())
		textBufInput = image.NewAlpha(dst.Bounds())
		textBufCursor = image.NewAlpha(dst.Bounds())
		textDirty = true
	}

	if textVersion != con.Version || textInputVersion != con.InputVersion {
		if textVersion != con.Version {
			textVersion = con.Version
			textDirty = true
		}
		textInputVersion = con.InputVersion
		textInputDirty = true
	}

	if textInputDirty {
		draw.Draw(textBufInput, textBufInput.Bounds(), image.Transparent, image.ZP, draw.Src)
		cInput := font.Context(textBufInput, false)
		if textDirty {
			draw.Draw(textBufGreen, textBufGreen.Bounds(), image.Transparent, image.ZP, draw.Src)
			draw.Draw(textBufRed, textBufRed.Bounds(), image.Transparent, image.ZP, draw.Src)
			cGreen := font.Context(textBufGreen, false)
			cGreenBold := font.Context(textBufGreen, true)
			cRed := font.Context(textBufRed, false)
			cRedBold := font.Context(textBufRed, true)

			pt := font.Start(cGreen)
			for _, word := range con.Output {
				if word.IsNewline() {
					pt = font.NextLine(cGreen, pt)
				} else {
					if word.Red {
						if word.Bold {
							pt, _ = cRedBold.DrawString(word.Text, pt)
						} else {
							pt, _ = cRed.DrawString(word.Text, pt)
						}
					} else {
						if word.Bold {
							pt, _ = cGreenBold.DrawString(word.Text, pt)
						} else {
							pt, _ = cGreen.DrawString(word.Text, pt)
						}
					}
					pt, _ = cGreen.DrawString(" ", pt)
				}
			}

			for _, word := range con.Prompt {
				if word.IsNewline() {
					pt = font.NextLine(cGreen, pt)
				} else {
					if word.Red {
						if word.Bold {
							pt, _ = cRedBold.DrawString(word.Text, pt)
						} else {
							pt, _ = cRed.DrawString(word.Text, pt)
						}
					} else {
						if word.Bold {
							pt, _ = cGreenBold.DrawString(word.Text, pt)
						} else {
							pt, _ = cGreen.DrawString(word.Text, pt)
						}
					}
					pt, _ = cGreen.DrawString(" ", pt)
				}
			}
			textInputStart = pt

			textDirty = false
		}

		pt := textInputStart
		for _, word := range con.Input {
			pt, _ = cInput.DrawString(word.Text, pt)
			pt, _ = cInput.DrawString(" ", pt)
		}

		draw.Draw(textBufCursor, textBufCursor.Bounds(), image.Transparent, image.ZP, draw.Src)
		font.Context(textBufCursor, true).DrawString("_", pt)

		textInputDirty = false
	}

	if (frame/FramesPerSecond)%4 == 0 {
		draw.DrawMask(dst, dst.Bounds(), textColorCursor, image.ZP, textBufCursor, image.ZP, draw.Over)
	}
}

func paintClearText(dst *image.RGBA) {
	draw.DrawMask(dst, dst.Bounds(), textColorGreen, image.ZP, textBufGreen, image.ZP, draw.Over)
	draw.DrawMask(dst, dst.Bounds(), textColorRed, image.ZP, textBufRed, image.ZP, draw.Over)
	draw.DrawMask(dst, dst.Bounds(), textColorGreen, image.ZP, textBufInput, image.ZP, draw.Over)
}
