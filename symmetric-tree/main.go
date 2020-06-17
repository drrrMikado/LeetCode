package main

import "fmt"

func main() {
	fmt.Println(isSymmetric(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
	//return isSymmetricNonRecusive(root)
}

func isMirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	} else if l.Val != r.Val {
		return false
	}
	return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}

func isSymmetricNonRecusive(root *TreeNode) bool {
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		left := q[0]
		right := q[1]
		q = q[2:]
		if left == nil && right == nil {
			continue
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		q = append([]*TreeNode{left.Left, right.Right, left.Right, right.Left}, q...)
	}
	return true
}
