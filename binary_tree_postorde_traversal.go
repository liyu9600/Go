package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	//return process(root)
	return process2(root)
}

func process(node *TreeNode) []int {

	if node == nil {
		return []int{}
	}

	res := make([]int, 0)

	l := process(node.Left)

	r := process(node.Right)

	res = append(res, l...)
	res = append(res, r...)
	res = append(res, node.Val)

	return res
}

func process2(node *TreeNode) []int {

	if node == nil {
		return []int{}
	}

	var pre *TreeNode
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

		if node.Right != nil && node.Right != pre {
			stack = append(stack, node)
			node = node.Right
		} else {
			res = append(res, node.Val)
			pre = node
			node = nil
		}
	}

	return res
}

func main() {

	d := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	/*d := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}*/

	fmt.Println(postorderTraversal(d))
}