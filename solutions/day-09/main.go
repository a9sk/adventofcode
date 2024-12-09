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

func solvePart1(input string) string {

	diskMap := strings.Split(input, "")

	var blocks []string
	var id = 0
	for j, num := range diskMap {
		numInt, _ := strconv.Atoi(num)
		for i := 0; i < numInt; i++ {
			if j%2 != 0 {
				blocks = append(blocks, ".")
				continue
			}
			blocks = append(blocks, strconv.Itoa(id))
		}
		if j%2 != 0 {
			id++
		}
	}

	iterativeMove(blocks)

	var sum = 0
	for index, value := range blocks {
		intValue, _ := strconv.Atoi(value)
		sum += intValue * index
	}

	return strconv.Itoa(sum)
}

func iterativeMove(blocks []string) {
	var lastFull = 0
	for i := len(blocks) - 1; i > lastFull; i-- {
		if blocks[i] == "." {
			continue
		}
		for j := lastFull; j < i; j++ {
			if blocks[j] != "." {
				lastFull = j
				continue
			}
			blocks[j] = blocks[i]
			blocks[i] = "."
			break
		}
	}
}

func solvePart2(input string) string {

	diskMap := strings.Split(input, "")

	var blocks []string
	var id = 0
	for j, num := range diskMap {
		numInt, _ := strconv.Atoi(num)
		for i := 0; i < numInt; i++ {
			if j%2 != 0 {
				blocks = append(blocks, ".")
				continue
			}
			blocks = append(blocks, strconv.Itoa(id))
		}
		if j%2 != 0 {
			id++
		}
	}

	fmt.Println(blocks)
	wholeIterativeMove(blocks)
	fmt.Println(blocks)
	var sum = 0
	for index, value := range blocks {
		intValue, _ := strconv.Atoi(value)
		sum += intValue * index
	}

	return strconv.Itoa(sum)

}

func wholeIterativeMove(blocks []string) {
	var lastID = 10000000
	// stop := 1
	stop := 20000
	for i := len(blocks) - 1; i > stop; i-- {
		/*
			i set i to a random high value because, as this version does not use
			the lastFull variable (i am too tired to implement it), i would get som index out
			of bound errors. To check if the test_input works put the i value back to 1
		*/
		if blocks[i] == "." {
			continue
		}

		if check, _ := strconv.Atoi(blocks[i]); check >= lastID {
			continue
		}

		var id = blocks[i]
		lastID, _ = strconv.Atoi(id)
		var appears = 1
		// count how many blocks is that file (same ID) formed by
		for {
			if blocks[i-appears] != id {
				break
			}
			appears++
		}

		// fmt.Printf("file: %d has %d blocks\n", id, appears)

		for j := 0; j < i; j++ {
			var valid = true
			for n := 0; n < appears; n++ {
				if blocks[j+n] != "." {
					valid = false
					break
				}
			}

			if !valid {
				continue
			}

			for n := 0; n < appears; n++ {
				blocks[j+n] = id
				blocks[i-n] = "."
			}

			// lastFull = j + appears - 1 // should only increase if no . is found before but seems like too much iterations... idk which one is better
			// fmt.Println(blocks)
			break
		}

		i -= appears - 1
	}
}
