package main

import (
	"fmt"
	"math/rand"
)

// use hashmap to keep track of index in array
// when removing an item from the array, switch with the last element
// this way, no gaps
// updating last index will always take constant time.

type RandomizedSet struct {
	Set    map[int]int
	Lookup []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int),
		make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.Set[val]
	if ok {
		return false
	}
	index := len(this.Lookup)
	this.Lookup = append(this.Lookup, val)

	this.Set[val] = index
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.Set[val]
	if !ok {
		return false
	}

	last_index := len(this.Lookup) - 1

	if index != last_index {
		this.Lookup[index] = this.Lookup[last_index]
		this.Set[this.Lookup[index]] = index
	}
	this.Lookup = this.Lookup[:last_index]
	delete(this.Set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	target_index := rand.Intn(len(this.Lookup))
	return this.Lookup[target_index]
}

func main() {
	set := Constructor()
	fmt.Println(set.Insert(1))
	fmt.Println(set.Insert(2))

	set.Remove(2)

	// fmt.Println(set.Remove(1))
	// fmt.Println(set.Remove(1))

	fmt.Println(set)
}
