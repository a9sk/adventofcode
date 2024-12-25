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
	var locks, keys [][5]int

	for _, schematic := range strings.Split(input, "\n\n") {
		if schematic[0] == '.' { // it is a simple string, not parsed in lines still
			keys = append(keys, parseSchema(schematic, true))
			continue
		}
		locks = append(locks, parseSchema(schematic, false))
	}

	var sum = 0
	for _, lock := range locks {
		for _, key := range keys {
			if key[0]+lock[0] <= 7 && key[1]+lock[1] <= 7 && key[2]+lock[2] <= 7 && key[3]+lock[3] <= 7 && key[4]+lock[4] <= 7 {
				sum++
			}
		}
	}

	return strconv.Itoa(sum)
}

// func matches(k, l [5]int) bool {
// 	return k[0]+l[0] == 6 && k[1]+l[1] == 6 && k[2]+l[2] == 6 && k[3]+l[3] == 6 && k[4]+l[4] == 6
// }

func parseSchema(schema string, isKey bool) [5]int {
	frequencies := [5]int{-1, -1, -1, -1, -1}

	for i, row := range strings.Split(schema, "\n") {
		for j, cell := range row {
			if frequencies[j] != -1 {
				continue
			}
			if isKey && cell == '#' {
				frequencies[j] = 7 - i
			} else if !isKey && cell == '.' {
				frequencies[j] = i
			}
		}
	}

	return frequencies
}

func solvePart2(input string) string {

	return "No solution needed for part 2, GGs"
}
