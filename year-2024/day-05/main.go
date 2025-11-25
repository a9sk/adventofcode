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
	parts := strings.Split(input, "\n\n")

	rules := strings.Split(parts[0], "\n")    // line by line the first part of the input
	upgrades := strings.Split(parts[1], "\n") // line by line the second part of the input

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

	parts := strings.Split(input, "\n\n")

	rules := strings.Split(parts[0], "\n")
	upgrades := strings.Split(parts[1], "\n")
	var sum = 0

	for i := 0; i < len(upgrades); i++ {
		if ok, _ := isOrdered(upgrades[i], rules); !ok { // only care about unordered lists
			orderedUpgrade := doOrder(upgrades[i], rules)
			midValue, _ := strconv.Atoi(orderedUpgrade[len(orderedUpgrade)/2])
			sum += midValue
		}
	}

	return strconv.Itoa(sum)
}

func doOrder(upgrade string, rules []string) []string {
	values := strings.Split(upgrade, ",")

	var done = false
	for {
		for _, rule := range rules {
			parts := strings.Split(rule, "|")
			first, second := parts[0], parts[1]
			firstIndex, secondIndex := -1, -1
			for i, v := range values {
				if v == first && firstIndex == -1 {
					firstIndex = i
				}
				if v == second && secondIndex == -1 {
					secondIndex = i
				}
			}
			if secondIndex != -1 && firstIndex > secondIndex {
				values[firstIndex] = second
				values[secondIndex] = first
			}

			if ok, _ := isOrdered(strings.Join(values, ","), rules); ok {
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	return values
}
