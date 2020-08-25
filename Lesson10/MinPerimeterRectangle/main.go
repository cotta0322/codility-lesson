package main

import (
	"fmt"
	"math"
)

func Solution(N int) int {
	rel := 0
	sqrt := math.Sqrt(float64(N))
	for i := 1; i <= int(sqrt); i++ {
		if N%i == 0 {
			perimeter := 2 * (i + N/i)
			if rel == 0 || rel > perimeter {
				rel = perimeter
			}
		}
	}
	return rel
}

func main() {
	N := 30

	fmt.Println(Solution(N))
}
