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
		for i := 0; i < len(stringList); i++ {
			list[i], _ = strconv.Atoi(stringList[i])
		}

		var safe = true
		// fmt.Print(list)

		if list[0] > list[1] { // meaning it is increasing
			for j := 0; j < len(list)-1; j++ { // start from the same cell to check abs here as well
				if list[j] <= list[j+1] || utils.AbsInt(list[j]-list[j+1]) > 3 {
					fmt.Printf("%d unsafe\n", list)
					safe = false
					break
				}
			}
		} else if list[0] < list[1] { // and here decreasing
			for j := 0; j < len(list)-1; j++ {
				if list[j] >= list[j+1] || utils.AbsInt(list[j]-list[j+1]) > 3 {
					fmt.Printf("%d unsafe\n", list)
					safe = false
					break
				}
			}
		} else {
			safe = false
		}

		if safe {
			fmt.Printf("%d safe\n", list)
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
