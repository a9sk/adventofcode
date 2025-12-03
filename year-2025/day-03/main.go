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
	for _, l := range strings.Split(input, "\n") {

		d := 0
		for i := range len(l) - 1 {
			if l[i] > l[d] {
				d = i
			}
		}

		u := len(l) - 1
		for i := len(l) - 2; i > d; i-- {
			if l[i] > l[u] {
				u = i
			}
		}

		sum += int(l[d]-'0')*10 + int(l[u]-'0')
	}

	return strconv.Itoa(sum)

}

func solvePart2(input string) string {
	var sum int
	for _, l := range strings.Split(input, "\n") {
		var n string
		var j int // index of last greatest
		for len(n) < 12 {
			j = 0
			for i := j; i < len(l)-11+len(n); i++ {
				if l[i] > l[j] {
					j = i
				}
			}
			n = n + string(l[j])
			l = l[j+1:]
		}
		s, _ := strconv.Atoi(n)
		sum += s
	}

	return strconv.Itoa(sum)
}
