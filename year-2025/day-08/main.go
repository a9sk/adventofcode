package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y, z int
}

func main() {
	in, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(in)), "\n")

	var p []Point
	for _, l := range lines {
		var x, y, z int
		fmt.Sscanf(l, "%d,%d,%d", &x, &y, &z)
		p = append(p, Point{x, y, z})
	}

	n := len(p)

	var es []struct{ i, j, d int }
	for i := range n {
		for j := i + 1; j < n; j++ {
			es = append(es, struct{ i, j, d int }{i, j, d(p[i], p[j])})
		}
	}
	sort.Slice(es, func(a, b int) bool { return es[a].d < es[b].d })

	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	for k := 0; k < 1000 && k < len(es); k++ {
		u(parent, size, es[k].i, es[k].j)
	}

	var sizes []int
	for i := range n {
		if f(parent, i) == i {
			sizes = append(sizes, size[i])
		}
	}
	sort.Slice(sizes, func(a, b int) bool { return sizes[a] > sizes[b] })

	fmt.Println("part 1:", sizes[0]*sizes[1]*sizes[2])

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	components := n
	var last struct{ i, j, d int }

	for _, e := range es {
		if u(parent, size, e.i, e.j) {
			last = e
			components--
			if components == 1 {
				break
			}
		}
	}

	fmt.Println("part 2:", p[last.i].x*p[last.j].x)
}

func d(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func f(parent []int, x int) int {
	for parent[x] != x {
		parent[x] = parent[parent[x]]
		x = parent[x]
	}
	return x
}

func u(parent, size []int, a, b int) bool {
	ra, rb := f(parent, a), f(parent, b)
	if ra == rb {
		return false
	}
	if size[ra] < size[rb] {
		ra, rb = rb, ra
	}
	parent[rb] = ra
	size[ra] += size[rb]
	return true
}
