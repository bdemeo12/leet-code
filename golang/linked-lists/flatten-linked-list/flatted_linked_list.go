package main

import "fmt"

//
// linked list:
//
// Imagine a linked list with not just a next node, but a down node as well.
//
// Write a function that would "Flatten" that linked list by taking all the downward
// segments and merging them up between their parent and their parent's next.
//
// input:
//
//     [1] -> [2] -> [3] -> [8] -> [10]
//                    |      |
//                    |     [9]
//                    |
//                   [4] -> [5] -> [6]
//                                  |
//                                 [7]
//
// output:
//
//     [1] -> [2] -> [3] -> [4] -> [5] -> [6] -> [7] -> [8] -> [9] -> [10]
// 1,2 ,3 -> 4,5 , 6,7, -> .8, 10

type ListNode struct {
	Val          int
	Next         *ListNode
	DownListNode *ListNode
}

func flattenList(head *ListNode) *ListNode { // tail

	if head == nil {
		return nil
	}

	curr := head
	var lastGoodVal *ListNode // keeps track of last value not equal to nil
	for curr != nil {
		if curr.DownListNode != nil {
			tail := flattenList(curr.DownListNode) // head of sublist -> 4, 5, 6, 7

			// reassign pointers
			tail.Next = curr.Next // keeps next point in mem // curr = 3 curr.next = 8
			curr.Next = curr.DownListNode
			curr.DownListNode = nil
		}

		lastGoodVal = curr
		curr = curr.Next
	}

	return lastGoodVal
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("List is empty")
	}

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {}
