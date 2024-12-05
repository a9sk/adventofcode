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
	rules := utils.ParseLines(strings.Split(input, "\n\n")[0])    // line by line the first part of the input
	upgrades := utils.ParseLines(strings.Split(input, "\n\n")[1]) // line by line the second part of the input
	var sum = 0

	for i := 0; i < len(upgrades); i++ {
		if ok, num := isOrdered(upgrades[i], rules); ok {

			midValue, _ := strconv.Atoi(num)
			sum += midValue
			// fmt.Printf("sum: %d\n\n", sum)
		}
	}

	return strconv.Itoa(sum)
}

// this isOrdered() function check is a very long and tedious way if a string is ordered following the rules given as input...
func isOrdered(upgrade string, rules []string) (bool, string) {
	values := strings.Split(upgrade, ",")
	// fmt.Printf("new value\n")

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		first, second := parts[0], parts[1]
		// fmt.Printf("first %s, second %s\n", first, second)
		firstIndex, secondIndex := -1, -1
		for i, v := range values {
			if v == first && firstIndex == -1 {
				firstIndex = i
			}
			if v == second && secondIndex == -1 {
				secondIndex = i
			}
		}
		// fmt.Printf("i1: %d, i2: %d\n", firstIndex, secondIndex)
		if secondIndex != -1 && firstIndex > secondIndex {
			return false, "0"
		}
	}

	return true, values[len(values)/2]
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
