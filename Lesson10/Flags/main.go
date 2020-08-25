package main

import (
	"fmt"
)

func findPeek(A []int) []int {
	rel := []int{}
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			rel = append(rel, i)
		}
	}
	return rel
}

func Solution(A []int) int {
	rel := 0
	peeks := findPeek(A)

	if len(peeks) == 0 {
		return 0
	}

	sum := []int{peeks[0]}
	for i := 1; i < len(peeks); i++ {
		sum = append(sum, sum[i-1]+peeks[i])
	}

	for k := 0; k < len(peeks); k++ {
		tmp := 0
		nokoriK := k
		for i := range sum {
			tmp += sum[i]
			if k < tmp {
				nokoriK--
				tmp = 0
			}
			if nokoriK == 0 {
				rel = k
			}
		}
	}

	return rel
}

func main() {
	A := []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}

	fmt.Println(Solution(A))
}
