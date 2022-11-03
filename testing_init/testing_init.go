package testing_init

import (
	"os"
	"path"
	"runtime"
)

// change the root path to project root so that file reading works fine
// ref: https://stackoverflow.com/questions/22457861/how-to-test-go-programs-with-relative-directories
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
