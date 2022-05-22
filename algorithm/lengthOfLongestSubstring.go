package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	s1 := "bbbbb"
	s2 := "pwwkew"
	n := lengthOfLongestSubstring(s)
	n1 := lengthOfLongestSubstring(s1)
	n2 := lengthOfLongestSubstring(s2)
	fmt.Println(n, n1, n2)
}

func lengthOfLongestSubstring(s string) int {
	start, tmp, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		tmp = i - start + 1
		for j := start; j < i; j++ {
			if s[i] == s[j] {
				tmp = i - start
				start = j + 1
				break
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
