package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	nums1 := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums), singleNumber(nums1))
}

func singleNumber(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp ^= nums[i]
	}
	return tmp
}
