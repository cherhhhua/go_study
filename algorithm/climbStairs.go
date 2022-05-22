package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	a := 1
	b := 1
	for i := 1; i < n; i++ {
		a = a + b
		b = a - b
	}
	return a
}
