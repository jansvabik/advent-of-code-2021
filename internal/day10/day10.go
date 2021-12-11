package day10

import (
	"sort"
)

// Variable input obviously contains the day input which one should pass to the
// program using stdin pipe in console (./program <day> < input.txt). The input
// string is already processed by this day function (RegisterInput).
var inputs []string

// The table of points of the closing brackets.
var points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// The table of points of the closing brackets.
var autocompletePoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var oppositeBracket = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var autocompletePointsList = []int{}

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	inputs = append(inputs, data)
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	penaltyPoints := 0
	for _, line := range inputs {
		var openedBrackets []rune
		for _, b := range line {
			if b == '(' || b == '[' || b == '{' || b == '<' {
				openedBrackets = append(openedBrackets, b)
			} else {
				if openedBrackets[len(openedBrackets)-1] != oppositeBracket[b] {
					penaltyPoints += points[b]
					openedBrackets = make([]rune, 0)
					break
				} else {
					openedBrackets = openedBrackets[:len(openedBrackets)-1]
				}
			}
		}

		// incomplete line autocomplete points calculation
		if len(openedBrackets) > 0 {
			acPoints := 0
			for i := len(openedBrackets) - 1; i >= 0; i-- {
				acPoints *= 5
				acPoints += autocompletePoints[oppositeBracket[openedBrackets[i]]]
			}
			autocompletePointsList = append(autocompletePointsList, acPoints)
		}
	}
	return penaltyPoints, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	sort.Ints(autocompletePointsList)
	return autocompletePointsList[len(autocompletePointsList)/2], nil
}
