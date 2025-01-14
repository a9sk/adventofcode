package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input := strings.Join(lines, "\n")

	resultPart1 := solvePart1(string(input))
	fmt.Println("Solution Part 1:", resultPart1)

	resultPart2 := solvePart2(string(input))
	fmt.Println("Solution Part 2:", resultPart2)
}

func solvePart1(input string) string {

	var parsedInput = strings.Split(input, "\n")

	var len = len(parsedInput)

	left := make([]int, len)
	right := make([]int, len)

	for i := 0; i < len; i++ {
		line := strings.Split(parsedInput[i], "   ")

		leftVal, _ := strconv.Atoi(line[0])
		rightVal, _ := strconv.Atoi(line[1])

		left[i] = leftVal
		right[i] = rightVal
	}

	// sort the two lists
	sort.Ints(left)
	sort.Ints(right)

	// sum the differences between the two lists
	sum := 0
	for i := 0; i < len; i++ {
		sum = sum + AbsInt(left[i]-right[i])
	}

	return strconv.Itoa(sum)
}

func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func solvePart2(input string) string {

	var parsedInput = strings.Split(input, "\n")

	var len = len(parsedInput)

	left := make([]int, len)
	right := make([]int, len)

	for i := 0; i < len; i++ {
		line := strings.Split(parsedInput[i], "   ")

		leftVal, _ := strconv.Atoi(line[0])
		rightVal, _ := strconv.Atoi(line[1])

		left[i] = leftVal
		right[i] = rightVal
	}

	// sort the two lists
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len; i++ {
		times := 0
		for j := 0; j < len && right[j] <= left[i]; j++ {
			if right[j] == left[i] {
				times++
			}
		}
		sum += left[i] * times
	}

	return strconv.Itoa(sum)
}
