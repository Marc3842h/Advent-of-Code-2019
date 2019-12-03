package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Marc3842h/Advent-of-Code-2019/inputs"
)

var middleX = 0
var middleY = 0

func PartA() {
	input := inputs.ReadInputsStr(3)
	result := 10000000.0

	first := strings.Split(input[0], ",")
	second := strings.Split(input[1], ",")

	// x,y
	grid := make([]Pair, 1000)

	grid = setValues(first, 1, grid)
	grid = setValues(second, 2, grid)

	length := len(grid)
	fmt.Printf("Total grid size: %d\n", length)

	for index, value1 := range grid {
		if value1.x == 0 && value1.y == 0 {
			continue
		}

		if index%5000 == 0 {
			fmt.Printf("index %d (out of %d)\n", index, length)
		}

		for _, value2 := range grid {
			if value2.x == 0 && value2.y == 0 {
				continue
			}

			if IsEqual(value1, value2) {
				//fmt.Printf("Found crossing at (%d,%d)\n", value1.x, value1.y)

				distance := math.Abs(float64(value1.x)) + math.Abs(float64(value1.y))

				if result > distance {
					result = distance
					fmt.Printf("new shortest distance %f\n", distance)
				}
			}
		}
	}

	fmt.Printf("Day3PartA: %f\n", result)
}

func PartB() {
	//
}

type Pair struct {
	x, y, id int
}

func setValues(input []string, id int, grid []Pair) []Pair {
	currX := middleX
	currY := middleY

	grid = append(grid, Pair{middleX, middleY, id})
	size := len(input)

	for index, value := range input {
		fmt.Printf("Filling first array [%d/%d]\r", index, size)

		direction := string(value[0])
		amount, _ := strconv.Atoi(value[1:])

		if direction == "R" {
			orig := currX
			currX += amount

			for i := 0; i < amount; i++ {
				grid = append(grid, Pair{orig + i, currY, id})
			}
		} else if direction == "L" {
			orig := currX
			currX -= amount

			for i := 0; i < amount; i++ {
				grid = append(grid, Pair{orig - i, currY, id})
			}
		} else if direction == "D" {
			orig := currY
			currY += amount

			for i := 0; i < amount; i++ {
				grid = append(grid, Pair{currX, orig + i, id})
			}
		} else if direction == "U" {
			orig := currY
			currY -= amount

			for i := 0; i < amount; i++ {
				grid = append(grid, Pair{currX, orig - i, id})
			}
		}
	}

	fmt.Printf("Filling first array [done]\n")
	return grid
}

func IsEqual(a Pair, b Pair) bool {
	return a.x == b.x && a.y == b.y && a.id != b.id
}
