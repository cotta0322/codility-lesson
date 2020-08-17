package main

import (
	"fmt"
)

func Solution(X int, A []int) int {
	remaining := map[int]int{}

	for i := 0; i < X+1; i++ {
		remaining[i] = -1
	}

	for i := range A {
		if remaining[A[i]] == -1 {
			remaining[A[i]] = i
		}
	}

	rel := 0
	for i := 1; i < X+1; i++ {
		if remaining[i] == -1 {
			return -1
		}
		if rel < remaining[i] {
			rel = remaining[i]
		}
	}
	return rel
}

func main() {
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	fmt.Println(Solution(X, A))
}
