package room

func init() {
	rooms[Coord{0, 0}] = &Room{
		Desc: `
A tattered BLANKET almost obscures the doorway to the NORTH.`,
		Pass: [directions]bool{
			North: true,
		},
		Direct: [directions]string{
			North: `A tattered blanket covers most of the doorway.
Cold air seeps through the small area of the doorway that is uncovered.`,
		},
		Object: map[string]*Object{
			"blanket": &Object{
				Desc: "It's old and tattered and has strange yellow stains on it.",
			},
		},
	}
}
