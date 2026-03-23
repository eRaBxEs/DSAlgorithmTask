package main

import (
	"fmt"
)

func IsPermutation(arr []int) int {
	seen := make([]bool, len(arr)+1)
	for _, num := range arr {
		if num < 1 || num > len(arr) || seen[num] {
			return 0
		}

		seen[num] = true
	}

	fmt.Println(seen)

	return 1
}

func main() {
	arrOne, arrTwo := []int{5, 2, 3, 1, 4}, []int{2, 3, 5, 1}
	fmt.Println("checking if the arrays/slice above are permutative")
	fmt.Println("an output of 1 means it is permutative while an output of 0 means it's not")
	fmt.Printf("array:: %#v :: %d\n", arrOne, IsPermutation(arrOne))
	fmt.Printf("array:: %#v :: %d\n", arrOne, IsPermutation(arrTwo))
}
