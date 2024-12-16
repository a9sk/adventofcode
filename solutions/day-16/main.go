package main

import (
	"bufio"
	"fmt"
	"math"
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

type directedPoint struct {
	point     point
	direction direction
}

type direction rune

const (
	north direction = '^'
	east  direction = '>'
	south direction = 'v'
	west  direction = '<'
)

type point struct {
	x int
	y int
}

type routeState struct {
	reindeer directedPoint
	steps    []directedPoint
	score    int
}

func solvePart1(input string) string {

	grid := utils.ParseGrid(utils.ParseLines(input))

	var reindeer directedPoint
	var end point

	maze := make([][]rune, len(grid))
	for y := range grid {
		maze[y] = make([]rune, len(grid[y]))
		for x := range grid[y] {
			switch grid[y][x] {
			case 'S':
				reindeer = directedPoint{point{x, y}, east}
				maze[y][x] = '.'
			case 'E':
				end = point{x, y}
				maze[y][x] = '.'
			default:
				maze[y][x] = grid[y][x]
			}
		}
	}

	bestScore := findShortestPath(reindeer, end, maze)

	return strconv.Itoa(bestScore)
}

func findShortestPath(reindeer directedPoint, end point, maze [][]rune) int {
	bestScore := math.MaxInt
	queue := []routeState{{reindeer, []directedPoint{reindeer}, 0}}
	visited := make(map[directedPoint]int)

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.steps) > 10000 {
			continue
		}

		if state.score > bestScore {
			continue
		}

		if state.reindeer.point == end {
			if state.score <= bestScore {
				bestScore = state.score
				continue
			}
		}

		for _, candidate := range getOffsets(state.reindeer) {
			if maze[candidate.point.y][candidate.point.x] == '.' {
				score := state.score + 1
				if state.reindeer.direction != candidate.direction {
					score += 1000
				}

				if previous, found := visited[candidate]; found {
					if previous < score {
						continue
					}
				}
				visited[candidate] = score

				newSteps := make([]directedPoint, len(state.steps))
				copy(newSteps, state.steps)

				queue = append(queue, routeState{
					reindeer: candidate,
					steps:    append(newSteps, candidate),
					score:    score})
			}
		}
	}

	return bestScore
}

func getOffsets(reindeer directedPoint) []directedPoint {
	n := directedPoint{point{x: reindeer.point.x + 0, y: reindeer.point.y - 1}, north}
	e := directedPoint{point{x: reindeer.point.x + 1, y: reindeer.point.y + 0}, east}
	s := directedPoint{point{x: reindeer.point.x + 0, y: reindeer.point.y + 1}, south}
	w := directedPoint{point{x: reindeer.point.x - 1, y: reindeer.point.y + 0}, west}

	switch reindeer.direction {
	case north:
		return []directedPoint{n, e, w}
	case east:
		return []directedPoint{e, s, n}
	case south:
		return []directedPoint{s, w, e}
	case west:
		return []directedPoint{w, n, s}
	}

	return []directedPoint{e, e, e} // this should not be a case, i am just doing it for the return
}

func solvePart2(input string) string {
	grid := utils.ParseGrid(utils.ParseLines(input))

	var reindeer directedPoint
	var end point

	maze := make([][]rune, len(grid))
	for y := range grid {
		maze[y] = make([]rune, len(grid[y]))
		for x := range grid[y] {
			switch grid[y][x] {
			case 'S':
				reindeer = directedPoint{point{x, y}, east}
				maze[y][x] = '.'
			case 'E':
				end = point{x, y}
				maze[y][x] = '.'
			default:
				maze[y][x] = grid[y][x]
			}
		}
	}

	bestPaths := findAllBestPaths(reindeer, end, maze)

	uniqueTiles := countUniqueTiles(bestPaths, maze)

	return strconv.Itoa(uniqueTiles)
}

func findAllBestPaths(reindeer directedPoint, end point, maze [][]rune) [][]directedPoint {
	bestScore := math.MaxInt
	queue := []routeState{{reindeer, []directedPoint{reindeer}, 0}}
	visited := make(map[directedPoint]int)
	bestPaths := [][]directedPoint{}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if len(state.steps) > 10000 || state.score > bestScore {
			continue
		}

		if state.reindeer.point == end {
			if state.score < bestScore {
				bestPaths = [][]directedPoint{state.steps}
				bestScore = state.score
			} else if state.score == bestScore {
				bestPaths = append(bestPaths, state.steps)
			}
			continue
		}

		for _, candidate := range getOffsets(state.reindeer) {
			if maze[candidate.point.y][candidate.point.x] == '.' {
				score := state.score + 1
				if state.reindeer.direction != candidate.direction {
					score += 1000
				}

				if previous, found := visited[candidate]; found {
					if previous < score {
						continue
					}
				}
				visited[candidate] = score

				newSteps := make([]directedPoint, len(state.steps))
				copy(newSteps, state.steps)

				queue = append(queue, routeState{
					reindeer: candidate,
					steps:    append(newSteps, candidate),
					score:    score})
			}
		}
	}

	return bestPaths
}

func countUniqueTiles(paths [][]directedPoint, maze [][]rune) int {
	uniqueTiles := make(map[point]bool)

	for _, path := range paths {
		for _, step := range path {
			if maze[step.point.y][step.point.x] != '#' {
				uniqueTiles[step.point] = true
			}
		}
	}

	return len(uniqueTiles)
}
