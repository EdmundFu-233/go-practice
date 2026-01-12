package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: val}
	}
	if val < root.Value {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}
	return root
}

func Inorder(root *TreeNode) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Printf("%d ", root.Value)
	Inorder(root.Right)
}

func main() {
	var root *TreeNode
	vals := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range vals {
		root = Insert(root, v)
	}
	Inorder(root)
	fmt.Println()
}
