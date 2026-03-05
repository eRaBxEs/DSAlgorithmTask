package main

import (
	"fmt"
)

func BinaryGap(n int) int {
	for n > 0 && n&1 == 0 {
		n >>= 1
	}

	maxGap, currentGap := 0, 0
	for n > 0 {
		if n&1 == 0 {
			currentGap++
		} else {
			if currentGap > maxGap {
				maxGap = currentGap
			}
			currentGap = 0
		}

		n >>= 1
	}
	return maxGap

}

func main() {
	p := 1041
	q := 65
	r := 32
	s := 9

	fmt.Printf("binary of %d = %b and the longest binary gap of 0's in between two 1's =%d\n", p, p, BinaryGap(p))

	fmt.Printf("binary of %d = %b and the longest binary gap of 0's in between two 1's =%d\n", q, q, BinaryGap(q))

	fmt.Printf("binary of %d = %b and the longest binary gap of 0's in between two 1's =%d\n", r, r, BinaryGap(r))

	fmt.Printf("binary of %d = %b and the longest binary gap of 0's in between two 1's =%d\n", s, s, BinaryGap(s))

}
