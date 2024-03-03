package advent

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DayTwo struct {
}
type RGBag struct {
	red   int
	green int
	blue  int
}

// Receive a "Game" result
func (bag RGBag) gameIsValid(game string) (bool, error) {
	matches := strings.Split(game, ";") // Separate each match from a game result
	re := regexp.MustCompile(`(\d+)\s(\w+)`)

	for _, m := range matches {
		colors := strings.Split(m, ",") // Separate each ball from the match result

		for _, c := range colors {
			match := re.FindStringSubmatch(c)
			if len(match) <= 0 {
				return false, fmt.Errorf("did not match string %s", game)
			}

			// Convert number of balls to number
			count, err := strconv.Atoi(match[1])
			if err != nil {
				return false, err
			}
			// If the number of balls with a certain color is higher than the allowed
			// it´s impossible for this match to have the result given
			switch match[2] {
			case "red":
				if bag.red < count {
					return false, nil
				}
			case "green":
				if bag.green < count {
					return false, nil
				}
			case "blue":
				if bag.blue < count {
					return false, nil
				}
			default:
				return false, fmt.Errorf("should never end up in default case for rgb")
			}
		}
	}
	return true, nil
}

func (d DayTwo) First() (string, error) {
	bag := RGBag{12, 13, 14}
	f, err := getResource()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		games := strings.Split(line, ":")

		isValid, err := bag.gameIsValid(games[1])
		if err != nil {
			return "", nil
		}
		if isValid {
			re := regexp.MustCompile(`\w+\s+(\d+)`)
			match := re.FindStringSubmatch(games[0])
			if len(match) <= 0 {
				return "", fmt.Errorf("did not match string %s", games[0])
			}
			gameNumber, err := strconv.Atoi(match[1])
			if err != nil {
				return "", err
			}
			sum += gameNumber
		}

	}
	return fmt.Sprint(sum), nil
}

func (d DayTwo) Second() (string, error) {
	return "", fmt.Errorf("Not solved")
}
