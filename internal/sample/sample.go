package day02

import (
	"strconv"
	"strings"
)

// This structure is different for each day and should be optimized for storing
// the input data in a way achieving simple puzzle solving afterwards.
type inputPart struct {
	value int
}

// Variable input obviously contains the day input which one should pass to the
// program using stdin pipe in console (./program <day> < input.txt). The input
// string is already processed by this day function (RegisterInput).
var input []inputPart

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	// split the string to the command and its valud
	split := strings.Split(data, " ")

	// try to convert the string to the integer
	num, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}

	// store the number in input var
	input = append(input, inputPart{
		value: num,
	})
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (int, error) {
	result := 0
	return result, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (int, error) {
	result := 0
	return result, nil
}
