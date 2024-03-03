package advent

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var (
	err      error
	solution string
)

func Challenge(d DaySolver) {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	if !runSecondHalf() {
		solution, err = d.First()
	} else {
		solution, err = d.Second()
	}
	if err != nil {
		panic(err)
	}
	logrus.Infof("Solution for %s is: %s\n", reflect.TypeOf(d), solution)
}

func runSecondHalf() bool {
	args := os.Args[1:]

	if len(args) > 0 {
		if arg, _ := strconv.Atoi(args[0]); arg == 0 {
			return true
		}
	}
	return false
}

type DaySolver interface {
	First() (string, error)
	Second() (string, error)
}

func GetResource() (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	// Get name of file that calls this function
	_, fpath, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("cannot find filename of executable")
	}
	dir := filepath.Dir(fpath)
	filename := filepath.Base(dir)
	inputFile := filepath.Join(filepath.Dir(wd), "resources", strings.Split(filename, ".")[0], "input.txt")
	return os.Open(inputFile)
}
