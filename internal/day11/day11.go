package day11

import (
	"strconv"
)

// Octopus is a one particular point in the octopus matrix. It has its energy
// which is consumed on flash and it also stores a boolean value if the octopus
// has already flashed in one particular step.
type octopus struct {
	energy          int
	flashedThisStep bool
}

// The 2D size of the octopus matrix
const size = 10

// This is a matrix of octopuses in the cavern.
var matrix [size][size]octopus

// The number of lines loaded.
var loadedLines = 0

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	for i, v := range data {
		// convert the string to int
		e, _ := strconv.Atoi(string(v))

		// add the octopus to the matrix
		matrix[loadedLines][i] = octopus{
			energy:          e,
			flashedThisStep: false,
		}
	}
	loadedLines++
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	flashes := 0
	for i := 0; i < 100; i++ {
		// increment all the octopuses' energy
		incrementEnergies(&matrix)

		// while there is at least one octopus that is charged (which means
		// that is has energy > 9 and hasn't flashed this step yet), keep
		// flashing these octopuses
		exists, row, col := chargedOctopusExists(&matrix)
		for exists {
			runOctopusFirework(&matrix, row, col)
			flashes++
			exists, row, col = chargedOctopusExists(&matrix)
		}

		// this is the end of this step, set all octopuses flashed status
		// to boolean false for the next step
		resetFlashes(&matrix)
	}
	return flashes, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	result := 0
	return result, nil
}

// This function checks if there is at least one octopus which is ready to
// flash (has energy > 9 and hasn't flashed yet in this step).
func chargedOctopusExists(matrix *[size][size]octopus) (exists bool, row int, col int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j].energy > 9 && !matrix[i][j].flashedThisStep {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

// This function handles one octoups which had been decided to be ready to
// flash previously. The octopus flashes (increments adjanced octopuses
// energy by 1 - in case that they haven't flashed this step - and sets
// itselves value to 0 - the value of the called octopus by row and col).
func runOctopusFirework(matrix *[size][size]octopus, row int, col int) {
	// set the octopus flashed
	matrix[row][col].flashedThisStep = true
	matrix[row][col].energy = 0

	// update the adjanced cells
	imin, imax, jmin, jmax := incrementingCellRange(row, col, size)
	for i := imin; i <= imax; i++ {
		for j := jmin; j <= jmax; j++ {
			if !matrix[row+i][col+j].flashedThisStep {
				matrix[row+i][col+j].energy++
			}
		}
	}
}

// This function resets the flash status of all octopuses in the given matrix
// to boolean false.
func resetFlashes(matrix *[size][size]octopus) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j].flashedThisStep = false
		}
	}
}

// This function increments the energy value of all octopuses in the given
// matrix by one.
func incrementEnergies(matrix *[size][size]octopus) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j].energy++
		}
	}
}

// This function determines the index change of one octopus that flashed.
// Those indexes are the minimum and maximum offset of the called octopus
// position and can be used in a loop to select all adjanced loops. The main
// purpose of this function is to exclude positions that doesn't exist (like
// position -1 in an array or size position - e.g. 5 in 5-sized array).
func incrementingCellRange(row, col, size int) (imin, imax, jmin, jmax int) {
	// minimal allowed row
	imin = -1
	if row == 0 {
		imin = 0
	}

	// maximal allowed row
	imax = 1
	if row == size-1 {
		imax = 0
	}

	// minimal allowed col
	jmin = -1
	if col == 0 {
		jmin = 0
	}

	// maximal allowed col
	jmax = 1
	if col == size-1 {
		jmax = 0
	}

	return
}
