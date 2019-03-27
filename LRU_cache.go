package main

import (
	"fmt"
)

// Doubly linked list
type ListNode struct {
	key, val   int
	prev, next *ListNode
}

// LCOJ No. 146
type LRUCache struct {
	capacity   int
	dict       map[int]*ListNode
	head, tail *ListNode
}

// Constructor
func NewLRUCache(capacity int) *LRUCache {
	cache := new(LRUCache)
	cache.capacity = capacity
	cache.dict = make(map[int]*ListNode, capacity)
	cache.head = new(ListNode)
	cache.tail = new(ListNode)
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) addListNode(node *ListNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeListNode(node *ListNode) {
	var prevNode *ListNode = node.prev
	var nextNode *ListNode = node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (this *LRUCache) moveToHead(node *ListNode) {
	this.removeListNode(node)
	this.addListNode(node)
}

func (this *LRUCache) popTail() *ListNode {
	var res *ListNode = this.tail.prev
	this.removeListNode(res)
	return res
}

func (this *LRUCache) isKeyInCache(i int) bool {
	// A slice of all keys in the cache
	keys := make([]int, len(this.dict))
	for k, _ := range this.dict {
		keys = append(keys, k)
	}
	for _, x := range keys {
		if x == i {
			return true
		}
	}
	return false
}

func (this *LRUCache) get(key int) int {
	var node *ListNode
	if this.isKeyInCache(key) {
		node = this.dict[key]
		// Mark the record as recently used
		this.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) set(key, val int) {
	// Insert the value if the key is not present
	if this.isKeyInCache(key) == false {
		newNode := new(ListNode)
		newNode.key = key
		newNode.val = val
		this.dict[key] = newNode
		this.addListNode(newNode)

		// Invalidate the LRU record when the cache reached its capacity
		if len(this.dict) > this.capacity {
			var tail *ListNode = this.popTail()
			delete(this.dict, tail.key)
		}
	} else { // Set the value if the key is already present
		this.dict[key].val = val
		// Mark the record as recently used
		this.moveToHead(this.dict[key])
	}
}

func main() {
	cache := NewLRUCache(10)
	for i := 0; i < 12; i++ {
		cache.set(i, i+10)
	}
	cached_val := cache.get(1)
	fmt.Printf("%d\n", cached_val)
	cached_val = cache.get(11)
	fmt.Printf("%d\n", cached_val)
	cached_val = cache.get(12)
	fmt.Printf("%d\n", cached_val)
}
