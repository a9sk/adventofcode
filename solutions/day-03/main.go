package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	mulPattern := `mul\(\d{1,10},\d{1,10}\)`
	numPattern := `\d{1,10}`

	re := regexp.MustCompile(mulPattern)

	matches := re.FindAllString(input, -1)

	var sum int

	for i := 0; i < len(matches); i++ {
		re := regexp.MustCompile(numPattern)
		numbers := re.FindAllString(matches[i], -1)
		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])
		sum += first * second
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {

	keepPattern := `mul\(\d{1,10},\d{1,10}\)|do\(\)|don\'t\(\)`
	numPattern := `\d{1,10}`

	re := regexp.MustCompile(keepPattern)

	matches := re.FindAllString(input, -1)

	var sum int

	doMul := true

	for i := 0; i < len(matches); i++ {
		if matches[i] == "do()" {
			doMul = true
		} else if matches[i] == "don't()" {
			doMul = false
		} else if doMul {
			re := regexp.MustCompile(numPattern)
			numbers := re.FindAllString(matches[i], -1)
			first, _ := strconv.Atoi(numbers[0])
			second, _ := strconv.Atoi(numbers[1])
			sum += first * second
		}
	}

	return strconv.Itoa(sum)
}
