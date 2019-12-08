package day8

import (
	"fmt"
	"strconv"

	"github.com/Marc3842h/Advent-of-Code-2019/inputs"
)

var imageWidth = 25
var imageHeight = 6

func PartA() {
	input := inputs.ReadInputStr(8)

	layers := make([]layer, 0)
	counter := 0
	layerCount := 0

	for _, c := range input {
		char := string(c)
		pixel, _ := strconv.Atoi(char)

		if counter >= imageWidth*imageHeight {
			counter = 0
			layerCount++
		}

		if !indexExists(layerCount, layers) {
			layers = append(layers, layer{contents: make([]int, 0)})
		}

		layers[layerCount].contents = append(layers[layerCount].contents, pixel)
		counter++
	}

	indexWithLeastZeros := 1000000000000
	indexZeroAmount := 10000000000000

	for index, layer := range layers {
		zeros := 0

		for _, pixel := range layer.contents {
			if pixel == 0 {
				zeros++
			}
		}

		if indexZeroAmount > zeros {
			indexZeroAmount = zeros
			indexWithLeastZeros = index
		}
	}

	layer := layers[indexWithLeastZeros]
	ones := 0
	twos := 0

	for _, pixel := range layer.contents {
		if pixel == 1 {
			ones++
		} else if pixel == 2 {
			twos++
		}
	}

	fmt.Printf("Day8PartA: %d\n", ones*twos)
}

func PartB() {
	input := inputs.ReadInputStr(8)

	layers := make([]layer, 0)
	counter := 0
	layerCount := 0

	for _, c := range input {
		char := string(c)
		pixel, _ := strconv.Atoi(char)

		if counter >= imageWidth*imageHeight {
			counter = 0
			layerCount++
		}

		if !indexExists(layerCount, layers) {
			layers = append(layers, layer{contents: make([]int, 0)})
		}

		layers[layerCount].contents = append(layers[layerCount].contents, pixel)
		counter++
	}

	image := make([]int, imageWidth*imageHeight)

	// Make the image transparent
	for index := range image {
		image[index] = 2
	}

	for _, layer := range layers {
		for index, pixel := range layer.contents {
			if pixel == 2 {
				continue
			}

			if image[index] != 2 {
				continue
			}

			image[index] = pixel
		}
	}

	fmt.Printf("Day8PartB:\n")

	pointer := 0

	for i := 0; i < imageHeight; i++ {
		slice := image[pointer : pointer+imageWidth]

		for _, pixel := range slice {
			if pixel == 0 {
				fmt.Printf(" ")
				continue
			}

			fmt.Printf("%d", pixel)
		}
		fmt.Printf("\n")

		pointer += imageWidth
	}
}

func indexExists(needleIndex int, haystack []layer) bool {
	for a := range haystack {
		if a == needleIndex {
			return true
		}
	}
	return false
}

type layer struct {
	contents []int
}
