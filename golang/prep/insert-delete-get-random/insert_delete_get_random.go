package main

import (
	"fmt"
	"math/rand"
)

// Implement the RandomizedSet class:

// RandomizedSet() Initializes the RandomizedSet object.

// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.

// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.

// int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

// You must implement the functions of the class such that each function works in average O(1) time complexity.

type RandomizedSet struct {
	Nums []int
	Set  map[int]int
}

func Constructor() RandomizedSet {
	set := make(map[int]int)
	return RandomizedSet{
		Nums: []int{},
		Set:  set,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.Set[val]; exists {
		return false
	}

	// if it doesnt exist, insert it
	this.Nums = append(this.Nums, val)
	this.Set[val] = len(this.Nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, exists := this.Set[val]; exists {

		// remove from nums array
		last := this.Nums[len(this.Nums)-1]
		this.Nums[index] = last
		this.Set[last] = index

		this.Nums = this.Nums[:len(this.Nums)-1]

		// remove from set map
		delete(this.Set, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.Nums))
	return this.Nums[randomIndex]
}

func main() {

	obj := Constructor()

	obj.Insert(0)
	obj.Insert(1)
	obj.Remove(0)
	obj.Insert(2)
	obj.Remove(1)
	fmt.Println(obj.GetRandom())
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
