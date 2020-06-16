package main

import "fmt"

func main() {
	var l *ListNode
	l = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	printListNode(deleteDuplicates(l))
	l = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}}
	printListNode(deleteDuplicates(l))
}

func printListNode(l *ListNode) {
	t := l
	for t.Next != nil {
		fmt.Printf("%d->", t.Val)
		t = t.Next
	}
	fmt.Printf("%d\n", t.Val)
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return head
}
