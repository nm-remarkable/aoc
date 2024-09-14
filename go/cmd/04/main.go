package main

import (
	"advent/advent"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func convert(line string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func winningPoints(line string) (int, error) {
	s := strings.Split(strings.Split(line, ":")[1], "|")
	winning := s[1]
	fmt.Println(winning)
	// choices := s[1]
	w, err := convert(winning)
	if err != nil {
		return 0, err
	}
	for _, ww := range w {
		return ww, nil
	}
	return 1, nil
}

func (d DayFour) First() (string, error) {
	f, err := advent.GetResource()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		p, err := winningPoints(line)
		if err != nil {
			return "", err
		}
		sum += p
	}
	return fmt.Sprintf("%d", sum), nil
}

func (d DayFour) Second() (string, error) {
	return d.First()
}

type DayFour struct {
}

func main() {
	advent.Challenge(DayFour{})
}
