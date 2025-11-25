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

	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	groups := findGroups(grid)
	var sum = 0
	for _, group := range groups {
		sum += group.area * group.perimeter
	}

	// this was my initial approach but then i found the "same plant but different region" problem
	// plot := make(map[struct{ x, y int }]struct {
	// 	plant     rune
	// 	perimeter int
	// }, len(grid)*len(grid[0]))

	// for y, row := range grid {
	// 	for x, _ := range row {
	// 		plant, perimeter := checkPlot(grid, x, y)
	// 		plot[struct{ x, y int }{x, y}] = struct {
	// 			plant     rune
	// 			perimeter int
	// 		}{plant, perimeter}
	// 	}
	// }

	return strconv.Itoa(sum)
}

func findGroups(grid [][]rune) map[int]struct {
	plant     rune
	area      int
	perimeter int
} {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	groups := make(map[int]struct {
		plant     rune
		area      int
		perimeter int
	})

	groupID := 0

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var floodFill = func(x, y int, plant rune) (int, int) {
		stack := [][2]int{{x, y}}
		visited[y][x] = true
		perimeter := 0
		area := 1

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currentX, currentY := current[0], current[1]

			for _, d := range directions {
				nextX, nextY := currentX+d[0], currentY+d[1]

				if nextX < 0 || nextY < 0 || nextX >= cols || nextY >= rows || grid[nextY][nextX] != plant {
					perimeter++
				} else if !visited[nextY][nextX] {
					visited[nextY][nextX] = true
					area++
					stack = append(stack, [2]int{nextX, nextY})
				}
			}
		}

		return perimeter, area
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if !visited[y][x] {
				plant := grid[y][x]
				groupID++
				perimeter, area := floodFill(x, y, plant)
				groups[groupID] = struct {
					plant     rune
					area      int
					perimeter int
				}{plant, area, perimeter}
			}
		}
	}

	return groups
}

type Point struct {
	x, y int
}

// this solution is inspired by many implementations i found on the AoC reddit page
func solvePart2(input string) string {

	grid := map[Point]rune{}
	for y, line := range strings.Split(input, "\n") {
		for x, rune := range line {
			grid[Point{x, y}] = rune
		}
	}

	seen := map[Point]bool{}
	var sum = 0

	for plot := range grid {
		if seen[plot] {
			continue
		}
		seen[plot] = true

		area := 1
		sides := 0
		queue := []Point{plot}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			for _, direction := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				neighbor := Point{x: current.x + direction.x, y: current.y + direction.y}
				if grid[neighbor] != grid[current] {
					r := Point{x: current.x - direction.y, y: current.y + direction.x}
					if grid[r] != grid[current] || grid[Point{x: r.x + direction.x, y: r.y + direction.y}] == grid[current] {
						sides++
					}
				} else if !seen[neighbor] {
					seen[neighbor] = true
					queue = append(queue, neighbor)
					area++
				}
			}
		}
		sum += area * sides
	}

	return strconv.Itoa(sum)

}
