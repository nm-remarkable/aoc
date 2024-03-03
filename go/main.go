package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"

	advent "advent/days"
)

type challenge struct {
	day        int
	first_part bool
}

var (
	err      error
	solution string
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	chall, err := parseArgs()
	if err != nil {
		panic(err)
	}

	var solver advent.DaySolver
	switch chall.day {
	case 1:
		solver = advent.DayOne{}
	case 2:
		solver = advent.DayTwo{}
	default:
		e := fmt.Sprintf("Day %d does not exist", chall.day)
		panic(e)
	}
	if chall.first_part {
		solution, err = solver.First()
	} else {
		solution, err = solver.Second()
	}
	if err != nil {
		panic(err)
	}
	logrus.Infof("Solution for day %d is: %s\n", chall.day, solution)
}

func parseArgs() (challenge, error) {
	args := os.Args[1:]

	if len(args) > 0 {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			e := fmt.Sprint("Positional argument could not be converted to int (ex: 01, 02, 03)")
			return challenge{}, errors.New(e)
		}

		var first_part bool = true
		if len(args) == 2 {
			if arg, _ := strconv.Atoi(args[1]); arg != 0 {
				first_part = false
			}
		}
		return challenge{day, first_part}, nil
	}
	e := fmt.Sprint("You must provide a [Day] argument. Example: `go run . [Day]`")
	return challenge{}, errors.New(e)
}
