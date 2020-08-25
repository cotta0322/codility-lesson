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
			rel = rel + 2
		}
	}

	if int(sqrt)*int(sqrt) == N {
		rel--
	}

	return rel
}

func main() {
	N := 24

	fmt.Println(Solution(N))
}
