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

	case guess.LetterPositionMissing:
		{
			// boxStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorRed)
			// letterStyle := tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorWhite)
			boxStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorBlue)
			letterStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)

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
func ColorLetter(index int, letter string) PositionStyle {
	var pos = guess.FindLetterPosition(index, letter)

	return GetLetterStyles(pos)
}
