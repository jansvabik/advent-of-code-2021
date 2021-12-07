package day04

import (
	"strconv"
	"strings"
)

// Type board is a data type for storing one particular bingo board. The board
// consists of 5 rows and 5 cols which are filled with numbers.
type board [5][5]int

// This holds the data about all bingo game boards.
var boards = []board{}

// This variable stores the list of numbers to be called when the bingo game
// begins. They are going to be called in the same order as in the slice.
var numbersToCall = []int{}

// This variable stores the list of numbers that has already been called.
var calledNumbers = map[int]struct{}{}

// This map holds the number of boards that has already won. It is used just
// for checking if one particular board already won or not.
var wonBoards = map[int]struct{}{}

// Default board for creating new board. It is filled with -1 values for
// determining if the particular row/value has already been initiated.
var defaultBoard = board{
	[5]int{-1, -1, -1, -1, -1},
	[5]int{-1, -1, -1, -1, -1},
	[5]int{-1, -1, -1, -1, -1},
	[5]int{-1, -1, -1, -1, -1},
	[5]int{-1, -1, -1, -1, -1},
}

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	// first initialization consists of initializing the called numbers array
	// and creating first board in the slices of them
	if len(numbersToCall) == 0 {
		initnumbersToCall(data)
		boards = append(boards, defaultBoard)
	}

	// determine the number of rows of the last board stored
	newRowIndex := -1
	for i, b := range boards[len(boards)-1] {
		if b[0] == -1 {
			newRowIndex = i
			break
		}
	}

	// the last board has been filled in so we are going to add a new one
	// and we will add the first row of numbers to it
	if newRowIndex == -1 {
		boards = append(boards, defaultBoard)
		newRowIndex = 0
	}

	// if there is not 5 substrings, move on because we have to find a line
	// containing 5 numbers delimited by whitespace for using the input as
	// a list of new values for the board that is being created
	split := strings.Fields(data)
	if len(split) != 5 {
		return nil
	}

	// let's create an array of 5 integers from the current input, this will
	// be used as a new board row
	boardRow := [5]int{}
	for i, v := range split {
		num, _ := strconv.Atoi(v)
		boardRow[i] = num
	}

	// add the row to the last board
	boards[len(boards)-1][newRowIndex] = boardRow
	return nil
}

// This functions takes a list of numbers delimited by a comma which is then
// split, converted to integers and added to the list of called numbers.
func initnumbersToCall(list string) {
	for _, v := range strings.Split(list, ",") {
		num, _ := strconv.Atoi(v)
		numbersToCall = append(numbersToCall, num)
	}
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	for _, calledNumber := range numbersToCall {
		calledNumbers[calledNumber] = struct{}{}
		for boardIndex, board := range boards {
			for row := range board {
				for col := range board[row] {
					if calledNumber == board[row][col] {
						won, points := checkBoard(&board, row, col)
						if won {
							wonBoards[boardIndex] = struct{}{}
							return points, nil
						}
					}
				}
			}
		}
	}
	return -1, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	var lastWonBoardPoints int
	for _, calledNumber := range numbersToCall[len(calledNumbers):] { // starting at the part 1 end point
		calledNumbers[calledNumber] = struct{}{}
		for boardIndex, board := range boards {
			for row := range board {
				for col := range board[row] {
					if calledNumber == board[row][col] {
						won, points := checkBoard(&board, row, col)
						if _, wonInPast := wonBoards[boardIndex]; !wonInPast && won {
							wonBoards[boardIndex] = struct{}{}
							lastWonBoardPoints = points
						}
					}
				}
			}
		}
	}
	return lastWonBoardPoints, nil
}

// This function checks whether the board has won right now and returns the
// boolean value of the decision and the number of points.
func checkBoard(board *board, row int, col int) (won bool, points int) {
	// check all rows on the column position
	fullRow := true
	for i := 0; i < 5; i++ {
		if called := isCalledNumber(board[i][col]); !called {
			fullRow = false
			break
		}
	}

	// check all cols on the row position
	fullCol := true
	for i := 0; i < 5; i++ {
		if called := isCalledNumber(board[row][i]); !called {
			fullCol = false
			break
		}
	}

	// check whether the fullrows or fullcols arrays are filled with booleans only
	return fullRow || fullCol, calculateBoardPoints(board, board[row][col])
}

// This function calculates the number of points of the specified board
// depending on the lastly called number which multiples the points.
func calculateBoardPoints(board *board, calledNumber int) int {
	points := 0
	for row := range board {
		for _, num := range board[row] {
			if called := isCalledNumber(num); !called {
				points += num
			}
		}
	}
	return points * calledNumber
}

// This function checks whether the specified number has already been called.
// It returns boolean true if the number has been called and false if not.
func isCalledNumber(n int) bool {
	_, called := calledNumbers[n]
	return called
}
