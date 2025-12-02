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
	var sum int

	for _, r := range strings.Split(input, ",") {
		rr := strings.Split(r, "-")
		f, _ := strconv.Atoi(rr[0])
		l, _ := strconv.Atoi(rr[1])
	nv:
		for i := range l - f + 1 {
			v := strconv.Itoa(f + i)
		nn:
			for j := len(v) / 2; j >= 1; j-- {
				if len(v)%j != 0 {
					continue
				}
				var k int
				var c []string
				for k = 0; k < len(v); k += j {
					c = append(c, string(v[k:min(k+j, len(v))]))
				}

				for a := range len(c) - 1 {
					if c[a+1] != c[a] {
						continue nn
					}
				}

				t, _ := strconv.Atoi(v)
				sum += t
				continue nv
			}
		}
	}

	return strconv.Itoa(sum)
}
