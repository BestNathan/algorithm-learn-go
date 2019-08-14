package dynamicProgramming

import (
	"fmt"
)

func LongestPalindrome()  {
	fmt.Println("Longest Palindrome: ")
	for _, s := range []string{"aba", "aa", "ababa", "abcdcb", "aaaa"} {
		l := longestPalindrome(s)
		fmt.Printf("longest palindrome of `%s` is `%s`\n", s, l)
	}
	fmt.Println()
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	dp := make([][]bool, l)
	for dpi := 0; dpi < l; dpi++ {
		dp[dpi] = make([]bool, l)
	}

	len := 1
	res := string(s[0])
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			// fmt.Println(s[i], s[j], i, j)
			if s[i] == s[j] && (j - i <=2 || dp[i + 1][j - 1]) {
				dp[i][j] = true
				ll := j - i + 1
				if ll > len {
					len = ll
					res = string(s[i:j + 1])
				}
			}
		}
	}
	return res
}