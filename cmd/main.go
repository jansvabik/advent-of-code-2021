package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jansvabik/advent-of-code-2021/internal/day01"
	"github.com/jansvabik/advent-of-code-2021/internal/day02"
	"github.com/jansvabik/advent-of-code-2021/internal/day03"
	"github.com/jansvabik/advent-of-code-2021/internal/day04"
)

// Structure with all the public day functions that are going to be called
// by the main process. Each day has different implementation of those functions.
// todo: one should maybe use an interface?
type dayFunc struct {
	registerInput func(string) error
	part1         func() (interface{}, error)
	part2         func() (interface{}, error)
}

// This array holds the list of accessible days (only those that are already
// programmed and tested).
var dayFuncs = [...]dayFunc{
	{day01.RegisterInput, day01.Part1, day01.Part2},
	{day02.RegisterInput, day02.Part1, day02.Part2},
	{day03.RegisterInput, day03.Part1, day03.Part2},
	{day04.RegisterInput, day04.Part1, day04.Part2},
}

func main() {
	// check that there is a day argument
	if len(os.Args) == 1 {
		fmt.Printf("You should specify the day argument: ./program <day> < dayInput.txt\n")
		os.Exit(1)
	}

	// extract the day from program arguments
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot use your input as a day number: %v\n", os.Args[0])
		os.Exit(1)
	}

	// check that the day exists
	if len(dayFuncs) < day {
		fmt.Printf("This day is not implemented yet or just doesn't exist. Advent of code has 25 days.\n")
		os.Exit(1)
	}

	// read all data from stdin
	fmt.Printf("Type your input:\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := dayFuncs[day-1].registerInput(scanner.Text())
		if err != nil {
			fmt.Println("Your input is not valid for this day:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Input accepted.\n")

	// handle read errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Cannot load input data: %s\n", err)
	}

	// call the puzzle solver for part 1
	resultPt1, err := dayFuncs[day-1].part1()
	if err != nil {
		fmt.Printf("Cannot solve day %d, part 1: %s\n", day, err)
	}
	fmt.Printf("Day %d, part 1: %v\n", day, resultPt1)

	// call the puzzle solver for part 2
	resultPt2, err := dayFuncs[day-1].part2()
	if err != nil {
		fmt.Printf("Cannot solve day %d, part 2: %s\n", day, err)
	}
	fmt.Printf("Day %d, part 2: %v\n", day, resultPt2)
}
