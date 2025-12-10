package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aclements/go-z3/z3"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	sol1, sol2 := 0, 0

	ctx := z3.NewContext(nil)

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
		for _, b := range btns {
			var s int
			for _, x := range strings.Split(b[1:len(b)-1], ",") {
				v, _ := strconv.Atoi(x)
				s += 1 << v
			}
			bb = append(bb, s)
		}

		sc := len(btns) + 1
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
		if sc <= len(btns) {
			sol1 += sc
		}

		rawT := p[len(p)-1]
		rawT = rawT[1 : len(rawT)-1]
		tParts := strings.Split(rawT, ",")
		var tgts []int
		for _, x := range tParts {
			v, _ := strconv.Atoi(x)
			tgts = append(tgts, v)
		}

		slv := z3.NewSolver(ctx)
		intSort := ctx.IntSort()
		zero := ctx.FromInt(0, intSort).(z3.Int)
		one := ctx.FromInt(1, intSort).(z3.Int)

		var xx []z3.Int
		for k := range btns {
			x := ctx.IntConst(fmt.Sprintf("x%d", k))
			slv.Assert(x.GE(zero))
			xx = append(xx, x)
		}

		for j := range tgts {
			var ts []z3.Int
			for i := range btns {
				if (bb[i]>>j)&1 == 1 {
					ts = append(ts, xx[i])
				}
			}

			rhs := ctx.FromInt(int64(tgts[j]), intSort).(z3.Int)

			if len(ts) == 0 {
				if tgts[j] > 0 {
					slv.Assert(zero.Eq(one))
				}
			} else {
				sum := ts[0]
				for _, t := range ts[1:] {
					sum = sum.Add(t)
				}
				slv.Assert(sum.Eq(rhs))
			}
		}

		tot := ctx.IntConst("tot")
		if len(xx) > 0 {
			sumAll := xx[0]
			for _, x := range xx[1:] {
				sumAll = sumAll.Add(x)
			}
			slv.Assert(tot.Eq(sumAll))
		} else {
			slv.Assert(tot.Eq(zero))
		}

		mn := -1
		for {
			sat, err := slv.Check()
			if !sat || err != nil {
				break
			}

			m := slv.Model()
			res := m.Eval(tot, true)
			val, _, _ := res.(z3.Int).AsInt64()

			mn = int(val)

			cur := ctx.FromInt(val, intSort).(z3.Int)
			slv.Assert(tot.LT(cur))
		}

		if mn != -1 {
			sol2 += mn
		}
	}

	fmt.Println("part 1:", sol1)
	fmt.Println("part 2:", sol2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
