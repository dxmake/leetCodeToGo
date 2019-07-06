package leetcode0004

// import (
// 	"log"
// )

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len2 < len1 {
		nums1, nums2 = nums2, nums1
		len1, len2 = len2, len1
	}

	lenAll := len1 + len2
	lenHalf := lenAll / 2
	// [lenHalf] numbers in the left part and [lenAll-lenHalf] numbers in the right part.
	// (lenAll>=2) makes sure there's at least one number in both left and half part
	if lenAll == 1 {
		switch {
		case len1 != 0:
			return float64(nums1[0])
		case len2 != 0:
			return float64(nums2[0])
		}
	}

	st1, end1 := 0, len1
	st2, end2 := 0, len2

	for {
		// 迭代 折半查找
		mid1 := (end1 + st1) / 2
		mid2 := lenHalf - mid1
		switch {
		case mid2 < st2:
			mid1 += mid2
			end1 = mid1
			mid2 = st2
		case mid2 > end2:
			mid1 += mid2 - len2
			end1 = mid1
			mid2 = end2
		default:
		}
		// log.Print((mid1 == 0 || mid2 == len2 || nums1[mid1-1] <= nums2[mid2]))
		// log.Print((mid1 == len1 || mid2 == 0 || nums2[mid2-1] <= nums1[mid1]))
		// log.Fatal(st1, end1, mid1, mid2)
		// 退出条件
		if (mid1 == 0 || mid2 == len2 || nums1[mid1-1] <= nums2[mid2]) && (mid1 == len1 || mid2 == 0 || nums2[mid2-1] <= nums1[mid1]) {
			var maxL, minR float64
			// Both left and right has a number.
			switch {
			case mid1 == 0:
				maxL = float64(nums2[mid2-1])
			case mid2 == 0:
				maxL = float64(nums1[mid1-1])
			default:
				maxL = float64(nums1[mid1-1])
				if float64(nums2[mid2-1]) > maxL {
					maxL = float64(nums2[mid2-1])
				}
			}
			switch {
			case mid1 == len1:
				minR = float64(nums2[mid2])
			case mid2 == len2:
				minR = float64(nums1[mid1])
			default:
				minR = float64(nums1[mid1])
				if float64(nums2[mid2]) < minR {
					minR = float64(nums2[mid2])
				}
			}
			// log.Print(lenAll&0x1 == 1)
			// log.Fatal(maxL, minR)
			if lenAll&0x1 == 1 {
				return minR
			} else {
				return (maxL + minR) / 2
			}
		}
		// 迭代过程
		switch {
		case !(mid1 == 0 || mid2 == len2 || nums1[mid1-1] <= nums2[mid2]):
			end1 = mid1
		case !(mid1 == len1 || mid2 == 0 || nums2[mid2-1] <= nums1[mid1]):
			st1 = mid1 + 1
		}

	}
}
