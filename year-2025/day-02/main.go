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
	var sum int

	for _, r := range strings.Split(input, ",") {
		rr := strings.Split(r, "-")
		f, _ := strconv.Atoi(rr[0])
		l, _ := strconv.Atoi(rr[1])
		for i := range l - f + 1 {
			v := strconv.Itoa(f + i)
			if v[:len(v)/2] == v[len(v)/2:] {
				// i sue f+i to avoid strconv (i guess it should be faster)
				sum += f + i
				continue
			}
		}
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
