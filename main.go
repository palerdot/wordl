package main

import (
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/palerdot/wordl/screen"
)

func main() {
	encoding.Register()

	s, state := screen.Setup()

	// quit handler
	quit := func() {
		// catch panics
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}

	defer quit()

	screen.Render(s, &state)
	screen.Listen(s, &state)
}
