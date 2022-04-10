package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	//return process(root)
	return process2(root)
}

func process(node *TreeNode) []int {

	if node == nil {
		return []int{}
	}

	res := []int{node.Val}

	l := process(node.Left)

	r := process(node.Right)

	res = append(res, l...)
	res = append(res, r...)

	return res
}

func process2(node *TreeNode) []int {

	if node == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for node != nil || len(stack) > 0 {

		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		l := len(stack)
		node = stack[l-1]
		stack = stack[0 : l-1]

		node = node.Right

		/*if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}*/
	}

	return res
}