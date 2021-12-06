package day01

import (
	"strconv"
)

// input obviously contains the day input which one should pass to the program
// using stdin pipe in console (./program <day> < input.txt).
var input []int

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	// try to convert the string to the integer
	num, err := strconv.Atoi(data)
	if err != nil {
		return err
	}

	// store the number in input var
	input = append(input, num)
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	return len(input), nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	return len(input), nil
}
