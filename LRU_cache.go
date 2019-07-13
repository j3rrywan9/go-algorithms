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

func (c *LRUCache) addListNode(node *ListNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeListNode(node *ListNode) {
	var prevNode *ListNode = node.prev
	var nextNode *ListNode = node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (c *LRUCache) moveToHead(node *ListNode) {
	c.removeListNode(node)
	c.addListNode(node)
}

func (c *LRUCache) popTail() *ListNode {
	var res *ListNode = c.tail.prev
	c.removeListNode(res)
	return res
}

func (c *LRUCache) isKeyInCache(i int) bool {
	// A slice of all keys in the cache
	keys := make([]int, len(c.dict))
	for k, _ := range c.dict {
		keys = append(keys, k)
	}
	for _, x := range keys {
		if x == i {
			return true
		}
	}
	return false
}

func (c *LRUCache) get(key int) int {
	var node *ListNode
	if c.isKeyInCache(key) {
		node = c.dict[key]
		// Mark the record as recently used
		c.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (c *LRUCache) set(key, val int) {
	// Insert the value if the key is not present
	if c.isKeyInCache(key) == false {
		newNode := new(ListNode)
		newNode.key = key
		newNode.val = val
		c.dict[key] = newNode
		c.addListNode(newNode)

		// Invalidate the LRU record when the cache reached its capacity
		if len(c.dict) > c.capacity {
			var tail *ListNode = c.popTail()
			delete(c.dict, tail.key)
		}
	} else { // Set the value if the key is already present
		c.dict[key].val = val
		// Mark the record as recently used
		c.moveToHead(c.dict[key])
	}
}

func main() {
	cache := NewLRUCache(10)
	for i := 0; i < 12; i++ {
		cache.set(i, i+10)
	}
	cachedValue := cache.get(1)
	fmt.Printf("%d\n", cachedValue)
	cachedValue = cache.get(11)
	fmt.Printf("%d\n", cachedValue)
	cachedValue = cache.get(12)
	fmt.Printf("%d\n", cachedValue)
}
