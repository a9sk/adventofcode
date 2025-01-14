package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

const (
	WALL  = '#'
	ROBOT = '@'
	BOX   = 'O'
	EMPTY = '.'
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

func solvePart1(input string) string {
	parts := strings.Split(input, "\n\n")
	mapLines := strings.Split(parts[0], "\n")
	moves := strings.ReplaceAll(parts[1], "\n", "")

	warehouse := make([][]rune, len(mapLines))
	for i, line := range mapLines {
		warehouse[i] = []rune(line)
	}

	var robotPos point
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == ROBOT {
				robotPos = point{x, y}
				break
			}
		}
		if robotPos.x != 0 || robotPos.y != 0 {
			break
		}
	}

	for _, move := range moves {
		newWarehouse, moved := moveRobot(warehouse, robotPos, move)
		if moved {
			warehouse = newWarehouse
			robotPos = point{
				x: robotPos.x + map[rune]point{
					'^': {0, -1},
					'v': {0, 1},
					'<': {-1, 0},
					'>': {1, 0},
				}[move].x,
				y: robotPos.y + map[rune]point{
					'^': {0, -1},
					'v': {0, 1},
					'<': {-1, 0},
					'>': {1, 0},
				}[move].y,
			}
		}
	}

	var sum = 0
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == BOX {
				sum += (y * 100) + x
			}
		}
	}

	return strconv.Itoa(sum)
}

func moveRobot(warehouse [][]rune, pos point, direction rune) ([][]rune, bool) {
	var directions = map[rune]point{
		'^': {x: 0, y: -1},
		'v': {x: 0, y: 1},
		'<': {x: -1, y: 0},
		'>': {x: 1, y: 0},
	}
	dir := directions[direction]

	newWarehouse := make([][]rune, len(warehouse))
	for i := range warehouse {
		newWarehouse[i] = make([]rune, len(warehouse[i]))
		copy(newWarehouse[i], warehouse[i])
	}

	newX, newY := pos.x+dir.x, pos.y+dir.y

	if newX < 0 || newY < 0 || newY >= len(newWarehouse) ||
		newX >= len(newWarehouse[0]) || newWarehouse[newY][newX] == WALL {
		return warehouse, false
	}

	if newWarehouse[newY][newX] == BOX {
		boxesToPush := []point{{x: newX, y: newY}}

		currentX, currentY := newX, newY
		for {
			nextX, nextY := currentX+dir.x, currentY+dir.y

			if nextX <= 0 || nextY <= 0 || nextY >= len(newWarehouse)-1 ||
				nextX >= len(newWarehouse[0])-1 || newWarehouse[nextY][nextX] == WALL {
				return warehouse, false
			}

			if newWarehouse[nextY][nextX] == BOX {
				boxesToPush = append(boxesToPush, point{x: nextX, y: nextY})
				currentX, currentY = nextX, nextY
			} else {
				break
			}
		}

		for i := len(boxesToPush) - 1; i >= 0; i-- {
			box := boxesToPush[i]
			newBoxX, newBoxY := box.x+dir.x, box.y+dir.y
			newWarehouse[newBoxY][newBoxX] = BOX
			newWarehouse[box.y][box.x] = EMPTY
		}
	}

	newWarehouse[pos.y][pos.x] = EMPTY
	newWarehouse[newY][newX] = ROBOT

	return newWarehouse, true
}

func solvePart2(input string) string {
	parts := strings.Split(input, "\n\n")
	grid := parts[0]
	moves := strings.ReplaceAll(parts[1], "\n", "")

	r := strings.NewReplacer(
		"#", "##",
		"O", "[]",
		".", "..",
		"@", "@.")

	grid = r.Replace(grid)

	gridMap := make(map[point]rune)
	robotPos := point{}

	for y, line := range strings.Fields(grid) {
		for x, ch := range line {
			if ch == '@' {
				robotPos = point{x, y}
				ch = '.'
			}
			gridMap[point{x, y}] = ch
		}
	}

	delta := map[rune]point{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
		'[': {1, 0},
		']': {-1, 0},
	}

loop:
	for _, move := range moves {
		queue := []point{robotPos}
		boxes := make(map[point]rune)

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if _, ok := boxes[p]; ok {
				continue
			}

			boxes[p] = gridMap[p]

			n := point{p.x + delta[move].x, p.y + delta[move].y}

			switch gridMap[n] {
			case '#':
				continue loop
			case '[', ']':
				queue = append(queue, point{n.x + delta[gridMap[n]].x, n.y + delta[gridMap[n]].y})
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for p := range boxes {
			gridMap[p] = '.'
		}
		for p, ch := range boxes {
			gridMap[point{p.x + delta[move].x, p.y + delta[move].y}] = ch
		}
		robotPos = point{robotPos.x + delta[move].x, robotPos.y + delta[move].y}
	}

	sum := 0
	for p, ch := range gridMap {
		if ch == 'O' || ch == '[' {
			sum += 100*p.y + p.x
		}
	}

	return strconv.Itoa(sum)
}
