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

	cache := make(map[string]int)

	totalLength := 0
	for _, stone := range stones {
		totalLength += blinkWithCache(stone, 25, cache)
	}

	return strconv.Itoa(totalLength)

	// stones := make([]int, len(strings.Split(input, " ")))

	// for i, stone := range strings.Split(input, " ") {
	// 	stones[i], _ = strconv.Atoi(stone)
	// }

	// for b := 0; b < 25; b++ {
	// 	stones = blink(stones)
	// }

	// return strconv.Itoa(len(stones))
}

// func blink(stones []int) []int {

// 	var i = 0
// 	for i < len(stones) {

// 		// if the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1
// 		if stones[i] == 0 {
// 			stones[i] = 1
// 			i++
// 			continue
// 		}

// 		// if the stone is engraved with a number that has an even number of digits,
// 		// it is replaced by two stones.
// 		// the left half of the digits are engraved on the new left stone,
// 		// and the right half of the digits are engraved on the new right stone
// 		if len(strconv.Itoa(stones[i]))%2 == 0 {
// 			stoneValue := strconv.Itoa(stones[i])
// 			mid := len(stoneValue) / 2
// 			firstPart, _ := strconv.Atoi(stoneValue[:mid])
// 			secondPart, _ := strconv.Atoi(stoneValue[mid:])
// 			stones = append(stones[:i], append([]int{firstPart, secondPart}, stones[i+1:]...)...)
// 			i += 2 // so it skips the new stone added
// 			continue
// 		}

// 		// if none of the other rules apply, the old stone's number multiplied by 2024 is engraved on the new stone
// 		stones[i] = stones[i] * 2024
// 		i++
// 	}

// 	return stones
// }

func solvePart2(input string) string {

	stones := make([]int, len(strings.Split(input, " ")))

	for i, stone := range strings.Split(input, " ") {
		stones[i], _ = strconv.Atoi(stone)
	}

	cache := make(map[string]int)

	totalLength := 0
	for _, stone := range stones {
		// using cache i waaaaaay faster that "bruteforcing" it normally
		totalLength += blinkWithCache(stone, 75, cache)
	}

	return strconv.Itoa(totalLength)
}

func blinkWithCache(stone int, iterations int, cache map[string]int) int {

	cacheKey := fmt.Sprintf("%d-%d", stone, iterations)

	if cachedResult, exists := cache[cacheKey]; exists {
		return cachedResult
	}

	if iterations == 0 {
		return 1
	}

	var result int
	if stone == 0 {
		// if the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1
		result = blinkWithCache(1, iterations-1, cache)
	} else {
		stoneValue := strconv.Itoa(stone)
		if len(stoneValue)%2 == 0 {
			// if the stone is engraved with a number that has an even number of digits,
			// it is replaced by two stones.
			// the left half of the digits are engraved on the new left stone,
			// and the right half of the digits are engraved on the new right stone
			mid := len(stoneValue) / 2
			firstPart, _ := strconv.Atoi(stoneValue[:mid])
			secondPart, _ := strconv.Atoi(stoneValue[mid:])
			result = blinkWithCache(firstPart, iterations-1, cache) + blinkWithCache(secondPart, iterations-1, cache)
		} else {
			// if none of the other rules apply, the old stone's number multiplied by 2024 is engraved on the new stone
			newStone := stone * 2024
			result = blinkWithCache(newStone, iterations-1, cache)
		}
	}

	cache[cacheKey] = result
	return result
}
