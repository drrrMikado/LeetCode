package main

import "fmt"

func main() {
	p := &TreeNode{Val: 1, Left: &TreeNode{2, nil, nil}, Right: &TreeNode{3, nil, nil}}
	q := &TreeNode{Val: 1, Left: &TreeNode{2, nil, nil}, Right: &TreeNode{3, nil, nil}}
	fmt.Println(isSameTree(p, q))
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}
