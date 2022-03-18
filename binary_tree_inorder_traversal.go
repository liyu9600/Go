package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return process(root)
	return process2(root)
}

func process(node *TreeNode) []int {

	if node == nil {
		return []int{}
	}

	l := process(node.Left)

	l = append(l, node.Val)

	r := process(node.Right)

	return append(l, r...)
}

func process2(node *TreeNode) []int {

	if node == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for node != nil || len(stack) > 0 {

		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		l := len(stack)
		node = stack[l-1]
		stack = stack[0 : l-1]
		res = append(res, node.Val)

		node = node.Right

		/*if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}*/
	}

	return res
}
