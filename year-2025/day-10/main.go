package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	sol1 := 0

	for _, l := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		p := strings.Split(l, " ")

		a := p[0][1 : len(p[0])-1]

		gn := 0
		for i, c := range a {
			if c == '#' {
				gn += 1 << i
			}
		}

		btns := p[1 : len(p)-1]

		var bb []int
		var ns []int
		for _, b := range btns {
			var n []int
			var s int
			for _, x := range strings.Split(b[1:len(b)-1], ",") {
				v, _ := strconv.Atoi(x)
				n = append(n, v)
				s += 1 << v
			}
			bb = append(bb, s)
			ns = append(ns, n...)
		}

		sc := len(btns)
		for i := range 1 << len(btns) {
			an, as := 0, 0
			for j := range len(btns) {
				if (i>>j)%2 == 1 {
					an ^= bb[j]
					as++
				}
			}
			if an == gn {
				sc = min(sc, as)
			}
		}
		sol1 += sc
	}

	fmt.Println("part 1:", sol1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
