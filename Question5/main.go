package main

import (
	"fmt"
)

/*
Odd Occurrences In Array

Find the element that appears an odd number of times in an array. All other elements appear an even number of times.
Example:
· Input: [9,3,9,3,9,7,9]
· Output: 7
*/

func OddOccurrences(arr []int) int {
	result := 0
	for _, val := range arr {
		result ^= val
	}
	return result
}

func main() {
	arrOne, arrTwo, arrThree := []int{9, 3, 9, 3, 9, 7, 9}, []int{9, 3, 9, 9, 3, 5, 7, 5, 5, 5, 7}, []int{2, 2, 3, 3, 1, 7, 5, 7, 5}

	fmt.Printf("Odd occurrence in %#v\n is %d\n", arrOne, OddOccurrences(arrOne))
	fmt.Printf("Odd occurrence in %#v\n is %d\n", arrTwo, OddOccurrences(arrTwo))
	fmt.Printf("Odd occurrence in %#v\n is %d\n", arrThree, OddOccurrences(arrThree))

}
