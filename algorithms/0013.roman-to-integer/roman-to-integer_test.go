package leetcode0013

import (
	"testing"
)

type question struct {
	q string
	a int
}

func TestRomanToInt(t *testing.T) {
	showDetail := true
	qss := []question{
		question{"I", 1},
		question{"II", 2},
		question{"III", 3},
		question{"IV", 4},
		question{"V", 5},
		question{"VI", 6},
		question{"VII", 7},
		question{"VIII", 8},
		question{"IX", 9},
		question{"X", 10},
		question{"LVIII", 58},
		question{"CMIX", 909},
		question{"MMMLXXXIX", 3089},
		question{"MMMDCCCLXXXVIII", 3888},
		question{"MMMCMXCIX", 3999},
	}

	for _, qs := range qss {
		ans := romanToInt(qs.q)
		if ans == qs.a {
			t.Logf("OK")
		} else {
			t.Logf("FAILED")
		}
		if showDetail {
			t.Logf("  romanToInt(\"%s\") get %d, expect %d.", qs.q, ans, qs.a)
		}
	}
}
