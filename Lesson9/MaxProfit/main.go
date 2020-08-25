package main

import (
	"fmt"
)

func Solution(A []int) int {
	max := 0
	current := 0
	for i := 1; i < len(A); i++ {
		diff := A[i] - A[i-1]
		if current <= 0 {
			current = diff
		} else {
			current = current + diff
		}
		if current > max {
			max = current
		}
	}
	return max
}

func main() {
	A := []int{23171, 21011, 21123, 21366, 21013, 21367}

	fmt.Println(Solution(A))
}
