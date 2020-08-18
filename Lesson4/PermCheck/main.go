package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)

	if len(A) == 1 {
		if A[0] != 1 {
			return 0
		}
		return 1
	}

	if A[0] != 1 {
		return 0
	}
	for i := 1; i < len(A); i++ {
		if A[i-1]+1 != A[i] {
			return 0
		}
	}
	return 1
}

func main() {
	A := []int{4, 1, 3, 2}
	fmt.Println(Solution(A))
}
