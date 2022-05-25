package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := MakeBoolSlice(m+1, n+1)

	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j-2]
				if matches(s, p, i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else {
				if matches(s, p, i, j) {
					f[i][j] = f[i-1][j-1]
				}
			}
		}
	}
	return f[m][n]
}

func MakeBoolSlice(len1, len2 int) [][]bool {
	slice := make([][]bool, len1)
	for i := range slice {
		slice[i] = make([]bool, len2)
	}
	return slice
}
func matches(s, p string, i, j int) bool {
	if i == 0 {
		return false
	}

	if p[j-1] == '.' {
		return true
	}

	return s[i-1] == p[j-1]
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))

}
