package advent

import (
	"os"
	"path/filepath"
)

type DaySolver interface {
	Solve() (string, error)
}

func getResource(resourceFile string) (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	inputFile := filepath.Join(filepath.Dir(wd), "resources", resourceFile) //"input-01.txt"

	return os.Open(inputFile)
}
