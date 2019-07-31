package leetcode0064

import (
	"testing"
)

type question struct {
	q [][]int
	a int
}

func TestMinPathSum(t *testing.T) {
	qss := []question{
		question{
			[][]int{
				[]int{0},
			},
			0,
		},
		question{
			[][]int{
				[]int{1},
			},
			1,
		},
		question{
			[][]int{
				[]int{0, -1},
				[]int{0, 0},
			},
			-1,
		},
		question{
			[][]int{
				[]int{0, -1, 1},
				[]int{0, -2, 0},
			},
			-3,
		},
		question{
			[][]int{
				[]int{0, 1, -1},
				[]int{0, 1, 0},
			},
			0,
		},
		question{
			[][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			7,
		},
	}
	for _, qs := range qss {
		ans := minPathSum(qs.q)
		if ans != qs.a {
			t.Errorf("FAILED: Get %d, expect %d.", ans, qs.a)
		} else {
			t.Logf("OK: Get %d.", ans)
		}
	}
}
