package leetcode0011

func maxArea(height []int) int {
	maxA := 0
	tmpA := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			tmpA = height[i] * (j - i)
			i++
		} else {
			tmpA = height[j] * (j - i)
			j--
		}
		if tmpA > maxA {
			maxA = tmpA
		}
	}
	return maxA
}
