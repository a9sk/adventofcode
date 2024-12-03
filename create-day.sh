#!/bin/bash

DAY=$(date +%d)
if [ -n "$1" ]; then
    DAY=$(printf "%02d" "$1")
fi

DAY_DIR="solutions/day-${DAY}"

if [ -d "$DAY_DIR" ]; then
    echo "Directory '$DAY_DIR' already exists."
    exit 0
else
    mkdir "$DAY_DIR"
    echo "Created directory '$DAY_DIR'."
fi

cat <<EOF > "$DAY_DIR/main.go"
package main

import (
	"bufio"
	"fmt"
	"os"
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

	return "Solution for Part 1 not implemented"
}

func solvePart2(input string) string {

	return "Solution for Part 2 not implemented"
}
EOF

touch "${DAY_DIR}/input.txt"
echo "Initialized input.txt in '$DAY_DIR'."

touch "${DAY_DIR}/test_input.txt"
echo "Initialized test_input.txt in '$DAY_DIR'."

echo "Setup complete for Day ${DAY}!"
