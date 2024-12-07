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

type operation struct {
	result   int
	operands []int
}

func solvePart1(input string) string {
	lines := utils.ParseLines(input)

	var operations []operation
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		splitOperands := strings.Split(splitLine[1], " ")
		splitResult, _ := strconv.Atoi(splitLine[0])

		var operands []int
		for _, operand := range splitOperands {
			operandValue, _ := strconv.Atoi(operand)
			operands = append(operands, operandValue)
		}

		operations = append(operations, operation{
			result:   splitResult,
			operands: operands,
		})
	}

	var sum = 0
	for _, operation := range operations {
		if ok := isCalibrationCorrect(operation, 0, 0); ok {
			sum += operation.result
		}
	}

	return strconv.Itoa(sum)
}

func isCalibrationCorrect(operation operation, currentResult int, index int) bool {
	if index == len(operation.operands) {
		return operation.result == currentResult // if you finish to check the whole thing and the currentResult is the expected result return true
	}

	nextNum := operation.operands[index]

	if isCalibrationCorrect(operation, currentResult+nextNum, index+1) {
		return true
	}

	if isCalibrationCorrect(operation, currentResult*nextNum, index+1) {
		return true
	}

	return false
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
