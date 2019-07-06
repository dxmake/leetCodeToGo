package leetcode0004

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{}, []int{2}) == 2)
	t.Log(findMedianSortedArrays([]int{2}, []int{}) == 2)
	t.Log(findMedianSortedArrays([]int{1}, []int{2}) == 1.5)
	t.Log(findMedianSortedArrays([]int{2}, []int{1}) == 1.5)
	t.Log(findMedianSortedArrays([]int{1, 3}, []int{2}) == 2)
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4}) == 2.5)
	t.Log(findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 4}) == 4)
	t.Log(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{7, 8}) == 4.5)
	t.Log(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{7}) == 4)
	t.Log(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{}) == 3.5)
	t.Log(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{-5}) == 3)
	t.Log(findMedianSortedArrays([]int{1, 2, 3, 4, 5, 6}, []int{-5, -4}) == 2.5)
	t.Log(findMedianSortedArrays([]int{}, []int{5, 6, 7, 8, 9}) == 7)
	t.Log(findMedianSortedArrays([]int{1}, []int{5, 6, 7, 8, 9}) == 6.5)
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{5, 6, 7, 8, 9}) == 6)
}
