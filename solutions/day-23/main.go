package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	connections := strings.Split(input, "\n")
	var containsT []string
	graph := make(map[string][]string)
	for _, connection := range connections {
		p := strings.Split(connection, "-")
		if p[0][0] == 't' {
			containsT = append(containsT, p[0])
		}
		if p[1][0] == 't' {
			containsT = append(containsT, p[1])
		}
		graph[p[0]] = append(graph[p[0]], p[1])
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	var sum = 0
	var cle []string
	for _, Tnode := range containsT {
		for _, cTnode := range graph[Tnode] {
			for _, ccTnode := range graph[cTnode] {
				if slices.Contains(graph[ccTnode], Tnode) {
					cycle := []string{Tnode, cTnode, ccTnode}
					slices.Sort(cycle)
					if !slices.Contains(cle, fmt.Sprintf("%s%s%s", cycle[0], cycle[1], cycle[2])) {
						cle = append(cle, fmt.Sprintf("%s%s%s", cycle[0], cycle[1], cycle[2]))
						sum++
					}
				}
			}
		}
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
