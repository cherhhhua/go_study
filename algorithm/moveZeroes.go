package main

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}
	//两个指针i和j
	j := 0
	for i := 0; i < len(nums); i++ {
		//当前元素!=0，就把其交换到左边，等于0的交换到右边
		if nums[i] != 0 {
			if i > j { // #1
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}
