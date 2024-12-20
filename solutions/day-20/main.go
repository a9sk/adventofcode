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

	resultPart1 := solvePart1(input)
	fmt.Println("Solution Part 1:", resultPart1)

	resultPart2 := solvePart2(input)
	fmt.Println("Solution Part 2:", resultPart2)
}

type Point struct {
	x, y int
}

func solvePart1(input string) string {

	field := map[Point]rune{}
	start := Point{}
	for y, line := range strings.Fields(input) {
		for x, cell := range line {
			point := Point{x: x, y: y}
			field[point] = cell
			if cell == 'S' {
				start = point
			}
		}
	}

	distances := adventOfBFS(field, start)

	var sum = 0
	for pointOne := range distances {
		for pointTwo := range distances {
			distance := abs(pointTwo.x-pointOne.x) + abs(pointTwo.y-pointOne.y)
			if distance <= 2 && distances[pointTwo] >= distances[pointOne]+distance+100 {
				sum++
			}
		}
	}
	return strconv.Itoa(sum)
}

func adventOfBFS(grid map[Point]rune, start Point) map[Point]int {
	queue := []Point{start}
	distances := map[Point]int{start: 0}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := Point{x: p.x + d.x, y: p.y + d.y}
			if _, ok := distances[n]; !ok && grid[n] != '#' {
				queue = append(queue, n)
				distances[n] = distances[p] + 1
			}
		}
	}

	return distances
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func solvePart2(input string) string {

	return "Solution for part 2 not implemented"
}
