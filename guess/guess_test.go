package guess

import (
	"fmt"
	_ "github.com/palerdot/wordl/testing_init"
	"os"
	"testing"
)

var state GuessState

func TestMain(m *testing.M) {
	// setup initial state
	state = GetInitialState()
	fmt.Printf("porumai ... starting guess tests with wordle => '%s' \n", state.Wordle)
	// run the test
	exitVal := m.Run()
	fmt.Println("porumai ... ending guess tests")
	// exit the test
	os.Exit(exitVal)
}

func TestGetSize(t *testing.T) {
	x, y := GetSize()

	if x != 8 && y != 4 {
		t.Error("invalid wordle grid size")
	}
}

func TestWordLength(t *testing.T) {
	if GetWordLength() != 5 {
		t.Error("invalid word length")
	}
}

func TestTotalTries(t *testing.T) {
	if GetTotalTries() != 6 {
		t.Error("invalid total guess tries")
	}
}

func Test_isValidGuess(t *testing.T) {
	var validGuess = "pious"
	var invalidGuess = "aeiou"

	if isValidGuess(state.ValidList, validGuess) != true {
		t.Errorf("%s should be a valid guess", validGuess)
	}

	if isValidGuess(state.ValidList, invalidGuess) != false {
		t.Errorf("%s should be a invalid guess", invalidGuess)
	}
}
