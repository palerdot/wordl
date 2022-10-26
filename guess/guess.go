package guess

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// total tries
var TotalTries int = 6

// active guess index
var ActiveIndex int = 0

// flag to indicate if word is guessed
var IsComplete bool = false

// wordle for the game
var Wordle string = getWordle()

// tried guesses
// var Tries = [6]string{"Hello", "Light", "Scout", "Aimer", "Foggy", "Clear"}
var Tries = [6]string{}

func check(err error) {
	if err != nil {
		panic(err)
	}
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
