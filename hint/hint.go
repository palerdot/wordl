package hint

import (
	"github.com/gdamore/tcell/v2"
	"github.com/palerdot/wordl/guess"
	"github.com/palerdot/wordl/ui"
)

// keyboard rows
var Rows [3]string = [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

// keyboard dimensions
var sizeX = 4
var sizeY = 2

// hint keyboard
func DrawKeyboard(s tcell.Screen) {
	style := ui.GetLetterStyles(guess.LetterPositionBlank)

	for row := 0; row < len(Rows); row++ {
		for col := 0; col < len(Rows[row]); col++ {
			drawKeyboardLetter(s, row, col, style, string(Rows[row][col]))
		}
	}
}

func drawKeyboardLetter(s tcell.Screen, row int, col int, style ui.PositionStyle, letter string) {
	var size = ui.Dimension{
		X: sizeX,
		Y: sizeY,
	}

	space := 0
	xmax, _ := s.Size()
	totalWidth := guess.WordLength*guess.LetterSizeX + ((guess.WordLength - 1) * space)
	gridHeight := guess.TotalTries * guess.LetterSizeY
	// startX := 15
	startX := row*2 + (xmax-totalWidth)/2
	startY := gridHeight + 9

	x1 := startX + (col * size.X) + (space * col)
	y1 := startY + (row * size.Y) + (space * row) - 4
	x2 := x1 + size.X
	y2 := y1 + size.Y

	ui.DrawBox(s, x1, y1, x2, y2, size, style, letter)
}
