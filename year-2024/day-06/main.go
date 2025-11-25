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

type position struct {
	y int
	x int
}

var directions = []struct {
	name   string
	dy, dx int
	next   string
}{
	{"up", -1, 0, "right"},
	{"right", 0, 1, "down"},
	{"down", 1, 0, "left"},
	{"left", 0, -1, "up"},
}

func solvePart1(input string) string {

	lines := strings.Split(input, "\n")

	field := make([][]rune, len(lines))
	for i, line := range lines {
		field[i] = []rune(line)
	}

	var pos position

	// find the starting position
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == '^' {
				pos = position{i, j}
				break
			}
		}
	}

	field, _ = recursiveMove(field, pos, "up")

	var sum = 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == 'X' {
				sum++
			}
		}
	}
	return strconv.Itoa(sum)
}

func recursiveMove(field [][]rune, pos position, direction string) ([][]rune, bool) {
	field[pos.y][pos.x] = 'X'

	var dir struct {
		name   string
		dy, dx int
		next   string
	}
	for _, d := range directions {
		if d.name == direction {
			dir = d
			break
		}
	}

	nextPos := position{pos.y + dir.dy, pos.x + dir.dx}

	if nextPos.y < 0 || nextPos.y >= len(field) || nextPos.x < 0 || nextPos.x >= len(field[0]) {
		return field, true
	}

	nextCell := field[nextPos.y][nextPos.x]
	if nextCell == '.' || nextCell == 'X' {
		return recursiveMove(field, nextPos, dir.name)
	} else if nextCell == '#' {
		return recursiveMove(field, pos, dir.next)
	}

	return field, false
}

func solvePart2(input string) string {

	lines := strings.Split(input, "\n")

	restartField := make([][]rune, len(lines))
	for i, line := range lines {
		restartField[i] = []rune(line)
	}

	var startPos position
	for i := 0; i < len(restartField); i++ {
		for j := 0; j < len(restartField[i]); j++ {
			if restartField[i][j] == '^' {
				startPos = position{i, j}
				break
			}
		}
	}

	baseReachable := make(map[position]bool)
	traceBasePathIterative(restartField, startPos, baseReachable)

	var sum = 0
	for pos := range baseReachable {
		if restartField[pos.y][pos.x] == '#' || restartField[pos.y][pos.x] == '^' {
			continue
		}

		if checkLoopFast(restartField, startPos, pos) {
			sum++
		}
	}
	return strconv.Itoa(sum)
}

func traceBasePathIterative(field [][]rune, start position, reachable map[position]bool) {
	type state struct {
		pos       position
		direction string
	}

	queue := []state{{pos: start, direction: "up"}}
	visited := make(map[state]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true
		reachable[current.pos] = true

		var dir struct {
			name   string
			dy, dx int
			next   string
		}
		for _, d := range directions {
			if d.name == current.direction {
				dir = d
				break
			}
		}

		nextPos := position{current.pos.y + dir.dy, current.pos.x + dir.dx}
		if nextPos.y < 0 || nextPos.y >= len(field) || nextPos.x < 0 || nextPos.x >= len(field[0]) {
			continue
		}

		nextCell := field[nextPos.y][nextPos.x]
		if nextCell == '.' || nextCell == '^' {
			queue = append(queue, state{pos: nextPos, direction: current.direction})
		} else if nextCell == '#' {
			queue = append(queue, state{pos: current.pos, direction: dir.next})
		}
	}
}

func checkLoopFast(field [][]rune, start position, blocked position) bool {
	type state struct {
		pos       position
		direction string
	}

	visited := make(map[state]bool)
	current := state{pos: start, direction: "up"}

	for {
		if visited[current] {
			return true
		}
		visited[current] = true

		var dir struct {
			name   string
			dy, dx int
			next   string
		}
		for _, d := range directions {
			if d.name == current.direction {
				dir = d
				break
			}
		}

		nextPos := position{current.pos.y + dir.dy, current.pos.x + dir.dx}
		if nextPos.y < 0 || nextPos.y >= len(field) || nextPos.x < 0 || nextPos.x >= len(field[0]) {
			return false
		}

		if nextPos == blocked {
			current.direction = dir.next
			continue
		}

		nextCell := field[nextPos.y][nextPos.x]
		if nextCell == '.' || nextCell == '^' {
			current.pos = nextPos
		} else if nextCell == '#' {
			current.direction = dir.next
		}
	}
}
