package leetcode0011

import (
	"testing"
)

type question struct {
	q []int
	a int
}

func TestMaxArea(t *testing.T) {
	showDetail := true
	qss := []question{
		question{[]int{0, 1}, 0},
		question{[]int{1, 0}, 0},
		question{[]int{1, 1}, 1},
		question{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, qs := range qss {
		ans := maxArea(qs.q)
		if ans == qs.a {
			t.Logf("OK")
		} else {
			t.Logf("FAILED")
		}
		if showDetail {
			t.Logf("  maxArea(%v) get %d, expect %d.", qs.q, ans, qs.a)
		}
	}
}
