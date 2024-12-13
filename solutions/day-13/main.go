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

type button struct {
	x, y int
}

type prize struct {
	x, y int
}

func solvePart1(input string) string {
	clawMachines := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	var sum = 0
	for _, clawMachine := range clawMachines {
		a, b, p := parseButtons(strings.Split(clawMachine, "\n"))
		sum += findPresses(a, b, p)
	}
	return strconv.Itoa(sum)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}

func parseButtons(clawMachine []string) (button, button, prize) {
	var a, b button
	var p prize
	fmt.Sscanf(clawMachine[0], "Button A: X+%d, Y+%d", &a.x, &a.y)
	fmt.Sscanf(clawMachine[1], "Button B: X+%d, Y+%d", &b.x, &b.y)
	fmt.Sscanf(clawMachine[2], "Prize: X=%d, Y=%d", &p.x, &p.y)
	return a, b, p
}

func findPresses(a button, b button, p prize) int {
	presses := 0
	D, Dx, Dy := a.x*b.y-b.x*a.y, p.x*b.y-b.x*p.y, a.x*p.y-p.x*a.y

	if D != 0 && Dx == (Dx/D)*D && Dy == (Dy/D)*D {
		presses += (Dx/D)*3 + (Dy / D)
	}

	return presses
}
