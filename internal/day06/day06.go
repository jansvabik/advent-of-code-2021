package day06

import (
	"strconv"
	"strings"
)

const days = 8

var fishArray = [days + 1]int{}

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	for _, counter := range strings.Split(data, ",") {
		counterInt, _ := strconv.Atoi(counter)
		fishArray[counterInt]++
	}
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	// process all the fish
	for i := 0; i < 80; i++ {
		fishArray[(i+7)%9] += fishArray[i%9]
	}

	// summarize the data
	sum := 0
	for _, v := range fishArray {
		sum += v
	}
	return sum, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	// process all the fish
	for i := 0; i < 256; i++ {
		fishArray[(i+7)%9] += fishArray[i%9]
	}

	// summarize the data
	sum := 0
	for _, v := range fishArray {
		sum += v
	}
	return sum, nil
}
