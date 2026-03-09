package main

import (
	"fmt"
)

/**
3. Cyclic Rotation
Rotate an array to the right by K steps.

Example:

Input: arr = [3,8,9,7,6], K = 3

Output: [9,7,6,3,8]
**/

func CyclicRotation(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}

	k = k % len(arr)
	if k == 0 {
		return arr
	}

	return append(arr[len(arr)-k:], arr[:len(arr)-k]...)
}

func main() {
	fmt.Println("CyclicRotation")
	arr := []int{5, 6, 1, 4, 2}
	k := 3
	// should give {1, 4, 2, 5, 6}
	fmt.Printf("array #%v with cyclic rotation with %d steps give #%v", arr, k, CyclicRotation(arr, k))

}
