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

/*
	NOTE: i could not solve todays part one, i just could not get the right output. I had to look for the solution
	on the AoC reddit and found TimFanelle's Golang implementation. Looking at how he did it, i noticed i had read the
	definition of antinode in the wrong way (i'll use the "language barrier" excuse here). After i noticed i just had to
	fix a few things and came up with the following solution for part 1 (and then for part 2 obv).
*/

func solvePart1(input string) string {

	lines := strings.Split(input, "\n")
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

	return strconv.Itoa(len(antinodes))
}

func solvePart2(input string) string {

	lines := strings.Split(input, "\n")
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

				currentNode := struct{ x, y int }{x: a2.x - distX, y: a2.y - distY}
				for currentNode.x >= 0 && currentNode.x < width &&
					currentNode.y >= 0 && currentNode.y < height {
					antinodes[currentNode] = true
					currentNode = struct{ x, y int }{x: currentNode.x - distX, y: currentNode.y - distY}
				}

				currentNode = struct{ x, y int }{x: a1.x + distX, y: a1.y + distY}
				for currentNode.x >= 0 && currentNode.x < width &&
					currentNode.y >= 0 && currentNode.y < height {
					antinodes[currentNode] = true
					currentNode = struct{ x, y int }{x: currentNode.x + distX, y: currentNode.y + distY}
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}
