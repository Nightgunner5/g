package console

import (
	"strings"
)

type Word struct {
	Text string
	Red  bool
	Bold bool
}

func AppendWords(base [][]Word, text string, red, bold bool) [][]Word {
	for _, s := range strings.FieldsFunc(text, func(r rune) bool {
		return r == ' '
	}) {
		for i, word := range strings.Split(s, "\n") {
			if i != 0 {
				base = append(base, nil)
			}
			if word != "" {
				i := len(base) - 1
				base[i] = append(base[i], Word{
					Text: word,
					Red:  red,
					Bold: bold,
				})
			}
		}
	}
	return base
}
