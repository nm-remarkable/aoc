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
	row, value int
	position   []int
	valid      bool
}

type Symbol struct {
	x, y int
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

func validateGear(p PartNumber) (Symbol, error) {
	for i := p.row - 1; i <= p.row+1; i++ {
		if symbolRow, ok := symbols[i]; ok {
			for j := p.position[0]; j <= p.position[1]+1; j++ {
				for _, s := range symbolRow {
					if s == j {
						return Symbol{i, j}, nil
					}
				}
			}
		}
	}
	return Symbol{}, fmt.Errorf("this should have been an optional")
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
	f, err := advent.GetResource()
	if err != nil {
		return "", err
	}
	parseData(f)
	gears := make(map[Symbol][]PartNumber, 0)
	for _, p := range parts {
		if s, err := validateGear(p); err == nil {

			if _, ok := gears[s]; !ok {
				gears[s] = []PartNumber{}
			}
			gears[s] = append(gears[s], p)
		}
	}

	sum := 0
	for _, g := range gears {
		inner_sum := 1
		if len(g) < 2 {
			continue
		}
		for _, p := range g {
			inner_sum *= p.value
		}
		sum += inner_sum
		inner_sum = 1
	}

	return fmt.Sprint(sum), nil
}

func main() {
	advent.Challenge(DayThree{})
}
