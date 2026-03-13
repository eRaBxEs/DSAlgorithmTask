package main

import (
	"fmt"
)

/**
8. Max Counters
You have N counters, initially set to 0. Two operations:

Increase(X): counter X increases by 1
Max Counter: all counters are set to the maximum value of any counter

Given a sequence of operations, return the final counter values.
Example:

Input: N = 5, operations = [3,4,4,6,1,4,4]
Output: [3,2,2,4,2]
* */

func MaxCounters(N int, operations []int) []int {
	lastMax := 0
	maxCounter := 0
	counters := make([]int, N)

	for _, ops := range operations {
		if ops >= 1 && ops <= N {
			idx := ops - 1

			if counters[idx] < lastMax {
				counters[idx] = lastMax
			}

			counters[idx]++

			// update maxCounter
			if counters[idx] > maxCounter {
				maxCounter = counters[idx]
			}
		} else if ops == N+1 {
			lastMax = maxCounter
		}

	}

	// lazy update
	for i := 0; i < N; i++ {
		if counters[i] < lastMax {
			counters[i] = lastMax
		}
	}
	return counters
}

func main() {
	operations := []int{3, 4, 4, 6, 1, 4, 4}
	N := 5
	fmt.Printf("Input: %#v for %d Counters: %#v", operations, N, MaxCounters(N, operations))
}
