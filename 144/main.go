package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
    out := make([]int, 0, 16)
	return preorderTraversalAppend(root, out)
}

func preorderTraversalAppend(root *TreeNode, out []int) []int {
	out = append(out, root.Val)
	
	if root.Left != nil {
		out = preorderTraversalAppend(root.Left, out)
	}
	if root.Right != nil {
		out = preorderTraversalAppend(root.Right, out)
	}
	
	return out
}