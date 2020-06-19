package main

import "fmt"

func main() {
	var tree *TreeNode
	//tree = nil
	//fmt.Println(levelOrderBottom(tree))
	tree = &TreeNode{
		Val:  3,
		Left: &TreeNode{9, nil, nil},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{15, nil, nil},
			Right: &TreeNode{7, nil, nil},
		},
	}
	fmt.Println(levelOrderBottom(tree))
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	dfs(root, 0, &ans)
	l := len(ans)
	end := l / 2
	for i := 0; i < end; i++ {
		ans[i], ans[l-i-1] = ans[l-i-1], ans[i]
	}

	return ans
}

func dfs(n *TreeNode, level int, ans *[][]int) {
	if n == nil {
		return
	}
	if len(*ans) == level {
		*ans = append(*ans, []int{})
	}
	(*ans)[level] = append((*ans)[level], n.Val)
	dfs(n.Left, level+1, ans)
	dfs(n.Right, level+1, ans)
}
