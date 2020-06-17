package main

import "fmt"

func main() {
	fmt.Println(maxDepth(nil))
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		ml := maxDepth(root.Left)
		mr := maxDepth(root.Right)
		if ml < mr {
			return mr + 1
		} else {
			return ml + 1
		}
	}
}

func maxDepthNonRecrusive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type S struct {
		K *TreeNode
		V int
	}
	depth := 0
	q := []S{{root, 1}}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		root = c.K
		curDepth := c.V
		if root != nil {
			if depth < curDepth {
				depth = curDepth
			}
			q = append([]S{{root.Left, curDepth + 1}}, q...)
			q = append([]S{{root.Right, curDepth + 1}}, q...)
		}
	}
	return depth
}
