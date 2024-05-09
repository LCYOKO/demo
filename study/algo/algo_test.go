package algo

import (
	"fmt"
	"testing"
)

func isPalindrome(x int) bool {
	str := string(rune(x))
	for i, j := 0, len(str); i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func TestAlgo1(t *testing.T) {
	fmt.Println(longestPalindrome("aaaaa"))
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, n+1)
	}
	startIdx, maxLen := -1, 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if s[i-1] == s[j-1] && (i-j <= 2 || dp[j+1][i-1]) {
				if i-j+1 > maxLen {
					maxLen = i - j + 1
					startIdx = j - 1
				}
				dp[j][i] = true
			}
		}
	}
	if -1 == startIdx {
		return ""
	} else {
		return s[startIdx : startIdx+maxLen]
	}
}
