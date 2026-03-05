package main

import (
	"fmt"
)

/**
1. Array Rotation
Rotate an array to the left by K steps.

Example:

Input: arr = [1,2,3,4,5], K = 2

Output: [3,4,5,1,2] (rotated left by 2 positions)

**/

func RotateLeft(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}

	k = k % len(arr)
	if k == 0 {
		return arr
	}

	return append(arr[k:], arr[:k]...)
}

func main() {
	arrayOne := []int{3, 4, 5, 7, 9}
	fmt.Printf("%#v\n", arrayOne)
	rotatedArray := RotateLeft(arrayOne, 2)
	fmt.Printf("%#v\n", rotatedArray)
}
