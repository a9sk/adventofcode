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

	numbers := strings.Split(input, "\n")

	total := make(map[[4]int]int)

	var sum = 0
	for _, number := range numbers {
		secret, _ := strconv.Atoi(number)
		last := secret % 10
		pList := make([][2]int, 0, 2000)

		for i := 0; i < 2000; i++ {
			helpSecret := secret * 64
			secret = (secret ^ helpSecret) % 16777216
			helpSecret = int(secret / 32)
			secret = (secret ^ helpSecret) % 16777216
			helpSecret = secret * 2048
			secret = (secret ^ helpSecret) % 16777216
			temp := secret % 10
			pList = append(pList, [2]int{temp - last, temp})
			last = temp
		}

		seen := make(map[[4]int]bool)

		for i := 0; i < len(pList)-4; i++ {
			var pattern [4]int
			for j := 0; j < 4; j++ {
				pattern[j] = pList[i+j][0]
			}

			val := pList[i+3][1]

			if !seen[pattern] {
				seen[pattern] = true
				total[pattern] += val
			}
		}
	}

	for _, value := range total {
		if value > sum {
			sum = value
		}
	}

	return strconv.Itoa(sum)
}
