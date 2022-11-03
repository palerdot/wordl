package guess

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Wordle Data
// loaded once during the lifetime of the game

// valid wordle list
var wordleList []string = getValidAnswerList()

// valid guess list
var validGuessList []string = getValidGuessList()

// valid guess list (does not include valid wordle list)
func getValidGuessList() []string {
	data, err := os.ReadFile("data/guess.txt")
	check(err)
	var guessList = strings.Split(string(data), "\n")

	return guessList
}

// valid answer list
func getValidAnswerList() []string {
	data, err := os.ReadFile("data/answer.txt")
	check(err)
	var splitted = strings.Split(string(data), "\n")

	return splitted
}

// guess word length
func GetWordLength() int {
	return 5
}

// total tries
func GetTotalTries() int {
	return 6
}

// grid dimensions
// wordle letter dimension
func GetSize() (x int, y int) {
	x, y = 8, 4

	return x, y
}

// guess state
type GuessState struct {
	// active guess index
	ActiveIndex int
	// flag to indicate if word is guessed
	IsOver bool
	// flag to decide if the user has correctly guessed
	IsSuccess bool
	// valid final list
	ValidList []string
	// tried guesses
	Tries [6]string
	// wordle for the game
	Wordle string
}

// initial guess state
func GetInitialState() GuessState {
	return GuessState{
		ActiveIndex: 0,
		IsOver:      false,
		IsSuccess:   false,
		ValidList:   append(validGuessList, wordleList...),
		Tries:       [6]string{},
		Wordle:      getWordle(),
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getWordle() string {
	rand.Seed(time.Now().UnixNano())
	var randomIndex = rand.Intn(len(wordleList))
	var word string = wordleList[randomIndex]

	return word
}

// letter position related stuffs
type LetterPosition string

const (
	// letter is in correct position
	LetterPositionCorrect LetterPosition = "correct"
	// letter is NOT in correct position
	LetterPositionInCorrect LetterPosition = "incorrect"
	// letter is not in word
	LetterPositionMissing LetterPosition = "missing"
	// neutral state; we have not yet calculated letter position
	LetterPositionBlank LetterPosition = "blank"
)

// helper function to calculate letter color
// takes in index and letter
// we will compare the letter at specified index of the wordle
// returns LetterPosition
func FindLetterPosition(wordle string, col int, letter string) (pos LetterPosition) {
	var wordleLetter = string(wordle[col])
	// CASE 1: letter is in correct position
	if strings.EqualFold(wordleLetter, letter) {
		// update keyboard hint
		return LetterPositionCorrect
	} else {
		// CASE 2: letter is in incorrect position
		if strings.Contains(wordle, letter) {
			return LetterPositionInCorrect
		}

		// CASE 3: letter is not present in the word
		return LetterPositionMissing
	}
}

// handle incoming letter
// if the current word is not complete, append the letter
// and notify the ui that ui has to be updated
// if the current word is full, just ignore the incoming letters
func HandleLetter(letter rune, state *GuessState) (row int, col int, err error) {
	// check if current word is full
	var currentWord *string = &state.Tries[state.ActiveIndex]
	var isFull bool = len(*currentWord) == GetWordLength()

	// if full ignore letters
	if isFull {
		return row, col, errors.New("word already full")
	}
	// append the letter to the word
	*currentWord = *currentWord + string(letter)

	return state.ActiveIndex, len(*currentWord) - 1, nil
}

// calculate word on hitting enter
// 1. Enter is hit before word is complete
// 2. Enter is hit after word is complete
func Calculate(state *GuessState) (err error) {
	// check if current word is full
	var currentWord = state.Tries[state.ActiveIndex]
	var isFull bool = len(currentWord) == GetWordLength()

	if !isFull {
		return errors.New("word not yet full")
	}

	// we have a full word
	// there are 2 cases;
	// 1 - valid guess word (proceed to color the word)
	// 2 - invalid guess word (clear current word)

	// CASE 1: valid guess word
	// 1a: word is the Wordle
	if currentWord == state.Wordle {
		// shift to next word
		// this is needed for animating the correct guess
		state.ActiveIndex = state.ActiveIndex + 1
		// mark complete
		state.IsOver = true
		state.IsSuccess = true

		return nil
	}

	// CASE 2: word is valid guess but not wordle
	if isValidGuess(state.ValidList, currentWord) {
		// shift to next word
		state.ActiveIndex = state.ActiveIndex + 1
		// mark over if user has reached six tries
		if state.ActiveIndex == GetTotalTries() {
			// mark game as over
			state.IsOver = true
			state.IsSuccess = false
		}

		return nil
	} else {
		// CASE 3: word is not valid guess
		// clear the current word
		state.Tries[state.ActiveIndex] = ""

		return errors.New("Invalid word")
	}
}

// clear word on backspace
// return row, col to clear the letter
// error if letter is not be cleared
func ClearLetter(state *GuessState) (row int, col int, err error) {
	// check if current word is full
	var currentWord = state.Tries[state.ActiveIndex]
	var isEmpty bool = len(currentWord) == 0

	if isEmpty {
		return row, col, errors.New("word already empty")
	}

	// clear the letter
	var position int = len(currentWord) - 1
	state.Tries[state.ActiveIndex] = currentWord[0:position]

	return state.ActiveIndex, position, nil
}

// find if word is a valid guess
func isValidGuess(validList []string, currentGuess string) bool {
	for _, word := range validList {
		if strings.EqualFold(word, currentGuess) {
			return true
		}
	}

	return false
}
