package leetcode0006

import (
	// "log"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	len1 := len(s)
	if len1 == 0 {
		return ""
	}
	numRepeat := 2*numRows - 2
	numPrds := len1 / numRepeat
	numOver := len1 % numRepeat
	// log.Print(numRepeat, numPrds, numOver)
	var pattrn []int
	var res strings.Builder
	res.Grow(len1)
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			pattrn = []int{i}
		} else {
			pattrn = []int{i, numRepeat - i}
		}
		// log.Print(i, numRows, pattrn)
		for j := 0; j < numPrds; j++ {
			// log.Print("    ", j, numPrds)
			for k := 0; k < len(pattrn); k++ {
				// log.Print("        ", k, len(pattrn), string(s[numRepeat*j+pattrn[k]]))
				res.WriteByte(s[numRepeat*j+pattrn[k]])
			}
		}
		for k := 0; k < len(pattrn); k++ {
			if pattrn[k] < numOver {
				res.WriteByte(s[numRepeat*numPrds+pattrn[k]])
			}
		}
	}
	return res.String()
}
