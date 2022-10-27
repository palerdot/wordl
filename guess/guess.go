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
var ValidList []string = getValidGuessList()

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

	// we are going to insert the valid answers as part of valid list
	ValidList = append(ValidList, splitted...)

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
		// mark complete
		IsOver = true
		IsSuccess = true

		return nil
	}

	// 1b: word is valid guess but not wordle
	if isValidGuess(currentWord) {
		// shift to next word
		ActiveIndex = ActiveIndex + 1

		return nil
	} else {
		// clear the current word
		Tries[ActiveIndex] = ""

		return nil
	}
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
