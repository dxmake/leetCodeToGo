package leetcode0007

import (
	// "log"
	"strconv"
)

func reverse(x int) int {
	ru := []rune(strconv.Itoa(x))
	var r []rune

	if ru[0] == '-' {
		r = ru[1:]
	} else {
		r = ru
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	s := string(ru)
	y, _ := strconv.Atoi(s)

	const minInt int = -2147483648
	const maxInt int = 2147483647
	if y < minInt || y > maxInt {
		return 0
	}
	return y
}
