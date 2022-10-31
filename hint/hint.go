package hint

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/palerdot/wordl/guess"
	"github.com/palerdot/wordl/ui"
)

// keyboard rows
var Rows [3]string = [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

// keyboard dimensions
var sizeX = 4
var sizeY = 2

type HintInfo struct {
	Correct   map[string]bool
	Incorrect map[string]bool
	Missing   map[string]bool
}

type HintStatus struct {
	Previous HintInfo
	Current  HintInfo
}

// hint status; contains info on keyboard letter hints
var initialLetterStatus = HintInfo{
	Correct:   make(map[string]bool),
	Incorrect: make(map[string]bool),
	Missing:   make(map[string]bool),
}

var LetterStatus HintStatus = HintStatus{
	Previous: initialLetterStatus,
	Current:  initialLetterStatus,
}

// hint keyboard
// shows previous hint (at the start of screen paint) or latest hint
func DrawKeyboard(s tcell.Screen, isLatest bool) {
	var info HintInfo
	if isLatest {
		info = LetterStatus.Current
	} else {
		info = LetterStatus.Previous
	}

	for row := 0; row < len(Rows); row++ {
		for col := 0; col < len(Rows[row]); col++ {
			letter := string(Rows[row][col])
			style := getLetterColor(info, letter)
			drawKeyboardLetter(s, row, col, style, letter)
		}
	}

	// if latest; copy latest info to previous for next guess
	if isLatest {
		LetterStatus.Previous = LetterStatus.Current
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
	startY := gridHeight + 10

	x1 := startX + (col * size.X) + (space * col)
	y1 := startY + (row * size.Y) + (space * row) - 4
	x2 := x1 + size.X
	y2 := y1 + size.Y

	ui.DrawBox(s, x1, y1, x2, y2, size, style, letter, false)
}

// update letter hint
func UpdateLetter(pos guess.LetterPosition, l string) {
	letter := strings.ToLower(l)

	if pos == guess.LetterPositionCorrect {
		LetterStatus.Current.Correct[letter] = true

		return
	}
	if pos == guess.LetterPositionInCorrect {
		LetterStatus.Current.Incorrect[letter] = true

		return
	}
	if pos == guess.LetterPositionMissing {
		LetterStatus.Current.Missing[letter] = true

		return
	}

	return
}

// helper function to get letter color
func getLetterColor(info HintInfo, letter string) ui.PositionStyle {
	// order of style calculation
	// 1. Correct position
	// 2. InCorrect position
	// 3. Missing position
	// if no info is present blank position

	correct, ok := info.Correct[letter]
	if ok && correct {
		// 1. Correct position
		return ui.GetLetterStyles(guess.LetterPositionCorrect)
	}

	incorrect, ok := info.Incorrect[letter]
	if ok && incorrect {
		// 2. InCorrect position
		return ui.GetLetterStyles(guess.LetterPositionInCorrect)
	}

	missing, ok := info.Missing[letter]
	if ok && missing {
		// 3. Missing position
		return ui.GetLetterStyles(guess.LetterPositionMissing)
	}

	return ui.GetLetterStyles(guess.LetterPositionBlank)
}
