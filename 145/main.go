package main

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func postorderTraversal(root *treeNode) []int {
	if root == nil {
		return nil
	}

	out := make([]int, 0, 16)
	return postorderTraversalAppend(root, out)
}

func postorderTraversalAppend(root *treeNode, out []int) []int {
	if root.Left != nil {
		out = postorderTraversalAppend(root.Left, out)
	}
	if root.Right != nil {
		out = postorderTraversalAppend(root.Right, out)
	}

	out = append(out, root.Val)

	return out
}
