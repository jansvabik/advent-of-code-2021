package day07

import (
	"sort"
	"strconv"
	"strings"
)

// This variable holds the list of all positions. After laoding the positions,
// the slice is sorted from the lowest position to the highest one. Those are
// the current/default/starting positions of the crabs.
var positions []int

// Memoization map is used for storing fuel costs (value) depending on the
// position change (key). It is used in the second part of the puzzle.
var memo = map[int]int{}

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	for _, counter := range strings.Split(data, ",") {
		counterInt, _ := strconv.Atoi(counter)
		positions = append(positions, counterInt)
	}
	sort.Ints(positions)
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
// ! Not all of the inputs have to work if the middle value is not the
// ! right target value, however it worked for two different inputs.
func Part1() (interface{}, error) {
	mid := positions[len(positions)/2]
	fuel := 0
	for _, p := range positions {
		fuel += abs(p - mid)
	}
	return fuel, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	// store the middle value and the minimal fuel found which is set to the
	// maximal int value in Golang in the beggining for comparison with newly
	// found numbers in the loop
	mid := positions[len(positions)/2]
	minFuelFound := int(^uint(0) >> 1)

	// try to get the most efficient move/position for crabs to go to
	for i := 0; i < len(positions)/2; i++ {
		fuelCostLowerPos, fuelCostHigherPos := 0, 0
		for _, p := range positions {
			posHigh := abs(p - (mid + i))
			posLow := abs(p - (mid - i))
			fuelCostLowerPos += calculateRealFuel(posLow, memo)
			fuelCostHigherPos += calculateRealFuel(posHigh, memo)
		}

		// determine the position change with lower fuel cost
		if fuelCostLowerPos < minFuelFound {
			minFuelFound = fuelCostLowerPos
		}
		if fuelCostHigherPos < minFuelFound {
			minFuelFound = fuelCostHigherPos
		}
	}

	return minFuelFound, nil
}

// This function calculates the dynamic fuel for part 2 of the puzzle. The fuel
// is calculated as a sum of all previous numbers (n-1)+(n-2)+(n-N)+...+(n-n).
// It uses memoization for storing values which could be accessed multiple times.
func calculateRealFuel(n int, memo map[int]int) int {
	// get the number from memoization map if already saved
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	// handle base cases
	if n <= 1 {
		memo[1] = 1
		return 1
	}

	// calculate the result and store it in memo map
	result := n + calculateRealFuel(n-1, memo)
	memo[n] = result
	return result
}

// This function calculate the absolute value of the giben number param. It
// is more efficient because we don't need to convert between floats and
// integers as while using math.Abs() function.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
