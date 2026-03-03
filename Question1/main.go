package main

import (
	"fmt"
)

func rotateLeft(arr []int, k int) []int {
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
	rotatedArray := rotateLeft(arrayOne, 2)
	fmt.Printf("%#v\n", rotatedArray)
}
