package screen

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/palerdot/wordl/guess"
	"github.com/palerdot/wordl/hint"
	"github.com/palerdot/wordl/ui"
)

// game state
type GameState struct {
	guess guess.GuessState
	hint  hint.HintStatus
}

// helper function to reset row
func resetRow(s tcell.Screen, row int) {
	style := ui.GetLetterStyles(guess.LetterPositionBlank)

	for col := 0; col < guess.GetWordLength(); col++ {
		drawGridLetter(s, row, col, style, "")
	}
}

func Setup() (tcell.Screen, GameState) {
	// init screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// game state
	state := GameState{
		guess: guess.GetInitialState(),
		hint:  hint.GetInitialState(),
	}

	return s, state
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

// wordle grid
func drawGrid(s tcell.Screen) {
	style := ui.GetLetterStyles(guess.LetterPositionBlank)

	for row := 0; row < guess.GetTotalTries(); row++ {
		for col := 0; col < guess.GetWordLength(); col++ {
			drawGridLetter(s, row, col, style, "")
		}
	}
}

// draws wordle letter
func drawGridLetter(s tcell.Screen, row int, col int, style ui.PositionStyle, letter string) {
	letterSizeX, letterSizeY := guess.GetSize()
	size := ui.Dimension{
		X: letterSizeX,
		Y: letterSizeY,
	}

	space := 0
	xmax, _ := s.Size()
	totalWidth := guess.GetWordLength()*size.X + ((guess.GetWordLength() - 1) * space)
	// startX := 15
	startX := (xmax - totalWidth) / 2
	startY := 5

	x1 := startX + (col * size.X) + (space * col)
	y1 := startY + (row * size.Y) + (space * row) - 4
	x2 := x1 + size.X
	y2 := y1 + size.Y

	ui.DrawBox(s, x1, y1, x2, y2, size, style, letter, true)
}

func populateGuess(s tcell.Screen, state *GameState) {
	for row, word := range state.guess.Tries {
		if row >= guess.GetTotalTries() {
			break
		}

		if len(word) != guess.GetWordLength() {
			return
		}

		for col, r := range word {
			var delay time.Duration

			if col == 0 {
				delay = 150
			} else {
				delay = 515
			}

			if row == state.guess.ActiveIndex-1 {
				// paint the screen
				s.Sync()
				time.Sleep(delay * time.Millisecond)
			}

			style, pos := ui.ColorLetter(state.guess.Wordle, col, string(r))
			// update hint status
			hint.UpdateLetter(&state.hint, pos, string(r))
			// draw grid letter
			letter := strings.ToUpper(string(r))
			drawGridLetter(s, row, col, style, letter)
		}
	}
}

func showGuessStatus(s tcell.Screen, guessState *guess.GuessState) {
	var style tcell.Style
	var status string
	var padding int

	letterSizeX, letterSizeY := guess.GetSize()
	size := ui.Dimension{
		X: letterSizeX,
		Y: letterSizeY,
	}

	xmax, _ := s.Size()
	totalWidth := guess.GetWordLength() * size.X
	gridHeight := guess.GetTotalTries()*size.Y + 2
	startX := (xmax - totalWidth) / 2

	if guessState.IsOver {
		// CASE 1: game is over: user guessed right
		if guessState.IsSuccess {
			style = tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorWhite)
			status = fmt.Sprintf("  Great! ")
			padding = 2 * size.X
		} else {
			// CASE 2: game is over: user didn't guess right
			style = tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorWhite)
			status = fmt.Sprintf("  Wordle: %s ", strings.ToUpper(guessState.Wordle))
			padding = 1*size.X + 4
		}
	} else {
		style = tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorWhite)
		status = fmt.Sprintf("  %d/%d left ", (guess.GetTotalTries() - guessState.ActiveIndex), guess.GetTotalTries())
		padding = 2*size.X - 1
	}

	ui.DrawText(s, startX+padding, gridHeight, startX+totalWidth, 55, style, status)
}

func displayStatus(s tcell.Screen) {
	letterSizeX, letterSizeY := guess.GetSize()
	size := ui.Dimension{
		X: letterSizeX,
		Y: letterSizeY,
	}

	xmax, _ := s.Size()
	totalWidth := guess.GetWordLength() * size.X
	gridHeight := guess.GetTotalTries()*size.Y + 2
	startX := (xmax - totalWidth) / 2

	// shos instructions
	infoStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.Color245)
	// display instructions
	ui.DrawText(s, startX-3*size.X, gridHeight+1, startX+totalWidth+3*size.X, 55, infoStyle, "Esc/Ctrl-C to Quit. Ctrl-N for new Wordle. Type and enter the guess. Backspace to clear.")

	// project url
	urlStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.ColorTeal)
	ui.DrawText(s, startX+(size.X/2), gridHeight+2, startX+totalWidth+(size.X/2), 55, urlStyle, "https://github.com/palerdot/wordl")
}

func Reset(state *GameState) {
	// reset wordle
	// guess.ResetWordle()
	state.guess = guess.GetInitialState()
	// reset hint data
	state.hint = hint.GetInitialState()
}

func Render(s tcell.Screen, state *GameState) {
	// clear the screen
	s.Clear()
	// draw bg
	drawBG(s)
	// draw grid
	drawGrid(s)
	// display status
	displayStatus(s)
	// show previous keyboard hint
	hint.DrawKeyboard(s, &state.hint, false)
	// populate guesses
	populateGuess(s, state)
	// guess status
	showGuessStatus(s, &state.guess)
	// show latest keyboard hint with data from latest guess
	hint.DrawKeyboard(s, &state.hint, true)
}

func Listen(s tcell.Screen, state *GameState) {
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
			} else if ev.Key() == tcell.KeyCtrlN {
				Reset(state)
				Render(s, state)
			} else {
				// if game is over do not handle keys
				if state.guess.IsOver {
					break
				}

				mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
				// handle enter key
				if key == tcell.KeyEnter {
					err := guess.Calculate(&state.guess)

					// if no error re-render
					if err == nil {
						Render(s, state)
					} else {
						// if invalid word clear the row
						if err.Error() == "Invalid word" {
							row := state.guess.ActiveIndex
							resetRow(s, row)
						}
					}

					break
				}

				// backspace/delete
				if key == tcell.KeyDelete || key == tcell.KeyBackspace || key == 127 {
					row, col, err := guess.ClearLetter(&state.guess)
					// if no error clear the letter
					if err == nil {
						style := ui.GetLetterStyles(guess.LetterPositionBlank)
						drawGridLetter(s, row, col, style, " ")
					}

					break
				}

				// 65 - 122; valid letters range
				if mod == 0 && ch >= 65 && ch <= 122 {
					// we have a valid character
					row, col, err := guess.HandleLetter(ch, &state.guess)
					// if no error populate letter
					if err == nil {
						// populate letter
						style := ui.GetLetterStyles(guess.LetterPositionBlank)
						drawGridLetter(s, row, col, style, strings.ToUpper(string(ch)))
					}
				}
			}

		}
	}
}
