package day02

import (
	"strconv"
	"strings"
)

// This structure is different for each day and should be optimized for storing
// the input data in a way achieving simple puzzle solving afterwards.
type inputPart struct {
	command string
	value   int
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
		command: split[0],
		value:   num,
	})
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (int, error) {
	horizontal, depth := 0, 0
	for _, v := range input {
		switch v.command {
		case "forward":
			horizontal += v.value
		case "up":
			depth -= v.value
		case "down":
			depth += v.value
		}
	}

	result := horizontal * depth
	return result, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (int, error) {
	horizontal, depth, aim := 0, 0, 0
	for _, v := range input {
		switch v.command {
		case "forward":
			horizontal += v.value
			depth += aim * v.value
		case "up":
			aim -= v.value
		case "down":
			aim += v.value
		}
	}

	result := horizontal * depth
	return result, nil
}
