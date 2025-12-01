package main

import (
	"bufio"
	"fmt"
	_ "math"
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

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "R") {
			rot, _ := strconv.Atoi(line[1:])
			s = s + rot
		} else if strings.HasPrefix(line, "L") {
			rot, _ := strconv.Atoi(line[1:])
			s = s - rot
		}
		if s%100 == 0 {
			c++
		}
	}

	return strconv.Itoa(c)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func solvePart2(input string) string {

	var n int
	s := 50
	d := map[byte]int{'R': 1, 'L': -1}

	for _, line := range strings.Split(input, "\n") {
		rot, _ := strconv.Atoi(line[1:])

		for range rot {
			if s += d[line[0]]; s%100 == 0 {
				n++
			}
		}
	}

	return strconv.Itoa(n)
}
