package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func reverse(head *Node) *Node {

	var prevNode *Node
	curr := head

	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prevNode
		prevNode = curr
		curr = nextTmp
	}

	return prevNode
}

func createLinkedList(values []int) *Node {

	if len(values) == 0 {
		return nil
	}

	head := &Node{Val: values[0]}
	current := head
	for _, value := range values {
		current.Next = &Node{Val: value}
		current = current.Next
	}

	return head
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	printList(reverse(head))
}
