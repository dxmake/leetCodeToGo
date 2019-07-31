package leetcode0070

import (
	"testing"
)

type question struct {
	q int
	a int
}

func TestClimbStairs(t *testing.T) {
	qss := []question{
		question{1, 1},
		question{2, 2},
		question{3, 3},
		question{4, 5},
		question{5, 8},
		question{6, 13},
		question{7, 21},
		question{8, 34},
		question{9, 55},
		question{10, 89},
	}
	for _, qs := range qss {
		ans := climbStairs(qs.q)
		if ans != qs.a {
			t.Errorf("FAILED: climbStairs(%d) get %d, expect %d.", qs.q, ans, qs.a)
		} else {
			t.Logf("OK: climbStairs(%d) get %d.", qs.q, ans)
		}
	}
}
