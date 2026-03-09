package main

import (
	"fmt"
)

/**
A frog wants to reach position Y from position X, jumping exactly D units each time. Calculate the minimum number of jumps needed.

Example:

Input: X = 10, Y = 85, D = 30

Output: 3 (10 → 40 → 70 → 100)
**/

func FrogJump(X, Y, D int) int {
	distance := Y - X
	if distance <= 0 {
		return 0
	}

	// jumps
	jumps := distance / D
	if distance%D != 0 { // check for remainders
		jumps++
	}

	return jumps
}

func main() {
	x, y, d := 10, 85, 30
	fmt.Printf("Given X = %d\n", x)
	fmt.Printf("Given Y = %d\n", y)
	fmt.Printf("Given D = %d\n", d)
	fmt.Printf("Minimal number of jumps needed = %d\n", FrogJump(x, y, d))
}
