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

	diskMap := strings.Split(input, "")

	var blocks []string
	var id = 0
	for j, num := range diskMap {
		numInt, _ := strconv.Atoi(num)
		for i := 0; i < numInt; i++ {
			if j%2 != 0 {
				blocks = append(blocks, ".")
				continue
			}
			blocks = append(blocks, strconv.Itoa(id))
		}
		if j%2 != 0 {
			id++
		}
	}

	iterativeMove(blocks)

	var sum = 0
	for index, value := range blocks {
		intValue, _ := strconv.Atoi(value)
		sum += intValue * index
	}

	return strconv.Itoa(sum)
}

func iterativeMove(blocks []string) {
	var lastFull = 0
	for i := len(blocks) - 1; i > lastFull; i-- {
		if blocks[i] == "." {
			continue
		}
		for j := lastFull; j < i; j++ {
			if blocks[j] != "." {
				lastFull = j
				continue
			}
			blocks[j] = blocks[i]
			blocks[i] = "."
			break
		}
	}
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
