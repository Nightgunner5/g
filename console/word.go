package console

import (
	"strings"
)

type Word struct {
	Text string
	Red  bool
	Bold bool
}

func (w Word) IsNewline() bool {
	return w.Text == ""
}

func AppendWords(base []Word, text string, red, bold bool) []Word {
	for _, s := range strings.FieldsFunc(text, func(r rune) bool {
		return r == ' '
	}) {
		for _, word := range strings.Split(s, "\n") {
			base = append(base, Word{
				Text: word,
				Red:  red,
				Bold: bold,
			})
		}
	}
	return base
}
