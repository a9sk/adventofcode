package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(in), "\n\n")
	var s, s2 int
	var m [][2]int
	c := strings.Split(lines[0], "\n")
	sort.Strings(c)

	for _, l := range c {
		t := strings.Split(l, "-")
		s, _ := strconv.Atoi(t[0])
		e, _ := strconv.Atoi(t[1])
	m:
		for i, n := range m {
			if s > n[1] || e < n[0] {
				continue
			}

			if s < n[0] {
				m[i][0] = s
			}
			if e > n[1] {
				m[i][1] = e
			}

			merged := m[i]
			m = append(m[:i], m[i+1:]...)

			s, e = merged[0], merged[1]
			goto m

		}
		m = append(m, [2]int{s, e})
	}

	for _, l := range strings.Split(lines[1], "\n") {
		v, _ := strconv.Atoi(l)
		for _, n := range m {
			if v >= n[0] && v <= n[1] {
				s++
				break
			}
		}
	}

	fmt.Println(m)

	for _, n := range m {
		s2 += (n[1] + 1) - (n[0])
	}

	fmt.Println("sol 1:", s)
	fmt.Println("sol 2:", s2)
}
