package ui

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

// ref: https://github.com/gdamore/tcell/blob/main/TUTORIAL.md
func drawText(s tcell.Screen, x1 int, y1 int, x2 int, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1

	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1 int, y1 int, x2 int, y2 int, style tcell.Style, text string) {
	// fix improper dimensions
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// draw rounded corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func InitScreen() tcell.Screen {
	// init screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	return s
}

func Setup(s tcell.Screen) {
	fmt.Println("porumai ... setting up screen ?")
	bgStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	// default style
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	// box style
	boxStyle := tcell.StyleDefault.Background(tcell.ColorTeal).Foreground(tcell.ColorWhite)
	// set default style
	s.SetStyle(defaultStyle)
	// clear the screen
	s.Clear()
	// draw initial boxes
	startX := 15
	startY := 2
	width := 50
	height := 5
	boxStyle.Bold(true)
	drawBox(s, startX, startY, width, height, boxStyle, "PORUMAI")
	boxStyle.Bold(false)
	drawBox(s, startX, height+startY, width, 2*height+startY, bgStyle, "H E L L O")
}

func Listen(s tcell.Screen) {
	fmt.Println("porumai ... listening screen ?")
	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	for {
		// update screen
		s.Show()

		// poll for event
		ev := s.PollEvent()

		// process event with type assertion
		switch ev := ev.(type) {
		// screen resize
		case *tcell.EventResize:
			s.Sync()
		// keyboard input
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC || ev.Key() == tcell.KeyEsc {
				s.Fini()
				os.Exit(0)
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				// s.Clear()
			} else {
				mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
				fmt.Sprintf("EventKey Modifiers: %d Key: %d Rune: %d", mod, key, ch)
			}

		}
	}
}
