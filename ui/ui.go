package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/palerdot/wordl/guess"
)

func Greet() {
	fmt.Println("porumai ... all the best with Wordle!!!")
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
	default:
		{
			boxStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.Color245)
			letterStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorWhite)

			style := PositionStyle{
				box:    boxStyle,
				letter: letterStyle,
			}

			return style
		}
	}

}
