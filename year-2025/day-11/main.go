package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	s := map[string][]string{}
	for _, l := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		p := strings.Split(l, " ")
		s[p[0][0:len(p[0])-1]] = p[1:]
	}

	fmt.Println("part 1:", c(s, s["you"]))
}

func c(m map[string][]string, y []string) (r int) {
	if slices.Contains(y, "out") {
		return 1
	}

	for _, i := range y {
		r += c(m, m[i])
	}

	return r
}
