package advent

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

type DayOne struct {
}

func (d DayOne) Solve() (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	inputFile := filepath.Join(filepath.Dir(wd), "resources", "input-01.txt")

	f, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	// counter := 1 // Used to print pretty lines
	sum := 0
	for scanner.Scan() {
		var first int
		var last int
		line := scanner.Text()
		//fmt.Printf("Line %02d: %s\n", counter, line)

		// Part 1 can be done with a double loop
		for i := range line {
			left := rune(line[i]) // iterate from the left
			if unicode.IsDigit(left) {
				first = int(left - '0')
				break
			}
		}
		for i := range line {
			right := rune(line[len(line)-i-1]) // iterate from the right
			if unicode.IsDigit(right) {
				last = int(right - '0')
				break
			}
		}
		// Must convert back and forth from int to string, as we want "3" + "5" = "35", and not 7
		add, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		if err != nil {
			return "", err
		}
		//counter++
		sum += add
	}
	return fmt.Sprintf("%d", sum), nil
}
