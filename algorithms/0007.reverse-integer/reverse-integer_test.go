package leetcode0007

import (
	"testing"
)

type question struct {
	q int
	a int
}

func TestReverse(t *testing.T) {
	const showDetail bool = true
	qss := []question{
		question{q: 0, a: 0},

		question{q: 1, a: 1},
		question{q: -1, a: -1},

		question{q: 120, a: 21},
		question{q: -120, a: -21},

		question{q: 123, a: 321},
		question{q: -123, a: -321},

		question{q: 2147483647, a: 0},
		question{q: -2147483647, a: 0},

		question{q: 2147483648, a: 0},
		question{q: -2147483648, a: 0},

		question{q: 7463847412, a: 2147483647},
		question{q: 8463847412, a: 0},

		question{q: -8463847412, a: -2147483648},
		question{q: -9463847412, a: 0},

		question{q: 2000000001, a: 1000000002},
		question{q: -2000000001, a: -1000000002},

		question{q: 12345678901, a: 0},
		question{q: -12345678901, a: 0},
	}

	for _, qs := range qss {
		ans := reverse(qs.q)
		t.Log(ans == qs.a)
		if showDetail {
			t.Logf("  reverse(%d) get %d, and expect %d.", qs.q, ans, qs.a)
		}
	}
}
