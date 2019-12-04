package day4

import (
	"fmt"
	"strconv"
)

// Input
var inputLowerLimit = 353096
var inputUpperLimit = 843212

func PartA() {
	result := 0

	for i := inputLowerLimit; i <= inputUpperLimit; i++ {
		if !containsIdenticalAdjacentDigits(i) || !containsNonDecreasingDigits(i) {
			continue
		}

		result++
	}

	fmt.Printf("Day4PartA: %d\n", result)
}

func PartB() {
	result := 0

	for i := inputLowerLimit; i <= inputUpperLimit; i++ {
		if !containsIdenticalAdjacentDigitsNoGroup(i) || !containsNonDecreasingDigits(i) {
			continue
		}

		result++
	}

	fmt.Printf("Day4PartB: %d\n", result)
}

func containsIdenticalAdjacentDigits(num int) bool {
	str := strconv.Itoa(num)

	for index, char := range str {
		if index == (len(str) - 1) {
			continue
		}

		c := string(char)
		next := string(str[index+1])

		if c == next {
			return true
		}
	}

	return false
}
func containsIdenticalAdjacentDigitsNoGroup(num int) bool {
	str := strconv.Itoa(num)
	skippingCount := 0

	for index, char := range str {
		if index == (len(str) - 1) {
			continue
		}

		if skippingCount > 0 {
			skippingCount--
			continue
		}

		c := string(char)
		next := string(str[index+1])

		if c == next {
			occurs := 1

			// Check that this is not a group
			for i := index + 2; i < len(str); i++ {
				character := string(str[i])

				if character == c {
					occurs++
				}
			}

			if occurs == 1 {
				return true
			}

			skippingCount = occurs
		}
	}

	return false
}

func containsNonDecreasingDigits(num int) bool {
	str := strconv.Itoa(num)

	for index, char := range str {
		if index == (len(str) - 1) {
			continue
		}

		current, _ := strconv.Atoi(string(char))
		next, _ := strconv.Atoi(string(str[index+1]))

		if current > next {
			return false
		}
	}

	return true
}
