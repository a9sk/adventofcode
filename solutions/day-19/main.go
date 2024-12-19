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

	var towels []string
	// ALSO towels = append(towels, strings.Split(strings.Split(input, "\n\n")[0], ", ")...)

	for _, line := range strings.Split(strings.Split(input, "\n\n")[0], ", ") {
		towels = append(towels, line)
	}

	var design []string
	// ALSO: design = append(design, strings.Split(strings.Split(input, "\n\n")[1], "\n")...)

	for _, line := range strings.Split(strings.Split(input, "\n\n")[1], "\n") {
		design = append(design, line)
	}

	var checkDesign func(string) bool
	checkDesign = func(d string) bool {
		if d == "" {
			return true
		}
		for _, t := range towels {
			if strings.HasPrefix(d, t) {
				return checkDesign(d[len(t):])
			}
		}

		return false
	}

	var sum = 0
	// check for every design if all the message can be covered using the toweles
	for _, d := range design {
		if checkDesign(d) {
			sum++
		}
	}

	return strconv.Itoa(sum)

}

func solvePart2(input string) string {

	var towels []string
	towels = append(towels, strings.Split(strings.Split(input, "\n\n")[0], ", ")...)

	var design []string
	design = append(design, strings.Split(strings.Split(input, "\n\n")[1], "\n")...)

	cache := map[string]int{}

	var checkDesign func(string) int
	checkDesign = func(d string) (sum int) {
		if sum, ok := cache[d]; ok {
			return sum
		}

		// ok, this is a very neat trick i saw on the AoC reddit page
		// what it does is it always updates the cache when the recursive function's recursion ends
		defer func() { cache[d] = sum }()

		if d == "" {
			return 1
		}
		for _, t := range towels {
			if strings.HasPrefix(d, t) {
				sum += checkDesign(d[len(t):])
			}
		}
		return sum
	}

	var sum = 0

	for _, s := range design {
		sum += checkDesign(s)
	}

	return strconv.Itoa(sum)
}
