package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	ps := []Point{}
	for _, l := range strings.Split(string(input), "\n") {
		var x, y int
		fmt.Sscanf(l, "%d,%d", &x, &y)
		ps = append(ps, Point{x, y})
	}

	ps = ps[:len(ps)-1]

	max := 0
	for i, p1 := range ps {
		for j := i + 1; j < len(ps); j++ {
			if a := area(p1, ps[j]); a > max {
				max = a
			}
		}
	}

	fmt.Println("part 1:", max)
}

func area(a, b Point) int {
	return (max(a.x, b.x) - min(a.x, b.x) + 1) * (max(a.y, b.y) - min(a.y, b.y) + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
