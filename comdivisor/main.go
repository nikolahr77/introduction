package main

import "fmt"

func greatestDivisor(a, b int) int {
	if a == 0 {
		return b
	}

	for b != 0 {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func main() {
	fmt.Print(greatestDivisor(63, 9))
}
