package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/palerdot/wordl/guess"
)

var sizeX = 6
var sizeY = 2

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

func drawBox(s tcell.Screen, x1 int, y1 int, x2 int, y2 int, style PositionStyle, text string) {
	boxStyle := style.box
	letterStyle := style.letter
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
			s.SetContent(col, row, ' ', nil, boxStyle)
		}
	}

	// draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, boxStyle)
		s.SetContent(col, y2, tcell.RuneHLine, nil, boxStyle)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, boxStyle)
		s.SetContent(x2, row, tcell.RuneVLine, nil, boxStyle)
	}

	// draw rounded corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, boxStyle)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, boxStyle)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, boxStyle)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, boxStyle)
	}

	xDiff := sizeX / 2
	yDiff := sizeY / 2

	letterStyle.Bold(true)
	drawText(s, x1+xDiff, y1+yDiff, x2-xDiff, y2-yDiff, letterStyle, text)
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

func drawBG(s tcell.Screen) {
	bgStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorWhite)
	xmax, ymax := s.Size()
	// fill background
	for row := 0; row <= ymax; row++ {
		for col := 0; col <= xmax; col++ {
			s.SetContent(col, row, ' ', nil, bgStyle)
		}
	}
}

func drawGrid(s tcell.Screen) {
	// for now default blank style
	style := GetLetterStyles(guess.LetterPositionBlank)

	for row := 0; row < guess.TotalTries; row++ {
		for col := 0; col < guess.WordLength; col++ {
			drawGridLetter(s, row, col, style, "")
		}
	}
}

func drawGridLetter(s tcell.Screen, row int, col int, style PositionStyle, letter string) {
	space := 1
	xmax, _ := s.Size()
	totalWidth := guess.WordLength*sizeX + ((guess.WordLength - 1) * space)
	// startX := 15
	startX := (xmax - totalWidth) / 2
	startY := 5

	x1 := startX + (col * sizeX) + (space * col)
	y1 := startY + (row * sizeY) + (space * row) - 3
	x2 := x1 + sizeX
	y2 := y1 + sizeY

	drawBox(s, x1, y1, x2, y2, style, letter)
}

func populateGuess(s tcell.Screen) {
	// for now default blank style
	style := GetLetterStyles(guess.LetterPositionBlank)

	for row, word := range guess.Tries {
		if row >= guess.TotalTries {
			break
		}

		if len(word) != guess.WordLength {
			return
		}

		for col, r := range word {
			letter := strings.ToUpper(string(r))
			drawGridLetter(s, row, col, style, letter)
		}
	}
}

func Render(s tcell.Screen) {
	// default style
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	// set default style
	s.SetStyle(defaultStyle)
	// clear the screen
	s.Clear()
	// draw bg
	drawBG(s)
	// draw grid
	drawGrid(s)
	// populate guesses
	populateGuess(s)
}

func Listen(s tcell.Screen) {
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
			} else {
				mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
				// handle enter key
				if key == tcell.KeyEnter {
					err := guess.Calculate()
					// if no error re-render
					if err != nil {
						fmt.Printf("not full")
					} else {
						Render(s)
					}

					break
				}

				// backspace/delete
				if key == tcell.KeyDelete || key == tcell.KeyBackspace || key == 127 {
					fmt.Printf("Clear")

					break
				}

				// 65 - 122; valid letters range
				if mod == 0 && ch >= 65 && ch <= 122 {
					// we have a valid character
					row, col, err := guess.HandleLetter(ch)
					// if no error populate letter
					if err == nil {
						// populate letter
						style := GetLetterStyles(guess.LetterPositionBlank)
						drawGridLetter(s, row, col, style, strings.ToUpper(string(ch)))
					}
				}
			}

		}
	}
}
