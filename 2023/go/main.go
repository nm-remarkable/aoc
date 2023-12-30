package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	advent "advent/days"
)

func main() {
	adventDay, err := parseArgs()
	if err != nil {
		panic(err)
	}
	slog.Info(fmt.Sprintf("Solving day: %d", adventDay))

	var solver advent.DaySolver
	switch adventDay {
	case 1:
		solver = advent.DayOne{}
	case 2:
		solver = advent.DayTwo{}
	}
	solution, err := solver.Solve()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Solution for day %d is: %s\n", adventDay, solution)
}

func parseArgs() (int, error) {
	args := os.Args[1:]
	exePath := filepath.Base(os.Args[0])
	slog.Debug(fmt.Sprintf("Executing: %s %v", exePath, args))

	if len(args) == 1 {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			e := fmt.Sprintf("Positional argument could not be converted to int (ex: 01, 02, 03)")
			return 0, errors.New(e)
		}
		return day, nil
	}
	e := fmt.Sprintf("Wrong arguments, example: %s [Day] --debug", exePath)
	return 0, errors.New(e)
}
