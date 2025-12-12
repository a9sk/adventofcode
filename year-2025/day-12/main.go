package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	sol1 := 0
	n := []int{6, 7, 5, 7, 7, 7}
	for i, p := range strings.Split(string(input), "\n\n") {
		if i < 6 {
			continue
		}
	m:
		for _, l := range strings.Split(p, "\n") {
			var x, y int
			ss := make([]int, 6)
			if n, _ := fmt.Sscanf(l, "%dx%d: %d %d %d %d %d %d", &x, &y, &ss[0], &ss[1], &ss[2], &ss[3], &ss[4], &ss[5]); n == 0 {
				continue
			}
			xy, s := x*y, 0
			for i, nn := range n {
				s += nn * ss[i]
				if s > xy {
					continue m
				}
			}
			sol1++
		}

	}

	fmt.Println("part 1:", sol1)
}
