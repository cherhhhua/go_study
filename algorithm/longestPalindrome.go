package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return ""
	}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		res := max(len1, len2)

		if res > end-start {
			start = i - (res-1)/2
			end = i + res/2
		}

	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
