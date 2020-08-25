package main

import (
	"fmt"
)

func Solution(A []int) int {
	rel := A[0]
	current := A[0]
	for i := 1; i < len(A); i++ {
		tmp := current + A[i]
		if tmp > A[i] {
			current = tmp
		} else {
			current = A[i]
		}
		if rel < current {
			rel = current
		}
	}

	return rel
}

func main() {
	A := []int{3, 2, -6, 4, 0}

	fmt.Println(Solution(A))
}
