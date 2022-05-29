package main

func canJump(nums []int) bool {
	//last表示的是能不能到达last这个位置
	last := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		//从倒数第2个位置往前遍历，如果当前位置能够跳
		//到last这个位置，就更新last，如果从当前位置
		//不能到达last这个位置就继续往前遍历
		if i+nums[i] >= last {
			last = i
		}
	}
	//如果last等于0，说明可以从第一个位置跳到最后
	return last == 0
}
