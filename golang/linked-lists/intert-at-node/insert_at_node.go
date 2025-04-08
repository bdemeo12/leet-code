package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtPosition(list *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {

	if position == 0 {
		return &SinglyLinkedListNode{data: data, next: list}
	}

	// loop through next position times -1
	// at position, next should point to new node data, and data should point to old next

	current := list
	for i := 0; i < int(position)-1; i++ {
		if current.next == nil {
			return list // position is out of bounds
		}

		current = current.next // this will give us the poisiton we want to update
	}

	newNode := &SinglyLinkedListNode{data: data, next: current.next}
	current.next = newNode

	return list
}

func arrToList(arr []int32) *SinglyLinkedListNode {

	head := &SinglyLinkedListNode{data: int32(arr[0])}
	current := head
	for i := 1; i < len(arr); i++ {
		current.next = &SinglyLinkedListNode{data: arr[i]}
		current = current.next
	}

	return head
}

func main() {

	arr := []int32{13, 13, 7}
	list := arrToList(arr)

	var data int32 = 1

	var position int32 = 2

	fmt.Println(insertNodeAtPosition(list, data, position))
}
