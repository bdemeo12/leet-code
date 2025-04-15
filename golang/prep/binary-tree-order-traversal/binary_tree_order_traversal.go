package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	colMap := make(map[int][]int) // map of col, to val

	queue := []data{{
		node: root,
		col:  0,
	}}

	for len(queue) > 0 {
		// pop queue
		nodeData := queue[0]
		queue = queue[1:]

		// put it in a map
		if val, exists := colMap[nodeData.col]; !exists {
			colMap[nodeData.col] = []int{nodeData.node.Val}
		} else {
			val = append(val, nodeData.node.Val)
			colMap[nodeData.col] = val
		}

		if nodeData.node.Left != nil {
			queue = append(queue, data{
				node: nodeData.node.Left,
				col:  nodeData.col - 1,
			})
		}

		if nodeData.node.Right != nil {
			queue = append(queue, data{
				node: nodeData.node.Right,
				col:  nodeData.col + 1,
			})
		}
	}

	// sort hash table by key
	keys := []int{}
	for key, _ := range colMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(a, b int) bool {
		return keys[a] < keys[b]
	})

	order := [][]int{}
	for _, key := range keys {
		order = append(order, colMap[key])
	}

	return order
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
	arr := []interface{}{3, 9, 8, 4, 0, 1, 7}
	root := createBinaryTree(arr)
	fmt.Println(verticalOrder(root))
}
