package day05

import (
	"strconv"
	"strings"
)

// This structure represents one particular line.
type line struct {
	start             [2]int
	end               [2]int
	isDiagonal        bool
	isReverseDiagonal bool
}

// This variable contains the list of lines loaded from the input file.
var lines []line

// Variables resX and resY represents the maximum number found during input
// loading on the concrete coordinate.
var resX, resY int

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	// split the string to get the start coordinates and the end coordinates
	split := strings.Split(data, " -> ")
	start := strings.Split(split[0], ",")
	end := strings.Split(split[1], ",")

	// convert the strings to integers, we are going to use it as a slice or
	// an array index
	startX, _ := strconv.Atoi(start[0])
	startY, _ := strconv.Atoi(start[1])
	endX, _ := strconv.Atoi(end[0])
	endY, _ := strconv.Atoi(end[1])

	// update the maximal resolution of the diagram depending on the maximal
	// coordinates found during the loading
	resX, resY = updateResolution(startX, startY, endX, endY)

	// diagonal determination
	isDiagonal := startX != endX && startY != endY
	isReverseDiagonal := (startX > endX && startY < endY) || (startX < endX && startY > endY)

	// determine minimal and maximal x and y to properly set up the start and end
	sx, sy, ex, ey := startX, startY, endX, endY
	if !isReverseDiagonal {
		sx, ex = normalizeVector(startX, endX)
		sy, ey = normalizeVector(startY, endY)
	}

	// store the line in the slice of all loaded lines
	lines = append(lines, line{
		start:             [2]int{sx, sy},
		end:               [2]int{ex, ey},
		isDiagonal:        isDiagonal,
		isReverseDiagonal: isReverseDiagonal,
	})
	return nil
}

// This function returns the biggest number for each coordinate. It compares
// both the provided x and y values and the stored x and y values.
func updateResolution(sx, sy, ex, ey int) (x int, y int) {
	var mx, my int

	// x resolution comparison
	if sx > ex {
		mx = sx
	} else {
		mx = ex
	}

	// y resolution comparison
	if sy > ey {
		my = sy
	} else {
		my = ey
	}

	// comparing with already stored values
	if resX > mx {
		mx = resX
	}
	if resY > my {
		my = resY
	}
	return mx, my
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	// prepare an empty diagram
	diagram := make([][]int, resY+1)
	for i := range diagram {
		diagram[i] = make([]int, resX+1)
	}

	// process all non-diagonal lines (only horizonal or vertical lines)
	for _, line := range lines {
		if !line.isDiagonal {
			for y := line.start[1]; y <= line.end[1]; y++ {
				for x := line.start[0]; x <= line.end[0]; x++ {
					diagram[y][x]++
				}
			}
		}
	}

	// return the number of overlapping points in the diagram
	return calculateOverlappingPlaces(diagram), nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	// prepare an empty diagram
	diagram := make([][]int, resY+1)
	for i := range diagram {
		diagram[i] = make([]int, resX+1)
	}

	// process all non-diagonal lines (only horizonal or vertical lines)
	for _, line := range lines {
		if !line.isReverseDiagonal {
			handleIncreasingLine(diagram, &line)
		} else {
			handleDecreasingLine(diagram, &line)
		}
	}

	// return the number of overlapping points in the diagram
	return calculateOverlappingPlaces(diagram), nil
}

// This function handles increasing lines (which means lines where both x and
// y are increasing - can be normalized by normalizeVector function) and
// increases the number in the diagram displaying the line overlapping.
func handleIncreasingLine(diagram [][]int, line *line) {
	for y := line.start[1]; y <= line.end[1]; y++ {
		for x := line.start[0]; x <= line.end[0]; x++ {
			diagram[y][x]++
			if line.isDiagonal {
				y++
			}
		}
	}
}

// This function handles lines and increases diagram point values in cases when
// the vector is decreasing in only one "vector" (x or y but not both of them).
func handleDecreasingLine(diagram [][]int, line *line) {
	if line.start[0] > line.end[0] {
		x, y := line.start[0], line.start[1]
		for x >= line.end[0] && y <= line.end[1] {
			diagram[y][x]++
			x--
			y++
		}
	} else {
		x, y := line.start[0], line.start[1]
		for x <= line.end[0] && y >= line.end[1] {
			diagram[y][x]++
			x++
			y--
		}
	}
}

// This function calculates the number of points in the diagram where more than
// one line went through and made a greater number than 0 tho.
func calculateOverlappingPlaces(diagram [][]int) int {
	count := 0
	for _, row := range diagram {
		for _, val := range row {
			if val > 1 {
				count++
			}
		}
	}
	return count
}

// This function is used for simplifying increasing vectors. It can be used
// only for cases when both x and y are increasing or both are decreasing.
func normalizeVector(start, end int) (min, max int) {
	if start < end {
		return start, end
	}
	return end, start
}
