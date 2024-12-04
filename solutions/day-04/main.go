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
	parsedInput := utils.ParseLines(input)
	matrix := utils.ParseGrid(parsedInput)

	total := 0

	for rotation := 0; rotation < 4; rotation++ {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i])-3; j++ {
				if matrix[i][j] == 'X' && matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
					total++
					// fmt.Printf("found: %c %c %c %c\n", matrix[i][j], matrix[i][j+1], matrix[i][j+2], matrix[i][j+3])
				}

				if i < len(matrix[i])-3 && matrix[i][j] == 'X' && matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
					total++
					// fmt.Printf("found: %c %c %c %c\n", matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3])
				}
			}
		}
		// fmt.Println("ROTATE")
		matrix = utils.Rotate90(matrix)
	}

	return strconv.Itoa(total)
}

func solvePart2(input string) string {
	parsedInput := utils.ParseLines(input)
	matrix := utils.ParseGrid(parsedInput)

	total := 0

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			// very long if statement coming
			if matrix[i][j] == 'A' && ((matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M') || (matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S')) && ((matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] == 'M') || (matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] == 'S')) {
				total++
				// fmt.Println("found an X-MAS")
			}
		}
	}

	return strconv.Itoa(total)
}
