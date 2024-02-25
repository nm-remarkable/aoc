package advent

import (
	"os"
	"path/filepath"
)

type DaySolver interface {
	Solve() (string, error)
}

func getResource(day string) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	inputFile := filepath.Join(filepath.Dir(wd), "resources", day, "input.txt")

	return os.Open(inputFile)
}
