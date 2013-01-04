package console

import (
	"os"
)

type Console struct {
	Output []Word
	Prompt []Word
	Input  []Word
}

func New() *Console {
	return &Console{
		Output: append(AppendWords([]Word{Word{">", false, true}},
			os.Args[0], false, false), Word{}),
		Prompt: []Word{
			Word{">", false, true},
		},
		Input: []Word{},
	}
}
