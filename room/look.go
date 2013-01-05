package room

import (
	"github.com/Nightgunner5/g/console"
	"github.com/Nightgunner5/g/paint"
	"strings"
)

func (r *Room) Look(argv []string) {
	if r == nil {
		paint.WithConsole(func(c *console.Console) {
			c.Println("You see nothing. You're not even a real person.")
		})
		return
	}

	if len(argv) == 1 {
		var output string

		switch a := strings.ToLower(argv[0]); a {
		case "north":
			if output = r.Direct[North]; output == "" {
				output = "Nothing stands out to the north."
			}

		case "south":
			if output = r.Direct[South]; output == "" {
				output = "Nothing stands out to the south."
			}

		case "east":
			if output = r.Direct[East]; output == "" {
				output = "Nothing stands out to the east."
			}

		case "west":
			if output = r.Direct[West]; output == "" {
				output = "Nothing stands out to the west."
			}

		case "dennis":
			if output = r.Direct[Dennis]; output == "" {
				output = "Dennis isn't a real direction. You're just being silly."
			}
		}

		if output != "" {
			paint.WithConsole(func(c *console.Console) {
				c.Println(output)
			})
			return
		}
	}

	if len(argv) != 0 {
		target := strings.ToLower(strings.Join(argv, " "))
		if d, ok := r.Object[target]; ok {
			paint.WithConsole(func(c *console.Console) {
				c.Println(d.Desc)
			})
		} else {
			paint.WithConsole(func(c *console.Console) {
				c.Errorln("I don't see any", strings.ToUpper(target)+".")
			})
		}
		return
	}

	var possibleExits []string
	if r.Pass[North] {
		possibleExits = append(possibleExits, "NORTH")
	}
	if r.Pass[South] {
		possibleExits = append(possibleExits, "SOUTH")
	}
	if r.Pass[East] {
		possibleExits = append(possibleExits, "EAST")
	}
	if r.Pass[West] {
		possibleExits = append(possibleExits, "WEST")
	}
	if r.Pass[Dennis] {
		possibleExits = append(possibleExits, "DENNIS")
	}

	var exits string
	switch l := len(possibleExits); l {
	case 0:
		exits = "There are no obvious exits."
	case 1:
		exits = "There is an exit to the " + possibleExits[0] + "."
	case 2:
		exits = "Possible exits are " + possibleExits[0] + " and " + possibleExits[1] + "."
	default:
		exits = "Possible exits are " + strings.Join(possibleExits[:l-1], ", ") + ", and " + possibleExits[l-1] + "."
	}

	paint.WithConsole(func(c *console.Console) {
		c.Println(r.Desc)
		c.Println(exits)
	})
}
