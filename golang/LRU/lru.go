package main

import (
	"fmt"
)

func main() {
	obj := Constructor(2)

	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}

// ______________________________________________
// With double linked list

type doubleLinkedList struct {
	prev *doubleLinkedList
	next *doubleLinkedList
	val  int
	key  int
}

type LRUCache struct {
	lNode *doubleLinkedList
	rNode *doubleLinkedList

	keyMap map[int]*doubleLinkedList

	capacity int
}

func Constructor(capacity int) LRUCache {
	keyMap := make(map[int]*doubleLinkedList, capacity)

	ln := &doubleLinkedList{
		val:  -1,
		prev: nil,
	}

	rn := &doubleLinkedList{
		val:  -1,
		next: nil,
	}

	ln.next = rn
	rn.prev = ln

	// 	nil -> ln -> rn -> nil

	return LRUCache{
		lNode:    ln,
		rNode:    rn,
		keyMap:   keyMap,
		capacity: capacity,
	}

}

func (this *LRUCache) Get(key int) int {

	if node, ok := this.keyMap[key]; ok {

		// move node in double linked list to end
		this.moveNodeToLastPlace(node)

		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if it exists in the map
	if node, exists := this.keyMap[key]; exists {
		// update value
		node.val = value
		this.keyMap[key] = node
		//move to last place
		this.moveNodeToLastPlace(node)
		return
	}

	// if its not present, check that we are not over capacity
	if len(this.keyMap) == this.capacity {
		// if were over capacity, delete first in list

		// delete from map
		delete(this.keyMap, this.lNode.next.key)

		// delete from double linked list
		this.deleteNode(this.lNode.next)
	}

	// add item to list place in the list
	newNode := &doubleLinkedList{
		val: value,
		key: key,
	}

	this.moveNodeToLastPlace(newNode)

	// add item to map
	this.keyMap[key] = newNode
}

func (this *LRUCache) deleteNode(node *doubleLinkedList) {

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
}

func (this *LRUCache) moveNodeToLastPlace(node *doubleLinkedList) {
	// nil -> left -> other node-> other node->  last place -> right -> nil

	this.rNode.next = node
	node.prev = this.rNode.prev
	node.next = this.rNode
	this.rNode.prev = node
}

// ___________________________________________________
// with map and counter :

// type LRUCache struct {
// 	cache        map[int]valueData
// 	currCapacity int
// 	capacity     int
// }

// type valueData struct {
// 	value        int
// 	lastAccessed time.Time
// }

// func Constructor(capacity int) LRUCache {
// 	cache := make(map[int]valueData)
// 	return LRUCache{
// 		cache:        cache,
// 		currCapacity: 0,
// 		capacity:     capacity,
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	valueData, ok := this.cache[key]
// 	if !ok {
// 		return -1
// 	}

// 	// update the last use timestamp
// 	valueData.lastAccessed = time.Now()
// 	// reassign to map
// 	this.cache[key] = valueData

// 	return valueData.value
// }

// func (this *LRUCache) Put(key int, value int) {

// 	// if a key already exists, update value
// 	if valueData, ok := this.cache[key]; ok {
// 		//update value
// 		valueData.value = value
// 		// update last accessed time
// 		valueData.lastAccessed = time.Now()

// 		// reassign
// 		this.cache[key] = valueData
// 		return
// 	}

// 	// if current capacity is at the total capacity, we need to delete the least recently used before we can add
// 	if this.currCapacity == this.capacity {

// 		// lets loop through out map and find the oldest time stamp

// 		first := true
// 		var lruKey int
// 		for key, value := range this.cache {
// 			lruVal := this.cache[lruKey]

// 			if first || value.lastAccessed.Before(lruVal.lastAccessed) {
// 				first = false
// 				lruKey = key
// 			}
// 		}

// 		// we have found the key for the least recently used item in our cache, lets delete it now
// 		delete(this.cache, lruKey)

// 		// decrement our capacity
// 		this.currCapacity--
// 	}

// 	// now we have space to add to our cache
// 	this.cache[key] = valueData{
// 		value:        value,
// 		lastAccessed: time.Now(),
// 	}

// 	this.currCapacity++
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
