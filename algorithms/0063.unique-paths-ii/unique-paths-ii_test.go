package leetcode0063

import (
	"testing"
)

type question struct {
	q [][]int
	a int
}

func TestUniquePathsWithObstacles(t *testing.T) {
	qss := []question{
		question{
			[][]int{
				[]int{0},
			},
			1,
		},
		question{
			[][]int{
				[]int{1},
			},
			0,
		},
		question{
			[][]int{
				[]int{0, 0},
				[]int{0, 0},
			},
			2,
		},
		question{
			[][]int{
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			1,
		},
		question{
			[][]int{
				[]int{0, 1, 0},
				[]int{0, 1, 0},
			},
			0,
		},
		question{
			[][]int{
				[]int{0, 1, 0},
				[]int{0, 0, 0},
				[]int{1, 0, 1},
				[]int{0, 0, 0},
			},
			1,
		},
	}
	for _, qs := range qss {
		ans := uniquePathsWithObstacles(qs.q)
		if ans != qs.a {
			t.Errorf("FAILED: Get %d, expect %d.", ans, qs.a)
		} else {
			t.Logf("OK: Get %d.", ans)
		}
	}
}
