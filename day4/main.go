package main

import (
	"fmt"
	"strconv"
)

const (
	lowerBound = 265275
	upperBound = 781584
)

func main() {
	part1()
	part2()
}

func part2() {
	var totalValid int
	for i := lowerBound; i < upperBound; i++ {
		if isValid(i, true) {
			totalValid++
		}
	}

	fmt.Printf("Part2 Result: %d\n", totalValid)
}

func part1() {
	var totalValid int
	for i := lowerBound; i < upperBound; i++ {
		if isValid(i, false) {
			totalValid++
		}
	}

	fmt.Printf("Part1 Result: %d\n", totalValid)
}

func isValid(n int, restrictThreeOrMore bool) bool {
	nStr := strconv.Itoa(n)

	var foundDouble bool
	var lastChar int32

	var foundLocalDouble bool
	var detectedAtLeastThree bool

	for i, curChar := range nStr {
		if i == 0 {
			lastChar = curChar
			continue
		}

		if curChar < lastChar {
			return false
		}

		if curChar == lastChar {

			// If we found the same letter, and we've detected same letter already
			if foundLocalDouble {
				detectedAtLeastThree = true
			}

			foundLocalDouble = true
		}

		if curChar != lastChar || i == (len(nStr)-1) {

			if detectedAtLeastThree && restrictThreeOrMore {
				// We don't count three or more occurrences as valid
			} else {
				foundDouble = foundDouble || foundLocalDouble
			}

			// Reset for next character
			foundLocalDouble = false
			detectedAtLeastThree = false
		}

		lastChar = curChar
	}

	return foundDouble
}
