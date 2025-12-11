package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type t struct {
	s        string
	fft, dac bool
}

func main() {
	input, _ := os.ReadFile("input.txt")

	s := map[string][]string{}
	for _, l := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		p := strings.Split(l, " ")
		s[p[0][0:len(p[0])-1]] = p[1:]
	}

	fmt.Println("part 1:", c(s, s["you"]))
	fmt.Println("part 2:", c2(s, s["svr"], false, false, map[t]int{}))
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

func c2(m map[string][]string, y []string, fft, dac bool, cache map[t]int) (r int) {
	if slices.Contains(y, "out") {
		if fft && dac {
			return 1
		}
		return 0
	}

	for _, i := range y {
		nfft, ndac := fft, dac
		if v, ok := cache[t{i, nfft, ndac}]; ok {
			r += v
			continue
		}

		switch i {
		case "fft":
			nfft = true
		case "dac":
			ndac = true
		default:
		}

		cache[t{i, nfft, ndac}] = c2(m, m[i], nfft, ndac, cache)
		r += cache[t{i, nfft, ndac}]
	}

	return r
}
