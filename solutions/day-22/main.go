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

	numbers := strings.Split(input, "\n")

	var sum = 0
	for _, number := range numbers {
		secret, _ := strconv.Atoi(number)
		for i := 0; i < 2000; i++ {
			helpSecret := secret * 64
			secret = (secret ^ helpSecret) % 16777216
			helpSecret = int(secret / 32)
			secret = (secret ^ helpSecret) % 16777216
			helpSecret = secret * 2048
			secret = (secret ^ helpSecret) % 16777216
		}
		sum += secret
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
