package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type Point struct {
	x, y int
}

var directions = map[string]Point{
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
}

var keypad = map[string]Point{
	"7": {0, 0},
	"8": {1, 0},
	"9": {2, 0},
	"4": {0, 1},
	"5": {1, 1},
	"6": {2, 1},
	"1": {0, 2},
	"2": {1, 2},
	"3": {2, 2},
	"#": {0, 3},
	"0": {1, 3},
	"A": {2, 3},
}

var secondKeypad = map[string]Point{
	"#": {0, 0},
	"^": {1, 0},
	"A": {2, 0},
	"<": {0, 1},
	"v": {1, 1},
	">": {2, 1},
}

func solvePart1(input string) string {

	keycodes := strings.Split(input, "\n")
	cache := make(map[string]int)
	var sum = 0

	var findPress func(input map[string]Point, code string, robot int) int
	findPress = func(input map[string]Point, code string, robot int) int {
		key := fmt.Sprintf("%s,%d", code, robot)
		if n, ok := cache[key]; ok {
			return n
		}

		current := "A"
		length := 0
		for i := 0; i < len(code); i++ {
			moves := findPaths(input, current, string(code[i]))
			if robot == 0 {
				length += len(moves[0])
			} else {
				minLength := math.MaxInt
				for _, move := range moves {
					minLength = int(math.Min(float64(minLength), float64(findPress(secondKeypad, move, robot-1))))
				}
				length += minLength
			}
			current = string(code[i])
		}

		cache[key] = length
		return length
	}

	for _, code := range keycodes {
		n := 0
		for _, c := range code {
			if c >= '0' && c <= '9' {
				n = n*10 + int(c-'0')
			}
		}
		sum += n * findPress(keypad, code, 2)
	}

	return strconv.Itoa(sum)
}

func findPaths(input map[string]Point, start, end string) []string {
	queue := []struct {
		x, y int
		path string
	}{{input[start].x, input[start].y, ""}}
	distances := make(map[string]int)
	allPaths := []string{}

	if start == end {
		return []string{"A"}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.x == input[end].x && current.y == input[end].y {
			allPaths = append(allPaths, current.path+"A")
		}
		if distances[fmt.Sprintf("%d,%d", current.x, current.y)] != 0 && distances[fmt.Sprintf("%d,%d", current.x, current.y)] < len(current.path) {
			continue
		}

		for direction, vec := range directions {
			pos := struct{ x, y int }{current.x + vec.x, current.y + vec.y}
			if input["#"].x == pos.x && input["#"].y == pos.y {
				continue
			}

			for _, button := range input {
				if button.x == pos.x && button.y == pos.y {
					newPath := current.path + direction
					if distances[fmt.Sprintf("%d,%d", pos.x, pos.y)] == 0 || distances[fmt.Sprintf("%d,%d", pos.x, pos.y)] >= len(newPath) {
						queue = append(queue, struct {
							x, y int
							path string
						}{pos.x, pos.y, newPath})
						distances[fmt.Sprintf("%d,%d", pos.x, pos.y)] = len(newPath)
					}
				}
			}
		}
	}

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return allPaths
}

func solvePart2(input string) string {

	keycodes := strings.Split(input, "\n")
	cache := make(map[string]int)
	var sum = 0

	var findPress func(input map[string]Point, code string, robot int) int
	findPress = func(input map[string]Point, code string, robot int) int {
		key := fmt.Sprintf("%s,%d", code, robot)
		if n, ok := cache[key]; ok {
			return n
		}

		current := "A"
		length := 0
		for i := 0; i < len(code); i++ {
			moves := findPaths(input, current, string(code[i]))
			if robot == 0 {
				length += len(moves[0])
			} else {
				minLength := math.MaxInt
				for _, move := range moves {
					minLength = int(math.Min(float64(minLength), float64(findPress(secondKeypad, move, robot-1))))
				}
				length += minLength
			}
			current = string(code[i])
		}

		cache[key] = length
		return length
	}

	for _, code := range keycodes {
		n := 0
		for _, c := range code {
			if c >= '0' && c <= '9' {
				n = n*10 + int(c-'0')
			}
		}
		// the only difference between part 1 and part 2 is here...
		sum += n * findPress(keypad, code, 25)
	}

	return strconv.Itoa(sum)
}
