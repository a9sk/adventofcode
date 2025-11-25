package main

import (
	"bufio"
	"fmt"
	"os"
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

	resultPart1 := solvePart1(input)
	fmt.Println("Solution Part 1:", resultPart1)

	resultPart2 := solvePart2(input)
	fmt.Println("Solution Part 2:", resultPart2)
}

func solvePart1(input string) string {

	A, B, C, commands := parseInput(input)

	var program []int
	for _, val := range strings.Split(commands, ",") {
		num, _ := strconv.Atoi(val)
		program = append(program, num)
	}

	output := calculateProgram(A, B, C, program)

	return strings.Trim(strings.Replace(fmt.Sprint(output), " ", ",", -1), "[]")
}

func solvePart2(input string) string {

	_, B, C, commands := parseInput(input)

	var program []int
	for _, val := range strings.Split(commands, ",") {
		num, _ := strconv.Atoi(val)
		program = append(program, num)
	}

	A := 0
	for n := len(program) - 1; n >= 0; n-- {
		A = A << 3
		for !equalOutput(calculateProgram(A, B, C, program), program[n:]) {
			A++
		}
	}
	return strconv.Itoa(A)
}

func parseInput(input string) (int, int, int, string) {
	var A, B, C int
	var program string

	fmt.Sscanf(input, "Register A: %d", &A)
	fmt.Sscanf(input, "Register B: %d", &B)
	fmt.Sscanf(input, "Register C: %d", &C)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Program:") {
			program = strings.TrimPrefix(line, "Program: ")
			program = strings.TrimSpace(program)
		}
	}

	return A, B, C, program
}

func calculateProgram(A, B, C int, program []int) []int {
	var output []int

	for instructionPointer := 0; instructionPointer < len(program); instructionPointer += 2 {
		op, literal := program[instructionPointer], program[instructionPointer+1]

		combo := literal
		switch combo {
		case 4:
			combo = A
		case 5:
			combo = B
		case 6:
			combo = C
		}

		switch op {
		case 0:
			A >>= combo
		case 1:
			B ^= literal
		case 2:
			B = combo % 8
		case 3:
			if A != 0 {
				instructionPointer = literal - 2
			}
		case 4:
			B ^= C
		case 5:
			output = append(output, combo%8)
		case 6:
			B = A >> combo
		case 7:
			C = A >> combo
		}
	}

	return output
}

func equalOutput(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
