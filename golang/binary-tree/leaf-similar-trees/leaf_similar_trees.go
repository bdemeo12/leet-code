package main

import (
	"fmt"
)

// Consider all the leaves of a binary tree, from left to right order,
// the values of those leaves form a leaf value sequence.

// Two binary trees are considered leaf-similar if their leaf value sequence is the same.

// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

// a leaf has no children
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	arr1 := addLeaf(root1)
	arr2 := addLeaf(root2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func addLeaf(n *TreeNode) []int {
	var leaves []int

	if n == nil {
		return leaves
	} else if n.Left == nil && n.Right == nil {
		leaves = append(leaves, n.Val)
	} else {
		leaves = append(leaves, addLeaf(n.Left)...)
		leaves = append(leaves, addLeaf(n.Right)...)
	}

	return leaves
}

func createBinaryTree(arr []interface{}) *TreeNode {
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

	return root
}

func main() {
	// root1 := createBinaryTree([]interface{}{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4})
	// root2 := createBinaryTree([]interface{}{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8})

	// root1 := createBinaryTree([]interface{}{1, 2, 3})
	// root2 := createBinaryTree([]interface{}{1, 3, 2})

	root1 := createBinaryTree([]interface{}{1, 2})
	root2 := createBinaryTree([]interface{}{2, 2})

	fmt.Println(leafSimilar(root1, root2))
}
