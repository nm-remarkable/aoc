package advent

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type DaySolver interface {
	Solve() (string, error)
}

func getResource() (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// Get name of file that calls this function
	_, fpath, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("cannot find filename of executable")
	}
	filename := filepath.Base(fpath)
	inputFile := filepath.Join(filepath.Dir(wd), "resources", strings.Split(filename, ".")[0], "input.txt")
	return os.Open(inputFile)
}
