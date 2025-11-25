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

const fieldWidth = 101
const fieldHeight = 103

const xmid = ((fieldWidth - 1) / 2) + 1
const ymid = ((fieldHeight - 1) / 2) + 1

type point struct {
	x, y int
}

type directions struct {
	x, y int
}

type robot struct {
	position point
	vector   directions
}

func solvePart1(input string) string {
	robots := parseRobots(input)

	quadrants := make([]int, 5)
	for _, robot := range robots {
		i := findEndingSector(robot)
		quadrants[i]++
	}

	sum := quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]

	return strconv.Itoa(sum)
}

func findEndingSector(robot robot) int {

	x := (((robot.position.x+robot.vector.x*100)%fieldWidth + fieldWidth) % fieldWidth) + 1
	y := (((robot.position.y+robot.vector.y*100)%fieldHeight + fieldHeight) % fieldHeight) + 1

	return checkQuadrants(x, y)
}

func checkQuadrants(x, y int) int {
	if x == xmid || y == ymid {
		return 0
	}

	if x < xmid && y < ymid {
		return 1
	}

	if x > xmid && y < ymid {
		return 2
	}

	if x < xmid && y > ymid {
		return 3
	}

	if x > xmid && y > ymid {
		return 4
	}

	return 0
}

func parseRobots(input string) []robot {
	lines := strings.Split(input, "\n")
	robots := []robot{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		var r robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.position.x, &r.position.y, &r.vector.x, &r.vector.y)
		robots = append(robots, r)
	}
	return robots
}

func solvePart2(input string) string {

	robots := parseRobots(input)

	for t := 1; ; t++ {
		seen := make(map[point]struct{})

		for i := range robots {
			robots[i].position = point{x: ((robots[i].position.x+robots[i].vector.x)%fieldWidth + fieldWidth) % fieldWidth, y: ((robots[i].position.y+robots[i].vector.y)%fieldHeight + fieldHeight) % fieldHeight}
			seen[robots[i].position] = struct{}{}
		}

		if len(seen) == len(robots) {
			return strconv.Itoa(t)
		}
	}
}
