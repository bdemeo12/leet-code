package main

import "fmt"

// There is a stream of n (idKey, value) pairs arriving in an arbitrary order,
//  where idKey is an integer between 1 and n and value is a string. No two pairs have the same id.

// Design a stream that returns the values in increasing order of their IDs
//  by returning a chunk (list) of values after each insertion.
//  The concatenation of all the chunks should result in a list of the sorted values.

// Implement the OrderedStream class:

// OrderedStream(int n) Constructs the stream to take n values.

// String[] insert(int idKey, String value) Inserts the pair (idKey, value) into the stream,
//  then returns the largest possible chunk of currently inserted values that appear next in the order.

type OrderedStream struct {
	stream  []string
	nextPtr int // points to the next expected key
}

func Constructor(n int) OrderedStream {
	stream := make([]string, n)
	return OrderedStream{
		stream:  stream,
		nextPtr: 0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {

	this.stream[idKey-1] = value

	// if it isnt the next expected key, return empty
	if idKey-1 != this.nextPtr {
		return []string{}
	}

	// otherwise collect largest chunk, starting at pointer
	largestChunk := []string{}
	for this.nextPtr < len(this.stream) && this.stream[this.nextPtr] != "" {
		largestChunk = append(largestChunk, this.stream[this.nextPtr])
		this.nextPtr++
	}

	return largestChunk
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

func main() {
	obj := Constructor(5)
	fmt.Println(obj.Insert(3, "ccccc"))
	fmt.Println(obj.Insert(1, "aaaaa"))
	fmt.Println(obj.Insert(2, "bbbbb"))
	fmt.Println(obj.Insert(5, "eeeee"))
	fmt.Println(obj.Insert(4, "ddddd"))
}
