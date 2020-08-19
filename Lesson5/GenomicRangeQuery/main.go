package main

import (
	"fmt"
	"strings"
)

func Solution(S string, P []int, Q []int) []int {
	rel := []int{}

	Sstr := strings.Split(S, "")
	fmt.Println(Sstr)

	a := []int{0}
	c := []int{0}
	g := []int{0}
	t := []int{0}
	for i := 0; i < len(Sstr); i++ {
		switch Sstr[i] {
		case "A":
			a = append(a, a[i]+1)
			c = append(c, c[i])
			g = append(g, g[i])
			t = append(t, t[i])
		case "C":
			a = append(a, a[i])
			c = append(c, c[i]+1)
			g = append(g, g[i])
			t = append(t, t[i])
		case "G":
			a = append(a, a[i])
			c = append(c, c[i])
			g = append(g, g[i]+1)
			t = append(t, t[i])
		case "T":
			a = append(a, a[i])
			c = append(c, c[i])
			g = append(g, g[i])
			t = append(t, t[i]+1)
		}
	}

	M := len(P)
	for i := 0; i < M; i++ {
		p := P[i]
		q := Q[i]

		// Aあるか？
		fmt.Println(a)
		if a[p] < a[q+1] {
			rel = append(rel, 1)
			continue
		}
		// Cあるか？
		if c[p] < c[q+1] {
			rel = append(rel, 2)
			continue
		}
		// Gあるか？
		if g[p] < g[q+1] {
			rel = append(rel, 3)
			continue
		}
		// Tあるか？
		if t[p] < t[q+1] {
			rel = append(rel, 4)
			continue
		}
	}

	return rel
}

func main() {
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}

	fmt.Println(Solution(S, P, Q))
}
