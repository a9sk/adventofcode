package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(in), "\n\n")
	var s int
	var m [][2]int

	for _, l := range strings.Split(lines[0], "\n") {
		t := strings.Split(l, "-")
		s, _ := strconv.Atoi(t[0])
		e, _ := strconv.Atoi(t[1])
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

	fmt.Println("sol 1:", s)
}
