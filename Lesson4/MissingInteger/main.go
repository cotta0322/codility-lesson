package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)

	if A[0] > 0 && A[0] != 1 {
		return 1
	}
	if A[len(A)-1] <= 0 {
		return 1
	}
	for i := 1; i < len(A); i++ {
		if A[i] == 1 {
			continue
		}
		if A[i] <= 0 {
			continue
		}
		if A[i-1] <= 0 && A[i] != 1 {
			return 1
		}
		if A[i-1] == A[i] || A[i-1]+1 == A[i] {
			continue
		}
		return A[i-1] + 1
	}
	return A[len(A)-1] + 1
}

func main() {
	A := []int{0, 1, 2, 100}
	fmt.Println(Solution(A))
}
