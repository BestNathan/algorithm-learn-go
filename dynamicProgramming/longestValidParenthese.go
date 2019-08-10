package dynamicProgramming

import (
	"fmt"
)

// https://leetcode-cn.com/problems/longest-valid-parentheses/comments/

func LongestValidParentheses()  {
	fmt.Println("Longest Valid Parenthese: ")

	for _, s := range []string{"(", "()", "(()", "(())", "()(()())", ")()())"} {
		l := longestValidParentheses(s)
		fmt.Printf("The length of longest valid parenthese of string `%s` is %d\n", s, l)
	}

	fmt.Println()
}

func longestValidParentheses(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	dp := make(map[int]int)
	res := 0
	for i := 0; i < l; i++ {
		if s[i] != ')' || i - 1 < 0 {
			continue
		}

		if s[i - 1] == '(' {
			dp[i] = dp[i - 2] + 2
		} else if s[i - 1] == ')' && dp[i-1] > 0 && i - dp[i - 1] - 1 >= 0 && s[i - dp[i - 1] - 1] == '(' {
			dp[i] = dp[i - 1] + 2 + dp[i - dp[i - 1] - 2]
		}

		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}