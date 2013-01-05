package console

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	"os"
)

type Console struct {
	Output     [][]Word
	Input      []Word
	InputSpace bool

	Exec func([]string)

	Version      uint64
	InputVersion uint64
}

func New() *Console {
	return &Console{
		Output:       AppendWords([][]Word{[]Word{Word{">", false, true}}}, os.Args[0]+"\n", false, false),
		InputSpace:   false,
		Version:      1,
		InputVersion: 1,
	}
}

func (c *Console) Print(args ...interface{}) {
	c.print(fmt.Sprint(args...), false, false)
}

func (c *Console) Println(args ...interface{}) {
	c.print(fmt.Sprintln(args...), false, false)
}

func (c *Console) Printf(format string, args ...interface{}) {
	c.print(fmt.Sprintf(format, args...), false, false)
}

func (c *Console) Bprint(args ...interface{}) {
	c.print(fmt.Sprint(args...), false, true)
}

func (c *Console) Bprintln(args ...interface{}) {
	c.print(fmt.Sprintln(args...), false, true)
}

func (c *Console) Bprintf(format string, args ...interface{}) {
	c.print(fmt.Sprintf(format, args...), false, true)
}

func (c *Console) Error(args ...interface{}) {
	c.print(fmt.Sprint(args...), true, false)
}

func (c *Console) Errorln(args ...interface{}) {
	c.print(fmt.Sprintln(args...), true, false)
}

func (c *Console) Errorf(format string, args ...interface{}) {
	c.print(fmt.Sprintf(format, args...), true, false)
}

func (c *Console) Berror(args ...interface{}) {
	c.print(fmt.Sprint(args...), true, true)
}

func (c *Console) Berrorln(args ...interface{}) {
	c.print(fmt.Sprintln(args...), true, true)
}

func (c *Console) Berrorf(format string, args ...interface{}) {
	c.print(fmt.Sprintf(format, args...), true, true)
}

func (c *Console) print(text string, red, bold bool) {
	c.Output = AppendWords(c.Output, text, red, bold)
	c.Version++
}

func (c *Console) Typed(key, glyph string) {
	if key == wde.KeyReturn {
		argv := make([]string, len(c.Input))
		for i, w := range c.Input {
			argv[i] = w.Text
		}
		go c.Exec(argv)
		c.Input = nil
		c.InputSpace = false
		c.InputVersion++
		return
	}

	if key == wde.KeySpace {
		c.InputSpace = true
		c.InputVersion++
		return
	}

	if key == wde.KeyBackspace {
		if c.InputSpace {
			c.InputSpace = false
		} else if len(c.Input) > 0 {
			i := len(c.Input) - 1
			c.Input[i].Text = c.Input[i].Text[:len(c.Input[i].Text)-1]
			if c.Input[i].Text == "" {
				c.Input = c.Input[:i]
			}
		}
		c.InputVersion++
		return
	}

	if glyph == "" {
		return
	}

	if c.InputSpace || len(c.Input) == 0 {
		c.Input = append(c.Input, Word{})
		c.InputSpace = false
	}

	c.Input[len(c.Input)-1].Text += glyph
	c.InputVersion++
}
