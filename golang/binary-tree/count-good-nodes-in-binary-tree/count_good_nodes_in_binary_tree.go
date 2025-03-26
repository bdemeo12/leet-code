package main

import "fmt"

// Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

// Return the number of good nodes in the binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {

	if root == nil {
		return 1
	}

	val := root.Val

	return 1 + dfs(root.Left, val) + dfs(root.Right, val)
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}

	count := 0
	if root.Val >= val {
		count = 1
		val = root.Val
	}

	count += dfs(root.Left, val)
	count += dfs(root.Right, val)

	return count
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

		parent := tree[(i-1)/2]

		node := &TreeNode{Val: arr[i].(int)}

		if parent != nil {
			if i%2 == 1 { //left
				parent.Left = node
			} else { //right
				parent.Right = node
			}

			tree = append(tree, node)
		}
	}

	return tree
}

func main() {

	root := []interface{}{3, 1, 4, 3, nil, 1, 5}
	tree := createBinaryTree(root)
	fmt.Println(goodNodes(tree[0]))
}
