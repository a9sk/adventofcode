package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

type Area struct {
	p1, p2 Point
	a      int
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
	as := make([]Area, 0)
	for i, p1 := range ps {
		for j := i + 1; j < len(ps); j++ {
			a := area(p1, ps[j])
			if a > max {
				max = a
			}
			as = append(as, Area{p1, ps[j], a})
		}
	}
	fmt.Println("part 1:", max)

	sort.Slice(as, func(a, b int) bool { return as[a].a > as[b].a })
	edg := make([]Point, 0)
	for i, p := range ps {
		var pp Point
		if i == len(ps)-1 {
			pp = ps[0]
		} else {
			pp = ps[i+1]
		}
		edg = append(edg, cedg(p, pp)...)
	}

	edg = append(edg, ps...)

m:
	for j := 0; j < len(as); j++ {
		for _, e := range edg {
			if i(as[j], e) {
				continue m
			}
		}
		fmt.Println("part 2:", as[j])
		break
	}
}

func i(a Area, e Point) bool {
	return e.x < max(a.p1.x, a.p2.x) && e.x > min(a.p1.x, a.p2.x) && e.y < max(a.p1.y, a.p2.y) && e.y > min(a.p1.y, a.p2.y)
}

// cedg (fast for "calculate edge") finds connection between points, supposes that
// if x is different then y is equals (since that is what the problem says)
// also the extremes are not added and need to be appended later
func cedg(a, b Point) (c []Point) {
	if a.x == b.x {
		for i := min(a.y, b.y) + 1; i < max(a.y, b.y); i++ {
			c = append(c, Point{a.x, i})
		}
		return c
	}
	for i := min(a.x, b.x) + 1; i < max(a.x, b.x); i++ {
		c = append(c, Point{i, a.y})
	}
	return c
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
