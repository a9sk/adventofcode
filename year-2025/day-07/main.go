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
	r(lines, Point{len(lines[0]) / 2, 0})

	fmt.Println("part 1:", Count)
}

func r(l []string, p Point) {
	if len(l)-1 == p.y+1 {
		return
	}

	if l[p.y+1][p.x] == '.' {
		r(l, Point{p.x, p.y + 1})
		return
	}

	if l[p.y+1][p.x] == '^' {
		if !v[Point{p.x, p.y + 1}] {
			Count++
			v[Point{p.x, p.y + 1}] = true
		} else {
			return // the most important else { return } i have ever written
		}

		if p.x+1 < len(l[0]) {
			r(l, Point{p.x + 1, p.y})
		}
		if p.x-1 >= 0 {
			r(l, Point{p.x - 1, p.y})
		}

		return
	}
}
