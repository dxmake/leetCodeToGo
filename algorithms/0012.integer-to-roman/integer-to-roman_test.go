package leetcode0012

import (
	"testing"
)

type question struct {
	q int
	a string
}

func TestIntToRoman(t *testing.T) {
	showDetail := true
	qss := []question{
		question{1, "I"},
		question{2, "II"},
		question{3, "III"},
		question{4, "IV"},
		question{5, "V"},
		question{6, "VI"},
		question{7, "VII"},
		question{8, "VIII"},
		question{9, "IX"},
		question{10, "X"},
		question{58, "LVIII"},
		question{909, "CMIX"},
		question{3089, "MMMLXXXIX"},
		question{3888, "MMMDCCCLXXXVIII"},
		question{3999, "MMMCMXCIX"},
	}

	for _, qs := range qss {
		ans := intToRoman(qs.q)
		if ans == qs.a {
			t.Logf("OK")
		} else {
			t.Logf("FAILED")
		}
		if showDetail {
			t.Logf("  intToRoman(%d) get \"%s\", expect \"%s\".", qs.q, ans, qs.a)
		}
	}
}
