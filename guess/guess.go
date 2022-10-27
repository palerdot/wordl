package guess

import (
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
