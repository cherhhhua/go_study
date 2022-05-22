package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums1 := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums), maxSubArray(nums1))
}

func maxSubArray(nums []int) int {
	s := 0
	ans := nums[0]
	for _, n := range nums {
		if s > 0 {
			s += n
		} else {
			s = n
		}
		ans = max(ans, s)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
