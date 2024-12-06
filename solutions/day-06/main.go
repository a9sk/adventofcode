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

func solvePart1(input string) string {
	field := utils.ParseGrid(utils.ParseLines(input))

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

	// var isOut bool
	// var direction = "up"
	// for {
	// 	if field, pos, direction, isOut = Move(field, pos, direction); isOut {
	// 		field[pos.y][pos.x] = 'X' // to make sure the last cell visited is actually considered as X
	// 		break
	// 	}
	// }

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

type position struct {
	y int
	x int
}

// func Move(field [][]rune, pos position, direction string) ([][]rune, position, string, bool) {
// 	switch direction {
// 	case "up":
// 		if pos.y-1 >= 0 && (field[pos.y-1][pos.x] == '.' || field[pos.y-1][pos.x] == 'X') { // len(field[pos.y]) > pos.x+1
// 			field[pos.y][pos.x] = 'X'
// 			newPos := position{pos.y - 1, pos.x}
// 			return field, newPos, "up", false
// 		} else if pos.y-1 >= 0 && field[pos.y-1][pos.x] == '#' {
// 			field[pos.y][pos.x] = 'X'
// 			return field, pos, "right", false
// 		} else if pos.y-1 < 0 { // you are at the top of the map and you are going out...
// 			return field, pos, "up", true
// 		}
// 	case "right":
// 		if pos.x+1 < len(field[pos.y]) && (field[pos.y][pos.x+1] == '.' || field[pos.y][pos.x+1] == 'X') {
// 			field[pos.y][pos.x] = 'X'
// 			newPos := position{pos.y, pos.x + 1}
// 			return field, newPos, "right", false
// 		} else if pos.x+1 < len(field[pos.y]) && field[pos.y][pos.x+1] == '#' {
// 			field[pos.y][pos.x] = 'X'
// 			return field, pos, "down", false
// 		} else if pos.x+1 == len(field[pos.y]) { // if you get to len() then you are out
// 			return field, pos, "right", true
// 		}
// 	case "down":
// 		if pos.y+1 < len(field) && (field[pos.y+1][pos.x] == '.' || field[pos.y+1][pos.x] == 'X') {
// 			field[pos.y][pos.x] = 'X'
// 			newPos := position{pos.y + 1, pos.x}
// 			return field, newPos, "down", false
// 		} else if pos.y+1 < len(field) && field[pos.y+1][pos.x] == '#' {
// 			field[pos.y][pos.x] = 'X'
// 			return field, pos, "left", false
// 		} else if pos.y+1 == len(field) {
// 			return field, pos, "down", true
// 		}
// 	case "left":
// 		if pos.x-1 >= 0 && (field[pos.y][pos.x-1] == '.' || field[pos.y][pos.x-1] == 'X') {
// 			field[pos.y][pos.x] = 'X'
// 			newPos := position{pos.y, pos.x - 1}
// 			return field, newPos, "left", false
// 		} else if pos.x-1 >= 0 && field[pos.y][pos.x-1] == '#' {
// 			field[pos.y][pos.x] = 'X'
// 			return field, pos, "up", false
// 		} else if pos.x == 0 {
// 			return field, pos, "left", true
// 		}
// 	}

// 	return field, pos, direction, false // should not be a possibility but if it happens i want the for to stop
// }

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

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
