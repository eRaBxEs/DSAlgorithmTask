package main

import (
	"fmt"
)

func firstMissingPositiveEasy(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 1
	}

	seen := make([]bool, n+1)

	// step 1: loop through the array nums to mark available numbers as index in seen array
	for _, num := range nums {
		if num > 0 && num <= n {
			seen[num] = true
		}
	}

	// step 2: loop through the num array starting from index 1 to find the number not existing in seen array, we using the index of seen as a checker
	for i := 1; i <= n; i++ {
		if !seen[i] {
			return i
		}
	}

	// if all are seen return n+1
	return n + 1
}

func main() {
	tests := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
		{1, 1},
		{},
		{-1, -2},
		{2, 1},
	}

	for _, arr := range tests {
		fmt.Printf("%v → %d\n", arr, firstMissingPositiveEasy(arr))
	}
}
