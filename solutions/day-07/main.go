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
	lines := strings.Split(input, "\n")
	var sum int
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(parts[0])
		operands := strings.Split(parts[1], " ")

		if checkOperations(operands, testValue, []string{"+", "*", "||"}) {
			sum += testValue
		}
	}
	return strconv.Itoa(sum)
}

func checkOperations(operands []string, testValue int, operators []string) bool {
	operatorCombinations := generateOperatorCombinations(len(operands)-1, operators)

	for _, ops := range operatorCombinations {
		if evaluateExpression(operands, ops) == testValue {
			return true
		}
	}
	return false
}

func generateOperatorCombinations(numOperators int, operators []string) [][]string {
	var combinations [][]string
	var generate func(int, []string)
	generate = func(index int, currentCombination []string) {
		if index == numOperators {
			combinations = append(combinations, append([]string{}, currentCombination...))
			return
		}
		for _, operator := range operators {
			currentCombination = append(currentCombination, operator)
			generate(index+1, currentCombination)
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}
	generate(0, []string{})
	return combinations
}

func evaluateExpression(operands []string, operators []string) int {
	nums := make([]int, len(operands))
	for i, operand := range operands {
		nums[i], _ = strconv.Atoi(operand)
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		operator := operators[i-1]
		switch operator {
		case "+":
			result += nums[i]
		case "*":
			result *= nums[i]
		case "||":
			result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, nums[i]))
		}
	}
	return result
}
