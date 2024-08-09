package main

package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
    out := make([]int, 0, 16)
	return postorderTraversalAppend(root, out)
}

func postorderTraversalAppend(root *TreeNode, out []int) []int {
	if root.Left != nil {
		out = postorderTraversalAppend(root.Left, out)
	}
	if root.Right != nil {
		out = postorderTraversalAppend(root.Right, out)
	}
	
	out = append(out, root.Val)

	
	return out
}