package main

import "fmt"

// Given the root of a binary tree, return its maximum depth.

// A binary tree's maximum depth is the number of nodes along the longest path
//  from the root node down to the farthest leaf node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftcount := maxDepth(root.Left)
	rightcount := maxDepth(root.Right)

	return 1 + max(leftcount, rightcount) // +1 because we need to add the root to the count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func createBinaryTree(arr []interface{}) []*TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	tree := []*TreeNode{root}

	for i := 1; i < len(arr); i++ {
		if arr[i] == nil {
			tree = append(tree, nil)
			continue
		}

		node := &TreeNode{
			Val: arr[i].(int),
		}
		parent := tree[(i-1)/2]

		if parent != nil {

			if i%2 == 1 { //2*i + 1
				parent.Left = node
			} else { // 2*i + 2
				parent.Right = node
			}

			tree = append(tree, node)

		}

	}

	return tree
}

func main() {
	root := []interface{}{3, 9, 20, nil, nil, 15, 7} // does this need to be sorted ?
	tree := createBinaryTree(root)
	fmt.Println(maxDepth(tree[0]))
}
