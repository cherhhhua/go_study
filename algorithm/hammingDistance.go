package main

func hammingDistance(x int, y int) int {
	xor := x ^ y
	count := 0
	for xor != 0 {
		xor = xor & (xor - 1)
		count++
	}
	return count
}
