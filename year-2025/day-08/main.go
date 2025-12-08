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
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	type edge struct{ i, j, d int }
	var ds []edge
	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := p[i].x - p[j].x
			dy := p[i].y - p[j].y
			dz := p[i].z - p[j].z
			d := dx*dx + dy*dy + dz*dz
			ds = append(ds, edge{i, j, d})
		}
	}

	sort.Slice(ds, func(a, b int) bool { return ds[a].d < ds[b].d })

	for k := 0; k < 1000 && k < len(ds); k++ {
		union(ds[k].i, ds[k].j)
	}

	var sizes []int
	for i := range n {
		if find(i) == i {
			sizes = append(sizes, size[i])
		}
	}

	sort.Slice(sizes, func(a, b int) bool { return sizes[a] > sizes[b] })
	fmt.Println("part 1: ", sizes[0]*sizes[1]*sizes[2])
}
