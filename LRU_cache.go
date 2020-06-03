package main

import (
	"fmt"
)

// Doubly linked list node
type ListNode struct {
	key, val   int
	prev, next *ListNode
}

// LC 146
type LRUCache struct {
	cache      map[int]*ListNode
	head, tail *ListNode
	capacity   int
}

// Constructor
func NewLRUCache(capacity int) *LRUCache {
	lruCache := new(LRUCache)
	lruCache.capacity = capacity
	lruCache.cache = make(map[int]*ListNode, capacity)
	lruCache.head = new(ListNode)
	lruCache.tail = new(ListNode)
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache
}

func (l *LRUCache) addNode(node *ListNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *ListNode) {
	prevNode := node.prev
	nextNode := node.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

func (l *LRUCache) moveToHead(node *ListNode) {
	l.removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) popTail() *ListNode {
	res := l.tail.prev
	l.removeNode(res)
	return res
}

func (l *LRUCache) containsKey(key int) bool {
	if _, found := l.cache[key]; found {
		return true
	}
	return false
}

func (l *LRUCache) get(key int) int {
	var node *ListNode
	if l.containsKey(key) {
		node = l.cache[key]
		// Mark the record as recently used
		l.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (l *LRUCache) put(key, val int) {
	// Insert the value if the key is not present
	if l.containsKey(key) == false {
		newNode := new(ListNode)
		newNode.key = key
		newNode.val = val
		l.cache[key] = newNode
		l.addNode(newNode)

		// Invalidate the LRU record when the cache reached its capacity
		if len(l.cache) > l.capacity {
			tail := l.popTail()
			delete(l.cache, tail.key)
		}
	} else { // Set the value if the key is already present
		l.cache[key].val = val
		// Mark the record as recently used
		l.moveToHead(l.cache[key])
	}
}

func main() {
	lruCache := NewLRUCache(10)
	for i := 0; i < 12; i++ {
		lruCache.put(i, i+10)
	}
	cachedValue := lruCache.get(1)
	fmt.Printf("%d\n", cachedValue)
	cachedValue = lruCache.get(11)
	fmt.Printf("%d\n", cachedValue)
	cachedValue = lruCache.get(12)
	fmt.Printf("%d\n", cachedValue)
}
