package day08

import (
	"sort"
	"strings"
)

// This structure is different for each day and should be optimized for storing
// the input data in a way achieving simple puzzle solving afterwards.
type line struct {
	sortedSegments []string
	displayValue   string
	digitSegments  map[int]string
}

// Variable input obviously contains the day input which one should pass to the
// program using stdin pipe in console (./program <day> < input.txt). The input
// string is already processed by this day function (RegisterInput).
var lines []line

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	split := strings.Split(data, " | ")
	digitList := split[0]

	// normalize all the strings used for determinatio of the display segments
	digits := strings.Split(digitList, " ")
	for i, d := range digits {
		s := strings.Split(d, "")
		sort.Strings(s)
		digits[i] = strings.Join(s, "")
	}

	// add the data to the list of input values
	lines = append(lines, line{
		sortedSegments: digits,
		displayValue:   split[1],
		digitSegments:  map[int]string{},
	})
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	// find digit 1 segments
	for lineIndex, l := range lines {
		for _, s := range l.sortedSegments {
			if len(s) == 2 {
				lines[lineIndex].digitSegments[1] = s
			}
		}
	}

	result := 0
	return result, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	result := 0
	return result, nil
}
