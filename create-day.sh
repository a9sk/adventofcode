#!/bin/bash

DAY=$(date +%d)
if [ -n "$1" ]; then
    DAY=$(printf "%02d" "$1")
fi

DAY_DIR="solutions/day-${DAY}"

if [ -d "$DAY_DIR" ]; then
    echo "Directory '$DAY_DIR' already exists."
else
    mkdir "$DAY_DIR"
    echo "Created directory '$DAY_DIR'."
fi

MAIN_GO="${DAY_DIR}/main.go"
if [ -f "$MAIN_GO" ]; then
    echo "'main.go' already exists in '$DAY_DIR'. Skipping creation."
else
  cat << EOF > "$MAIN_GO"
package main

import (
	"bufio"
	"fmt"
	"os"
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

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := solve(input)
	fmt.Println("Solution:", result)
}

func solve(input []string) string {
	return "Implement me!"
}
EOF
    
    echo "Created 'main.go' in '$DAY_DIR'."
fi

# Create empty input files
touch "${DAY_DIR}/input.txt"
echo "Initialized input.txt in '$DAY_DIR'."

echo "Setup complete for Day ${DAY}!"
