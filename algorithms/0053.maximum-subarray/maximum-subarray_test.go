package leetcode0053

import (
	"testing"
)

type question struct {
	q []int
	a int
}

func TestMaxSubArray(t *testing.T) {
	qss := []question{
		question{[]int{0}, 0},
		question{[]int{-1}, -1},
		question{[]int{1}, 1},
		question{[]int{-1, 0, 1}, 1},
		question{[]int{1, 0, -1}, 1},
		question{[]int{1, -1, 1, 2}, 3},
		question{[]int{-1, -2, -3}, -1},
		question{[]int{-3, -2, -1}, -1},
		question{[]int{10, -5, 1, 2, 3}, 11},
		question{[]int{-1, 2, -1, 2, -1, 2}, 4},
		question{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, qs := range qss {
		ans := maxSubArray(qs.q)
		if ans != qs.a {
			t.Errorf("FAILED: maxSubArray(%v) get %d, expect %d.", qs.q, ans, qs.a)
		} else {
			t.Logf("OK: maxSubArray(%v) get %d.", qs.q, ans)
		}
	}
}
