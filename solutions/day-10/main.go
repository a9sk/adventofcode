package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type point struct {
	y, x int
}

func solvePart1(input string) string {

	field := utils.ParseGrid(utils.ParseLines(input))

	// save the starting points in a slice
	startingPoints := []point{}
	for y, row := range field {
		for x, _ := range row {
			if field[y][x] == '0' {
				startingPoints = append(startingPoints, point{y: y, x: x})
			}
		}
	}

	var sum = 0
	for _, point := range startingPoints {
		sum += calculateTrailheadScore(field, point)
	}

	return strconv.Itoa(sum)
}

func calculateTrailheadScore(field [][]rune, start point) int {
	visited := make(map[point]bool)
	ninePositions := make(map[point]bool)
	queue := []point{{start.y, start.x}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}
		visited[node] = true

		if field[node.y][node.x] == '9' {
			ninePositions[node] = true
		}

		directions := []point{
			{node.y + 1, node.x},
			{node.y - 1, node.x},
			{node.y, node.x + 1},
			{node.y, node.x - 1},
		}

		for _, next := range directions {
			if isValidMove(field, node, next, visited) {
				queue = append(queue, next)
			}
		}
	}

	return len(ninePositions)
}

func isValidMove(field [][]rune, current, next point, visited map[point]bool) bool {
	if next.y < 0 || next.y >= len(field) ||
		next.x < 0 || next.x >= len(field[0]) {
		return false
	}

	if visited[next] {
		return false
	}

	return field[next.y][next.x] == field[current.y][current.x]+1

}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
