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

	stones := make([]int, len(strings.Split(input, " ")))

	for i, stone := range strings.Split(input, " ") {
		stones[i], _ = strconv.Atoi(stone)
	}

	for b := 0; b < 25; b++ {
		stones = blink(stones)
	}

	return strconv.Itoa(len(stones))
}

func blink(stones []int) []int {

	var i = 0
	for i < len(stones) {

		// if the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1
		if stones[i] == 0 {
			stones[i] = 1
			i++
			continue
		}

		// if the stone is engraved with a number that has an even number of digits,
		// it is replaced by two stones.
		// the left half of the digits are engraved on the new left stone,
		// and the right half of the digits are engraved on the new right stone
		if len(strconv.Itoa(stones[i]))%2 == 0 {
			stoneValue := strconv.Itoa(stones[i])
			mid := len(stoneValue) / 2
			firstPart, _ := strconv.Atoi(stoneValue[:mid])
			secondPart, _ := strconv.Atoi(stoneValue[mid:])
			stones = append(stones[:i], append([]int{firstPart, secondPart}, stones[i+1:]...)...)
			i += 2 // so it skips the new stone added
			continue
		}

		// if none of the other rules apply, the old stone's number multiplied by 2024 is engraved on the new stone
		stones[i] = stones[i] * 2024
		i++
	}

	return stones
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
