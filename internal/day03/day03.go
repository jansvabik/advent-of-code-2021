package day03

import (
	"strconv"
)

// The number of bits of each input
const bits = 12

// The number of input lines processed by RegisterInput func.
var inputs []string

// The number of 0 bits on each position. Summarized for all lines.
var bits0 [bits]int

// The number of 1 bits on each position. Summarized for all lines.
var bits1 [bits]int

// RegisterInput processes each line of the input and stores it in prepared
// input variables above. It does the check of each line to match the day
// requirements.
func RegisterInput(data string) error {
	for i, v := range data {
		if v == '0' {
			bits0[i]++
		} else {
			bits1[i]++
		}
	}

	inputs = append(inputs, data)
	return nil
}

// Part 1 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the first part.
func Part1() (interface{}, error) {
	gammaRateRunes := make([]rune, bits)
	epsilonRateRunes := make([]rune, bits)
	for i := 0; i < bits; i++ {
		if bits0[i] > bits1[i] {
			gammaRateRunes[i] = '0'
			epsilonRateRunes[i] = '1'
		} else {
			gammaRateRunes[i] = '1'
			epsilonRateRunes[i] = '0'
		}
	}

	// convert the binary string number to real integer
	gammaRate, err := strconv.ParseInt(string(gammaRateRunes), 2, 0)
	if err != nil {
		return nil, err
	}

	// convert the binary string number to real integer
	epsilonRate, err := strconv.ParseInt(string(epsilonRateRunes), 2, 0)
	if err != nil {
		return nil, err
	}

	result := gammaRate * epsilonRate
	return result, nil
}

// Part 2 solves the puzzle itself and returns the result. It calls other
// internal helper functions to get the result. It solves the second part.
func Part2() (interface{}, error) {
	// last added values to the considered ones will be those values we want
	// to find and multiply in the end of the process
	var oxyRate, CO2Rate string

	// run the filtering loop for every bit (position) in the inputs
	finalConsideredOxy := inputs[:]
	finalConsideredCO2 := inputs[:]
	for bit := 0; bit < bits; bit++ {
		// temporary slices for this bit's round, will be used for filtering
		// the data in the next round
		consideredOxy := []string{}
		consideredCO2 := []string{}

		// find the most common bit in considered oxygen rate and co2 rate data
		oxyBits := map[byte]int{'0': 0, '1': 0}
		CO2Bits := map[byte]int{'0': 0, '1': 0}
		for _, v := range finalConsideredOxy {
			oxyBits[v[bit]]++
		}
		for _, v := range finalConsideredCO2 {
			CO2Bits[v[bit]]++
		}

		// determine the most and least common bits for every input
		mostCommonBit := '1'
		leastCommonBit := '0'
		if oxyBits['0'] > oxyBits['1'] {
			mostCommonBit = '0'
		}
		if CO2Bits['1'] < CO2Bits['0'] {
			leastCommonBit = '1'
		}

		// check all the inputs on the bit position
		for _, v := range finalConsideredOxy {
			if v[bit] == byte(mostCommonBit) {
				consideredOxy = append(consideredOxy, v)
				oxyRate = v
			}
		}
		for _, v := range finalConsideredCO2 {
			if v[bit] == byte(leastCommonBit) {
				consideredCO2 = append(consideredCO2, v)
				CO2Rate = v
			}
		}

		finalConsideredOxy = consideredOxy
		finalConsideredCO2 = consideredCO2
	}

	// convert the binary string number to real integer
	oxyRateNum, err := strconv.ParseInt(oxyRate, 2, 0)
	if err != nil {
		return nil, err
	}

	// convert the binary string number to real integer
	CO2RateNum, err := strconv.ParseInt(CO2Rate, 2, 0)
	if err != nil {
		return nil, err
	}

	result := oxyRateNum * CO2RateNum
	return result, nil
}
