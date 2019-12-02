package day2

import (
	"fmt"
	"github.com/Marc3842h/Advent-of-Code-2019/inputs"
	"strconv"
	"strings"
)

func PartA() {
	input := inputs.ReadInputStr(2)
	s := strings.Split(input, ",")

	s[1] = "12"
	s[2] = "2"

	fmt.Printf("Day2PartA: %s\n", calcInstructions(s)[0])
}

func PartB() {
	input := inputs.ReadInputStr(2)
	s := strings.Split(input, ",")

iter:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			s[1] = strconv.Itoa(i)
			s[2] = strconv.Itoa(j)

			s = calcInstructions(s)

			if s[0] == "19690720" {
				noun, _ := strconv.Atoi(s[1])
				verb, _ := strconv.Atoi(s[2])

				fmt.Printf("Day2PartB: %d\n", 100*noun+verb)
				break iter
			}

			s = strings.Split(input, ",")
		}
	}
}

func calcInstructions(s []string) []string {
	pos := 0

	for {
		c := s[pos]
		lhsPos, _ := strconv.Atoi(s[pos+1])
		lhs, _ := strconv.Atoi(s[lhsPos])
		rhsPos, _ := strconv.Atoi(s[pos+2])
		rhs, _ := strconv.Atoi(s[rhsPos])
		store, _ := strconv.Atoi(s[pos+3])

		if c == "1" {
			s[store] = strconv.Itoa(lhs + rhs)
		} else if c == "2" {
			s[store] = strconv.Itoa(lhs * rhs)
		} else if c == "99" {
			break
		}

		pos += 4
	}

	return s
}
