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

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	adventDay, err := parseArgs()
	if err != nil {
		panic(err)
	}

	var solver advent.DaySolver
	switch adventDay {
	case 1:
		solver = advent.DayOne{}
	case 2:
		solver = advent.DayTwo{}
	default:
		e := fmt.Sprintf("Day %d does not exist", adventDay)
		panic(e)
	}
	solution, err := solver.Solve()
	if err != nil {
		panic(err)
	}
	logrus.Infof("Solution for day %d is: %s\n", adventDay, solution)
}

func parseArgs() (int, error) {
	args := os.Args[1:]

	if len(args) == 1 {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			e := fmt.Sprint("Positional argument could not be converted to int (ex: 01, 02, 03)")
			return 0, errors.New(e)
		}
		return day, nil
	}
	e := fmt.Sprint("You must provide a [Day] argument. Example: `go run . [Day]`")
	return 0, errors.New(e)
}
