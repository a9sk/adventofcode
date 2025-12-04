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
	dir := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var s int

	ls := strings.Fields(string(input))

	g := make([][]rune, len(ls))

	for y, l := range ls {
		g[y] = make([]rune, len(ls[0]))
		for x, c := range l {
			g[y][x] = c
		}
	}

	for y, l := range ls {
		for x := range l {
			if g[y][x] != '@' {
				continue
			}
			var i int
			for _, d := range dir {
				ny, nx := y+d[0], x+d[1]
				if ny < 0 || nx < 0 || ny >= len(ls) || nx >= len(l) {
					continue
				}
				if g[ny][nx] == '@' {
					i++
				}
			}
			if i < 4 {
				s++
			}
		}
	}

	return strconv.Itoa(s)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
