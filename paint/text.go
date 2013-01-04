package paint

import (
	"github.com/Nightgunner5/g/console"
	"github.com/Nightgunner5/g/font"
	"image"
	"image/color"
	"image/draw"
)

var (
	textColorGreen = image.NewUniform(color.NRGBA{0, 255, 0, 128})
	textColorRed   = image.NewUniform(color.NRGBA{255, 0, 0, 128})
	textBufGreen   *image.Alpha
	textBufRed     *image.Alpha
	textBufCursor  *image.Alpha
	textDirty      = true
	textVersion    uint64
)

func paintText(dst *image.RGBA, con *console.Console) {
	if textBufCursor == nil || textBufCursor.Bounds() != dst.Bounds() {
		textBufGreen = image.NewAlpha(dst.Bounds())
		textBufRed = image.NewAlpha(dst.Bounds())
		textBufCursor = image.NewAlpha(dst.Bounds())
		textDirty = true
	}

	if textVersion != con.Version {
		textVersion = con.Version
		textDirty = true
	}

	if textDirty {
		draw.Draw(textBufGreen, textBufGreen.Bounds(), image.Transparent, image.ZP, draw.Src)
		draw.Draw(textBufRed, textBufRed.Bounds(), image.Transparent, image.ZP, draw.Src)
		draw.Draw(textBufCursor, textBufCursor.Bounds(), image.Transparent, image.ZP, draw.Src)
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

		for _, word := range con.Input {
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

		font.Context(textBufCursor, true).DrawString("_", pt)

		textDirty = false
	}

	draw.DrawMask(dst, dst.Bounds(), textColorGreen, image.ZP, textBufGreen, image.ZP, draw.Over)
	draw.DrawMask(dst, dst.Bounds(), textColorRed, image.ZP, textBufRed, image.ZP, draw.Over)
	if (frame/FramesPerSecond)%4 == 0 {
		draw.DrawMask(dst, dst.Bounds(), textColorGreen, image.ZP, textBufCursor, image.ZP, draw.Over)
	}
}
