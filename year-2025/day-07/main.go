package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

var Count int
var v = make(map[Point]bool)

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Println("part 2:", r(lines, Point{len(lines[0]) / 2, 0}, map[Point]int{}))
	fmt.Println("part 1:", Count)
}

func r(l []string, p Point, cache map[Point]int) int {
	if v, ok := cache[p]; ok {
		return v
	}

	if len(l)-1 == p.y+1 {
		cache[p] = 1
		return 1
	}

	n := Point{p.x, p.y + 1}

	if l[p.y+1][p.x] == '.' {
		cache[p] = r(l, n, cache)
		return cache[p]
	}

	t := 0
	if !v[n] {
		Count++
		v[n] = true
	}

	if p.x+1 < len(l[0]) {
		t += r(l, Point{p.x + 1, p.y}, cache)
	}
	if p.x-1 >= 0 {
		t += r(l, Point{p.x - 1, p.y}, cache)
	}

	return t
}
