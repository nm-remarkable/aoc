package advent

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
)

type DayOne struct {
}

var wordToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func intify(number string) (int, error) {
	if val, ok := wordToNumber[number]; ok {
		return val, nil
	}
	return strconv.Atoi(number)
}

func MatchDigit(input string, reverse bool) (string, error) {
	line := []rune(input) // Reverse works on the reference, so we need to assign line to variable
	numbers := []rune(`one|two|three|four|five|six|seven|eight|nine`)
	digits := `\d` + `|`

	if reverse {
		slices.Reverse(line)
		slices.Reverse(numbers)
	}
	re := regexp.MustCompile(digits + string(numbers))
	match := re.Find([]byte(string(line)))

	if reverse {
		slices.Reverse(match)
	}
	digit, err := intify(string(match))
	if err != nil {
		return "", fmt.Errorf("cannot convert to integer: %v", err)
	}
	return fmt.Sprintf("%d", digit), nil
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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first_match, err := MatchDigit(line, false)
		if err != nil {
			return "", err
		}
		last_match, err := MatchDigit(line, true)
		if err != nil {
			return "", err
		}

		add, err := strconv.Atoi(first_match + last_match)
		if err != nil {
			return "", err
		}
		sum += add
	}
	return fmt.Sprintf("%d", sum), nil
}
