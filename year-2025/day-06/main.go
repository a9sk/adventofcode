package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	rows := strings.Split(strings.TrimSpace(string(input)), "\n")
	h := len(rows)

	ops := []rune{}
	for _, f := range strings.Fields(rows[h-1]) {
		ops = append(ops, rune(f[0]))
	}

	var nums [][]int
	for _, row := range rows[:h-1] {
		fields := strings.Fields(row)
		arr := make([]int, len(fields))
		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			arr[i] = n
		}
		nums = append(nums, arr)
	}

	problems := len(nums[0])
	sol1 := 0

	for p := range problems {
		op := ops[p]
		result := nums[0][p]

		for r := 1; r < len(nums); r++ {
			if op == '+' {
				result += nums[r][p]
			} else {
				result *= nums[r][p]
			}
		}

		sol1 += result
	}

	fmt.Println("part 1:", sol1)
}
