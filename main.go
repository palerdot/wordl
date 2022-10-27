package main

import (
	"github.com/palerdot/wordl/ui"
)

func main() {
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
