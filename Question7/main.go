package main

import (
	"fmt"
)

func DistinctCounter(arr []int) int {
	seen := make(map[int]bool)
	for _, num := range arr {
		seen[num] = true
	}

	return len(seen)
}

func main() {
	arrOne := []int{2, 1, 1, 2, 3, 1}
	fmt.Printf("array :: %#v has %d distinct values\n", arrOne, DistinctCounter(arrOne))
}
