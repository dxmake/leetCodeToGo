package leetcode0015

import (
	"testing"
)

type question struct {
	q []int
	a [][]int
}

func TestThreeSum(t *testing.T) {
	showDetail := true
	qss := []question{
		question{
			q: []int{-1, 0},
			a: [][]int{},
		},
		question{
			q: []int{-1, 0, 0},
			a: [][]int{},
		},
		question{
			q: []int{-1, 0, 1},
			a: [][]int{
				[]int{-1, 0, 1},
			},
		},
		question{
			q: []int{1, 2, -2, -1},
			a: [][]int{},
		},
		question{
			q: []int{-1, -1, -1, 0, 0, 0, 1, 1, 1},
			a: [][]int{
				[]int{-1, 0, 1},
				[]int{0, 0, 0},
			},
		},
		question{
			q: []int{-1, 0, 1, 2, -1, -4},
			a: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
	}

	for _, qs := range qss {
		ans := threeSum(qs.q)
		// if ans == qs.a {
		// 	t.Logf("OK")
		// } else {
		// 	t.Logf("FAILED")
		// }
		if showDetail {
			t.Logf("  threeSum(%v) get %v, expect %v.", qs.q, ans, qs.a)
		}
	}
}
