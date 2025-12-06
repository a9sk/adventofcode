#!/bin/bash

DAY=$(date +%d)
if [ -n "$1" ]; then
    DAY=$(printf "%02d" "$1")
fi

DAY_DIR="year-2025/day-${DAY}"

if [ -d "$DAY_DIR" ]; then
    echo "directory '$DAY_DIR' already exists."
    exit 0
else
    mkdir "$DAY_DIR"
    echo "created directory '$DAY_DIR'."
fi

cat <<EOF > "$DAY_DIR/main.go"
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("part 1:", solvePart1(string(input)))

	fmt.Println("part 2:", solvePart2(string(input)))
}

func solvePart1(input string) string {
	return "no solution yet"
}

func solvePart2(input string) string {
	return "no solution yet"
}
EOF

touch "${DAY_DIR}/input.txt"
echo "initialized input.txt in '$DAY_DIR'."

touch "${DAY_DIR}/test_input.txt"
echo "initialized test_input.txt in '$DAY_DIR'."

echo "setup complete for day ${DAY}!"
