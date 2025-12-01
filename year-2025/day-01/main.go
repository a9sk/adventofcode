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

	resultPart1 := solvePart1(string(input))
	fmt.Println("Solution Part 1:", resultPart1)

	resultPart2 := solvePart2(string(input))
	fmt.Println("Solution Part 2:", resultPart2)
}

func solvePart1(input string) string {

	s := 50
	var c int
	// var r []int
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "R") {
			rot, _ := strconv.Atoi(line[1:])
			// r = append(r, rot)
			//	fmt.Println(rot)
			s = mod(s+rot, 100)
		} else if strings.HasPrefix(line, "L") {
			rot, _ := strconv.Atoi(line[1:])
			// rot = -rot
			// r = append(r, rot)
			// fmt.Println(rot)
			s = mod(s-rot, 100)
		}
		if s == 0 {
			c++
		}
	}

	return strconv.Itoa(c)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
