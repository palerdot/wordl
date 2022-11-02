package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/palerdot/wordl/guess"
)

type Dimension struct {
	X int
	Y int
}

func Greet() {
	fmt.Println("Happy Wordle !!!")
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

// generic function to draw box
func DrawBox(s tcell.Screen, x1 int, y1 int, x2 int, y2 int, size Dimension, style PositionStyle, text string, isBold bool) {
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
	for row := y1 + 1; row < y2; row++ {
		for col := x1 + 1; col < x2; col++ {
			s.SetContent(col, row, ' ', nil, boxStyle)
		}
	}

	xDiff := size.X / 2
	yDiff := size.Y / 2

	letterStyle = letterStyle.Bold(isBold)
	DrawText(s, x1+xDiff, y1+yDiff, x2-xDiff, y2-yDiff, letterStyle, text)
}

func DrawText(s tcell.Screen, x1 int, y1 int, x2 int, y2 int, style tcell.Style, text string) {
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

// helper function to get letter styles
// there are 4 variations
// Correct position - green
// InCorrect position - orange
// Missing position - gray
// Neutral/default position - transparent/black

type PositionStyle struct {
	box    tcell.Style
	letter tcell.Style
}

func GetLetterStyles(posType guess.LetterPosition) PositionStyle {
	switch posType {

	case guess.LetterPositionMissing:
		{
			boxStyle := tcell.StyleDefault.Background(tcell.Color242).Foreground(tcell.Color242)
			letterStyle := tcell.StyleDefault.Background(tcell.Color242).Foreground(tcell.ColorWhite)

			style := PositionStyle{
				box:    boxStyle,
				letter: letterStyle,
			}

			return style
		}

	case guess.LetterPositionInCorrect:
		{
			boxStyle := tcell.StyleDefault.Background(tcell.Color178).Foreground(tcell.ColorWhite)
			letterStyle := tcell.StyleDefault.Background(tcell.Color178).Foreground(tcell.ColorWhite)

			style := PositionStyle{
				box:    boxStyle,
				letter: letterStyle,
			}

			return style
		}

	case guess.LetterPositionCorrect:
		{
			boxStyle := tcell.StyleDefault.Background(tcell.Color28).Foreground(tcell.ColorWhite)
			letterStyle := tcell.StyleDefault.Background(tcell.Color28).Foreground(tcell.ColorWhite)

			style := PositionStyle{
				box:    boxStyle,
				letter: letterStyle,
			}

			return style
		}

	default:
		{
			boxStyle := tcell.StyleDefault.Background(tcell.Color236).Foreground(tcell.ColorWhite)
			letterStyle := tcell.StyleDefault.Background(tcell.Color236).Foreground(tcell.ColorWhite)

			style := PositionStyle{
				box:    boxStyle,
				letter: letterStyle,
			}

			return style
		}
	}

}

// color letter
func ColorLetter(wordle string, col int, letter string) (PositionStyle, guess.LetterPosition) {
	var pos = guess.FindLetterPosition(wordle, col, letter)

	return GetLetterStyles(pos), pos
}
