package main

import "fmt"

type DayThree struct {
}

func (d DayThree) First() (string, error) {
	return "sad days", nil
}

func main() {
	s, _ := DayThree{}.First()
	fmt.Printf("%s", s)
}
