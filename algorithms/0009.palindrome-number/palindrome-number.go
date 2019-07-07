package leetcode0009

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// I see a solution has exactly same solution with me except he uses []rune instead of string.
	// His Code cost 0ms and mine cost 20ms.
	// It's somewhere we chan go deep in to.
	// s := []rune(strconv.Itoa(x))
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
