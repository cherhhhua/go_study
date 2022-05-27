package main

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		swapNum(nums, i, j)
	}
	reverse(nums, i+1)
}

func swapNum(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start int) {
	left := start
	right := len(nums)

	for left < right {
		swapNum(nums, left, right)
		left++
		right--
	}
}
