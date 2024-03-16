package main

import (
	"advent/advent"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type DayThree struct {
}

type PartNumber struct {
	row      int
	value    int
	position []int
	valid    bool
}

var symbols = make(map[int][]int)
var parts []PartNumber = []PartNumber{}

func parseData(f *os.File) {
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()

		re := regexp.MustCompile(`(\d+|[^.])`)
		match := re.FindAllStringIndex(line, -1)

		for _, m := range match {
			n, err := strconv.Atoi(line[m[0]:m[1]])
			// Digits branch
			if err == nil {
				parts = append(parts, PartNumber{i, n, m, false})
				continue
			}
			// Symbols branch
			if _, ok := symbols[i]; !ok {
				symbols[i] = make([]int, 0)
			}
			symbols[i] = append(symbols[i], m[1])
		}
	}
}

func validatePart(p PartNumber) bool {
	for i := p.row - 1; i <= p.row+1; i++ {
		if symbolRow, ok := symbols[i]; ok {
			for j := p.position[0]; j <= p.position[1]+1; j++ {
				if slices.Contains(symbolRow, j) {
					return true
				}
			}
		}
	}
	return false
}

func (d DayThree) First() (string, error) {
	f, err := advent.GetResource()
	if err != nil {
		return "", err
	}
	parseData(f)
	sum := 0
	for _, p := range parts {
		if validatePart(p) {
			sum += p.value
		}
	}
	return fmt.Sprint(sum), nil
}

func (d DayThree) Second() (string, error) {
	return "Nop", nil
}

func main() {
	advent.Challenge(DayThree{})
}
