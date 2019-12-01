package inputs

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInputsStr(day int) (strings []string) {
	file, _ := os.Open("day" + strconv.Itoa(day) + "/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return
}

func ReadInputsInt(day int) (numbers []int) {
	file, _ := os.Open("day" + strconv.Itoa(day) + "/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, result)
	}

	return
}

func ReadInputStr(day int) string {
	file, _ := os.Open("day" + strconv.Itoa(day) + "/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func ReadInputInt(day int) int {
	file, _ := os.Open("day" + strconv.Itoa(day) + "/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result, _ := strconv.Atoi(scanner.Text())
		return result
	}

	return 0
}
