package main

import (
	"fmt"
)

func sumArray(A []int) int {
	rel := 0
	for i := range A {
		rel += A[i]
	}
	return rel
}

func Solution(A []int) int {
	left := 0
	right := sumArray(A)
	rel := 0
	for i := 0; i < len(A)-1; i++ {
		left = left + A[i]
		right = right - A[i]

		tmp := left - right

		if tmp == 0 {
			rel = tmp
			return 0
		}
		if tmp < 0 {
			tmp = -tmp
		}
		if rel == 0 || rel > tmp {
			rel = tmp
		}
	}
	return rel
}

func main() {
	A := []int{3, 1, 2, 4, 3}
	fmt.Println(Solution(A))
}
