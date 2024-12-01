package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/a9sk/adventofcode/utils"
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

	var parsedInput = utils.ParseLines(input)

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
		sum = sum + utils.AbsInt(left[i]-right[i])
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
