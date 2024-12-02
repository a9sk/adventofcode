package main

import (
	"bufio"
	"fmt"
	"os"
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
	safeCount := 0

	for i := 0; i < len(parsedInput); i++ {
		stringList := utils.ParseSpace(parsedInput[i])

		list := make([]int, len(stringList))

		// it is easier if we use a list of integers, mainly for the jump which can be of maximum 3
		for j := 0; j < len(stringList); j++ {
			list[j], _ = strconv.Atoi(stringList[j])
		}

		if safe := isListSafe(list); safe {
			fmt.Printf("%d safe\n", list)
			safeCount++
		} else {
			fmt.Printf("%d unsafe\n", list)
		}
	}

	return strconv.Itoa(safeCount)
}

func solvePart2(input string) string {
	parsedInput := utils.ParseLines(input)
	safeCount := 0

	for i := 0; i < len(parsedInput); i++ {
		stringList := utils.ParseSpace((parsedInput[i]))

		list := make([]int, len(stringList))

		for j := 0; j < len(stringList); j++ {
			list[j], _ = strconv.Atoi(stringList[j])
		}

		if isListSafe(list) {
			fmt.Printf("%d safe\n", list)
			safeCount++
			continue
		}

		// skipping each element
		for k := 0; k < len(list); k++ {
			if isListSafe(removeIndex(list, k)) {
				fmt.Printf("%d safe\n", list)
				safeCount++
				break
			}
		}
		fmt.Printf("%d unsafe\n", list)
	}

	return strconv.Itoa(safeCount)
}

func removeIndex(list []int, index int) []int {
	newList := make([]int, 0, len(list)-1)
	newList = append(newList, list[:index]...)
	return append(newList, list[index+1:]...)
}

func isListSafe(list []int) bool {
	if len(list) <= 2 { // list with 2 or less elements is always safe (as you can skip the second if they are the same)
		return true
	}

	increasing := list[0] < list[1]
	for i := 1; i < len(list); i++ {
		diff := list[i] - list[i-1]
		if increasing {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || -diff > 3 {
				return false
			}
		}
	}

	return true
}
