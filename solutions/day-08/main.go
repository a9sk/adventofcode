package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/a9sk/adventofcode/utils"
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

	lines := utils.ParseLines(input)
	height := len(lines)
	width := len(lines[0])

	antennas := make(map[rune][]struct{ x, y int })

	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				antennas[ch] = append(antennas[ch], struct{ x, y int }{x, y})
			}
		}
	}

	antinodes := make(map[struct{ x, y int }]bool)

	for _, points := range antennas {
		if len(points) < 2 {
			continue
		}

		for i := 0; i < len(points)-1; i++ {
			for j := i + 1; j < len(points); j++ {
				a1, a2 := points[i], points[j]

				distX := a2.x - a1.x
				distY := a2.y - a1.y

				antiNode1 := struct{ x, y int }{x: a1.x - distX, y: a1.y - distY}
				antiNode2 := struct{ x, y int }{x: a2.x + distX, y: a2.y + distY}

				if antiNode1.x >= 0 && antiNode1.x < width &&
					antiNode1.y >= 0 && antiNode1.y < height {
					antinodes[antiNode1] = true
				}

				if antiNode2.x >= 0 && antiNode2.x < width &&
					antiNode2.y >= 0 && antiNode2.y < height {
					antinodes[antiNode2] = true
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(antinodes))
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
