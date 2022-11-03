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
