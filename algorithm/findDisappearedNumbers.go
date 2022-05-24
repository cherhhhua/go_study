package main

func findDisappearedNumbers(nums []int) []int {
	var res []int

	n := len(nums)
	for _, num := range nums {
		x := (num - 1) % n
		nums[x] += n
	}
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res
}
