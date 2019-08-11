package dynamicProgramming

import (
	"fmt"
)

func IsMatch()  {
	fmt.Println("is match: ")
	for _, items := range [][]string{
		[]string{"abc", "a*c"},
		[]string{"aa", "a"},
		[]string{"aa", "a*"},
		[]string{"aa", "a?"},
	} {
		m := isMatch(items[0], items[1])
		fmt.Printf("`%s` match `%s`: %v\n", items[0], items[1], m)
	}
	fmt.Println()
}

// 定义状态:
// dp[i][j] 表示 s 的前 i 个字符 和 p 的前 j 个字符是否匹配
// 初始状态:
// 空 匹配 空
// dp[0][0] = true
// s 为空 则只有 p 为 * 的时候匹配
// dp[0][j] = dp[0][j-1] && p[j] == '*'
// p 为空 都不匹配
// dp[i][0] = false
// 状态转移
// 如果 p[j] 不是 *
// dp[i][j] = dp[i-1][j-1] && (s[i] == p[j] || p[j] == '?')
// 如果 p[j] 是 *
// 	可以选择匹配也可以选择不匹配
// 	

func isMatch(s, p string) bool {
	ls := len(s)
	lp := len(p)

	dp := make([][]bool, ls + 1)
	for dpi := 0; dpi < ls + 1; dpi++ {
		dp[dpi] = make([]bool, lp + 1)
	}

	dp[0][0] = true

	// s 空串的匹配
	for j := 1; j <= lp; j++ {
		dp[0][j] = dp[0][j - 1] && p[j - 1] == '*'
	}

	// p 空串不匹配
	for i := 1; i <= ls; i++ {
		dp[i][0] = false
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if s[i - 1] == p[j - 1] || p[j - 1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j - 1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j - 1]
			}
		}
	}
	return dp[ls][lp]
}