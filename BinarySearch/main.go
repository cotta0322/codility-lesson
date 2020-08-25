package main

import (
	"fmt"
	"strings"
)

func numAlphabet(searchStr string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetList := strings.Split(alphabet, "")

	low := 0
	high := len(alphabetList) - 1
	for low <= high {
		mid := (high + low) / 2
		if alphabetList[mid] == searchStr {
			return mid + 1
		}
		if alphabetList[mid] > searchStr {
			high = mid - 1
		}
		if alphabetList[mid] < searchStr {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(numAlphabet("y"))
}
