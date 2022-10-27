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

// wordle for the game
var Wordle string = getWordle()

// valid guess list
var ValidList = getValidGuessList()

// tried guesses
// var Tries = [6]string{"Hello", "Light", "Scout", "Aimer", "Foggy", "Clear"}
var Tries = [6]string{}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getValidGuessList() []string {
	data, err := os.ReadFile("data/guess.txt")
	check(err)
	var guessList = strings.Split(string(data), "\n")

	return guessList
}

func getWordle() string {
	rand.Seed(time.Now().UnixNano())
	data, err := os.ReadFile("data/answer.txt")
	check(err)
	var splitted = strings.Split(string(data), "\n")
	var randomIndex = rand.Intn(len(splitted))
	var word string = splitted[randomIndex]

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
