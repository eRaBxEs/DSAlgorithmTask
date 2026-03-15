package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// step 1: set all numbers <= 0 and numbers > n to n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	// step 2: check for numbers <= n and set numbers at index n - 1 to negative
	for i := 0; i < n; i++ {
		val := abs(nums[i])

		if val <= n {
			if nums[val-1] > 0 {
				nums[val-1] = -nums[val-1] // negating such numbers
			}
		}
	}

	// step 3: return the index i+1 of any number that is still positive; that is the number
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	// if non of the numbers are positive, return n+1 as the number
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		fmt.Printf("%v → %d\n", arr, firstMissingPositive(arr))
	}
}
