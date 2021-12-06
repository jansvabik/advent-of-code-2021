package day01

import (
	"strconv"
)

// input obviously contains the day input which one should pass to the program
// using stdin pipe in console (./program <day> < input.txt).
var input []int

// This variable holds the value of previous processed input for further
// processing (e.g. comparison with the new input). The default value is the
// maximum int data type value because we don't want to count the first input.
var lastInput = int(^uint(0) >> 1)

// The increments variable holds the number of increments when comparison with
// the previous input value. It is used as a part 1 result.
var increments = 0

// Variable groupSums is used for summarizing numbers in groups specified by
// the second part. The summarization is done during input registering.
var groupSums []int

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

	// increment the number of value increments if the number is greater than
	// the previous one, we need this value for part 1
	if num > lastInput {
		increments++
	}
	lastInput = num

	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
// It is calculated during the input registering for faster processing so
// this function only returns the result.
func Part1() (int, error) {
	return increments, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (int, error) {
	lastGroupSum := int(^uint(0) >> 1)
	incrementsPt2 := 0
	for i := range input {
		if len(input) > i+2 {
			sum := input[i] + input[i+1] + input[i+2]
			if sum > lastGroupSum {
				incrementsPt2++
			}
			lastGroupSum = sum
		}
	}
	return incrementsPt2, nil
}
