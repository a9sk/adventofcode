package main

import (
	"bufio"
	"fmt"
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

type Gate struct {
	gateType string
	in1, in2 string
	out      string
}

func solvePart1(input string) string {

	parts := strings.Split(input, "\n\n")

	wireValues := make(map[string]int)
	gates := make([]Gate, 0)

	for _, line := range strings.Split(parts[0], "\n") {
		lineParts := strings.Split(line, ": ")
		if val := lineParts[1][0] - '0'; val <= 1 {
			wireValues[lineParts[0]] = int(val)
		}
	}

	for _, line := range strings.Split(parts[1], "\n") {
		lineParts := strings.Split(line, " -> ")
		inputs := strings.Split(lineParts[0], " ")
		if len(inputs) == 3 {
			gates = append(gates, Gate{
				gateType: inputs[1],
				in1:      inputs[0],
				in2:      inputs[2],
				out:      lineParts[1],
			})
		}
	}

	changed := true
	for changed {
		changed = false
		for _, gate := range gates {
			if _, exists := wireValues[gate.out]; !exists {
				if evaluate(gate, wireValues) {
					changed = true
				}
			}
		}
	}

	result := 0
	zCount := 0
	for wire, value := range wireValues {
		if strings.HasPrefix(wire, "z") {
			pos := 0
			fmt.Sscanf(wire[1:], "%d", &pos) // i belive this to be pretty neat
			if value == 1 {
				result |= 1 << pos
			}
			if pos > zCount {
				zCount = pos
			}
		}
	}

	return strconv.Itoa(result)
}

func evaluate(gate Gate, wireValues map[string]int) bool {
	v1, ok1 := wireValues[gate.in1]
	v2, ok2 := wireValues[gate.in2]
	if !ok1 || !ok2 {
		return false
	}

	var result int
	switch gate.gateType {
	case "AND":
		result = v1 & v2
	case "OR":
		result = v1 | v2
	case "XOR":
		result = v1 ^ v2
	}
	wireValues[gate.out] = result
	return true
}

// i could not solve today's part 2, so i had to look for the correct approach on the reddit
func solvePart2(input string) string {
	parts := strings.Split(input, "\n\n")
	gates := strings.Split(parts[1], "\n")
	return swapAndJoinWires(gates)
}

func find(a, b, operator string, gates []string) string {
	for _, gate := range gates {
		if strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", a, operator, b)) ||
			strings.HasPrefix(gate, fmt.Sprintf("%s %s %s", b, operator, a)) {
			parts := strings.Split(gate, " -> ")
			return parts[len(parts)-1]
		}
	}
	return ""
}

func swapAndJoinWires(gates []string) string {
	var swapped []string
	var c0 string

	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		var m1, n1, r1, z1, c1 string

		m1 = find("x"+n, "y"+n, "XOR", gates)
		n1 = find("x"+n, "y"+n, "AND", gates)

		if c0 != "" {
			r1 = find(c0, m1, "AND", gates)
			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				r1 = find(c0, m1, "AND", gates)
			}

			z1 = find(c0, m1, "XOR", gates)
			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			c1 = find(r1, n1, "OR", gates)
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if c0 == "" {
			c0 = n1
		} else {
			c0 = c1
		}
	}

	sort.Strings(swapped)
	return strings.Join(swapped, ",")
}
