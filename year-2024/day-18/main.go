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

// i just now noticed that structs in go should have the fist letter in caps
type Point struct {
	x, y int
}

func solvePart1(input string) string {

	bytes := []Point{}

	for _, line := range strings.Split(input, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		bytes = append(bytes, Point{x: x, y: y})
	}

	field := map[Point]bool{}
	for y := 0; y < 71; y++ {
		for x := 0; x < 71; x++ {
			field[Point{x: x, y: y}] = true
		}
	}

mainFor:
	for bite := range bytes { // i use bite instead of byte cos i do not like byte's colour being different
		field[bytes[bite]] = false

		queue, distance := []Point{{0, 0}}, map[Point]int{{0, 0}: 0}
		for len(queue) > 0 {
			position := queue[0]
			queue = queue[1:]

			if position == (Point{70, 70}) {
				if bite == 1024 {
					return strconv.Itoa(distance[position])
				}
				continue mainFor
			}

			for _, direction := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				new := Point{x: position.x + direction.x, y: position.y + direction.y}
				if _, ok := distance[new]; !ok && field[new] {
					queue, distance[new] = append(queue, new), distance[position]+1
				}
			}
		}
		break
	}

	return ""
}

func solvePart2(input string) string {

	/*  NOTE:
	    Part 2 is pretty much the same to part one, the only difference is
		that you do not have to leave the loop when you find the shortest path
		and wait untill the field gets blocked out (when the length of the
		queue is zero)
	*/

	bytes := []Point{}

	for _, line := range strings.Split(input, "\n") {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		bytes = append(bytes, Point{x: x, y: y})
	}

	field := map[Point]bool{}
	for y := 0; y < 71; y++ {
		for x := 0; x < 71; x++ {
			field[Point{x: x, y: y}] = true
		}
	}

mainFor:
	for bite := range bytes {
		field[bytes[bite]] = false

		queue, distance := []Point{{0, 0}}, map[Point]int{{0, 0}: 0}
		for len(queue) > 0 {
			position := queue[0]
			queue = queue[1:]

			if position == (Point{70, 70}) {
				continue mainFor
			}

			for _, direction := range []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				new := Point{x: position.x + direction.x, y: position.y + direction.y}
				if _, ok := distance[new]; !ok && field[new] {
					queue, distance[new] = append(queue, new), distance[position]+1
				}
			}
		}
		return fmt.Sprintf("%d,%d", bytes[bite].x, bytes[bite].y)
	}

	return "the input is probably wrong"
}
