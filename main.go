package main

import (
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/palerdot/wordl/ui"
)

func main() {
	encoding.Register()
	ui.Greet()

	s := ui.InitScreen()

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

	ui.Render(s)
	ui.Listen(s)
}
