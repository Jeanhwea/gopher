package main

import ("fmt")

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal2(root *TreeNode) []int {
	var traverse func(*TreeNode)
	ans := []int{}
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return ans
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return ans
}

func main() {
	root := &TreeNode{
		Val:  0,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: nil,
		},
	}
	ans := preorderTraversal2(root)
	fmt.Printf("ans = %v\n", ans)
}
