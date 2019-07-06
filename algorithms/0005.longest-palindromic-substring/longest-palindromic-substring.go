package leetcode0005

func longestPalindrome(s string) string {
	len1 := len(s)
	pNum := make([]int, 2*len1)
	pMax := 0
	pPos := 0
	for i, j := 0, 0; j < len1-i; j++ {
		pNum[2*j+i] = 1
		pMax = 1
	}
	for i := 1; i < len1; i++ {
		for j := 0; j < len1-i; j++ {
			if s[j] == s[j+i] {
				pNum[2*j+i] += 2
				if pNum[2*j+i] == i+1 && pNum[2*j+i] > pMax {
					pMax = pNum[2*j+i]
					pPos = j
				}
			}
		}
	}
	return s[pPos : pPos+pMax]
}
