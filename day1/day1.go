package day1

import (
	"fmt"
	"math"

	"github.com/Marc3842h/Advent-of-Code-2019/inputs"
)

func getFuel(input int) int {
	return int(math.Trunc(float64(input)/3.0)) - 2
}

func PartA() {
	input := inputs.ReadInputsInt(1)
	total := 0

	for _, number := range input {
		total += getFuel(number)
	}

	fmt.Printf("Day1PartA: %d\n", total)
}

func PartB() {
	input := inputs.ReadInputsInt(1)
	total := 0

	for _, number := range input {
		current := getFuel(number)
		module := current

		for {
			current = getFuel(current)

			if current < 1 {
				break
			}

			module += current
		}

		total += module
	}

	fmt.Printf("Day1PartB: %d\n", total)
}
