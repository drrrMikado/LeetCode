package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	t := ans
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			t.Next = l2
			l2 = l2.Next
		} else {
			t.Next = l1
			l1 = l1.Next
		}
		t = t.Next
	}

	if l1 != nil {
		t.Next = l1
	}
	if l2 != nil {
		t.Next = l2
	}

	return ans.Next
}
