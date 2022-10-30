package guess

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

// guess word length
var WordLength int = 5

// total tries
var TotalTries int = 6

// active guess index
var ActiveIndex int = 0

// flag to indicate if word is guessed
var IsOver bool = false

// flag to decide if the user has correctly guessed
var IsSuccess bool = false

// valid wordle list
var wordleList []string = getValidAnswerList()

// valid guess list
var validGuessList []string = getValidGuessList()

// valid final list
var ValidList []string = append(validGuessList, wordleList...)

// tried guesses
var Tries = [6]string{}

// wordle for the game
var Wordle string = getWordle()

// grid dimensions
// wordle letter dimension
var LetterSizeX = 8
var LetterSizeY = 4

// keyboard hint dimensions
var KbdSizeX = 4
var KbdSizeY = 2

// helper function to reset wordle
func ResetWordle() {
	// active guess index
	ActiveIndex = 0
	// flag to indicate if word is guessed
	IsOver = false
	// flag to decide if the user has correctly guessed
	IsSuccess = false
	// tried guesses
	Tries = [6]string{}
	// wordle for the game
	Wordle = getWordle()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

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

// helper function to calculate color letter
// takes in index and letter
// we will compare the letter at specified index of the wordle
// returns LetterPosition
func FindLetterPosition(index int, letter string) (pos LetterPosition) {
	if index >= WordLength {
		return LetterPositionBlank
	}

	var wordleLetter string = string(Wordle[index])
	// CASE 1: letter is in correct position
	if strings.EqualFold(wordleLetter, letter) {
		return LetterPositionCorrect
	} else {
		// CASE 2: letter is in incorrect position
		if strings.Contains(Wordle, letter) {
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
func HandleLetter(letter rune) (row int, col int, err error) {
	// check if current word is full
	var currentWord = Tries[ActiveIndex]
	var isFull bool = len(currentWord) == WordLength

	// if full ignore letters
	if isFull {
		return row, col, errors.New("word already full")
	}
	// append the letter to the word
	currentWord = currentWord + string(letter)
	// update the original Tries
	Tries[ActiveIndex] = currentWord

	return ActiveIndex, len(currentWord) - 1, nil
}

// calculate word on hitting enter
// 1. Enter is hit before word is complete
// 2. Enter is hit after word is complete
func Calculate() (err error) {
	// check if current word is full
	var currentWord = Tries[ActiveIndex]
	var isFull bool = len(currentWord) == WordLength

	if !isFull {
		return errors.New("word not yet full")
	}

	// we have a full word
	// there are 2 cases;
	// 1 - valid guess word (proceed to color the word)
	// 2 - invalid guess word (clear current word)

	// CASE 1: valid guess word
	// 1a: word is the Wordle
	if currentWord == Wordle {
		// shift to next word
		// this is needed for animating the correct guess
		ActiveIndex = ActiveIndex + 1
		// mark complete
		IsOver = true
		IsSuccess = true

		return nil
	}

	// CASE 2: word is valid guess but not wordle
	if isValidGuess(currentWord) {
		// shift to next word
		ActiveIndex = ActiveIndex + 1
		// mark over if user has reached six tries
		if ActiveIndex == TotalTries {
			// mark game as over
			IsOver = true
			IsSuccess = false
		}

		return nil
	} else {
		// CASE 3: word is not valid guess
		// clear the current word
		Tries[ActiveIndex] = ""

		return errors.New("Invalid word")
	}
}

// clear word on backspace
// return row, col to clear the letter
// error if letter is not be cleared
func ClearLetter() (row int, col int, err error) {
	// check if current word is full
	var currentWord = Tries[ActiveIndex]
	var isEmpty bool = len(currentWord) == 0

	if isEmpty {
		return row, col, errors.New("word already empty")
	}

	// clear the letter
	var position int = len(currentWord) - 1
	Tries[ActiveIndex] = currentWord[0:position]

	return ActiveIndex, position, nil
}

// find if word is a valid guess
func isValidGuess(currentGuess string) bool {
	for _, word := range ValidList {
		if strings.EqualFold(word, currentGuess) {
			return true
		}
	}

	return false
}
