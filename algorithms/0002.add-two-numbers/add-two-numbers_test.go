package leetcode0002

import (
	"testing"
)

type question struct {
	q1  *ListNode
	q2  *ListNode
	ans *ListNode
}

func makeList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
	head := &ListNode{}
	p := head
	for i := 0; i < len(l); i++ {
		p.Next = &ListNode{l[i], nil}
		p = p.Next
	}
	return head.Next
}

func listEqual(l1, l2 *ListNode) bool {
	for {
		if l1 != nil && l2 != nil {
			if l1.Val != l2.Val {
				return false
			}
		} else {
			if l1 != l2 {
				return false
			} else {
				return true
			}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func TestAddTwoNumbers(t *testing.T) {
	qss := []question{
		question{
			q1:  makeList([]int{2, 4, 3}),
			q2:  makeList([]int{5, 6, 4}),
			ans: makeList([]int{7, 0, 8}),
		},
		question{
			q1:  makeList([]int{9, 9, 9}),
			q2:  makeList([]int{1, 1, 1}),
			ans: makeList([]int{0, 1, 1, 1}),
		},
		question{
			q1:  makeList([]int{0}),
			q2:  makeList([]int{1, 1, 1}),
			ans: makeList([]int{1, 1, 1}),
		},
		question{
			q1:  makeList([]int{0}),
			q2:  makeList([]int{0}),
			ans: makeList([]int{0}),
		},
	}

	for _, qs := range qss {
		t.Log(listEqual(addTwoNumbers(qs.q1, qs.q2), qs.ans))
	}
}
