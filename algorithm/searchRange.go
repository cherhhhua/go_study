package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	lower := lower_bound(nums, target)
	upper := upper_bound(nums, target) - 1
	if lower == len(nums) || nums[lower] != target {
		return []int{-1, -1}
	}
	return []int{lower, upper}
}

func lower_bound(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
func upper_bound(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
