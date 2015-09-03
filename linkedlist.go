package main

import (
	"fmt"
)

type ListNode struct {
	val int
	next *ListNode
}

func (this *ListNode) CreateLinkedList(length int) {
	current := this
	for i := 2; i <= length; i++ {
		pListNode := &ListNode{i, nil}
		current.next = pListNode
		current = pListNode
	}
}

func (this *ListNode) PrintLinkedList() {
	current := this
	for {
		fmt.Printf("%d", current.val)
		if current.next != nil {
			fmt.Print("->")
			current = current.next
		} else {
			break
		}
	}
	fmt.Println()
}

// LCOJ No. 206
func (this *ListNode) ReverseLinkedList() *ListNode {
	var head *ListNode = nil
	for this != nil {
		temp := this
		this = this.next
		temp.next = head
		head = temp
	}
	return head
}

// LCOJ No. 237
func (this *ListNode) DeleteNode(pListNode *ListNode) {
	if pListNode != nil && pListNode.next != nil {
		pListNode.val = pListNode.next.val
		pListNode.next = pListNode.next.next
	}
}

// LCOJ No. 203
func (this *ListNode) RemoveElements(val int) {
	if this != nil {
		current := this
		for current.next != nil {
			if current.val == val {
				current.val = current.next.val
				current.next = current.next.next
			}
			current = current.next
		}
	}
}

// LCOJ No. 83
func (this *ListNode) DeleteDuplicates() {
	if this == nil {
		return
	} else {
		current := this
		for current.next != nil {
			if current.val == current.next.val {
				current.next = current.next.next
			} else {
				current = current.next
			}
		}
	}
}

// LCOJ No. 19
func (this *ListNode) RemoveNthFromEnd(n int) {
	fast, slow := this, this
	for i := 0; i < n; i++ {
		fast = fast.next
	}
	// Nth node from end equals head node
	if fast == nil {
		this = this.next
	}
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	slow.next = slow.next.next
}

func main() {
	mylistnode := &ListNode{1, nil}
	fmt.Println("Creating a 10-node linked list...")
	mylistnode.CreateLinkedList(10)
	mylistnode.PrintLinkedList()
	fmt.Println("Deleting 3rd node from above linked list...")
	mylistnode.DeleteNode(mylistnode.next.next)
	mylistnode.PrintLinkedList()
	fmt.Println("Removing the node with value equals 5 from above linked list...")
	mylistnode.RemoveElements(5)
	mylistnode.PrintLinkedList()
	fmt.Println("Adding a duplicate node to above linked list...")
	newlistnode := &ListNode{1, nil}
	newlistnode.next = mylistnode
	mylistnode = newlistnode
	mylistnode.PrintLinkedList()
	fmt.Println("Deleting duplicate nodes from above linked list...")
	mylistnode.DeleteDuplicates()
	mylistnode.PrintLinkedList()
	fmt.Println("Removing 4th node from end of above linked list...")
	mylistnode.RemoveNthFromEnd(4)
	mylistnode.PrintLinkedList()
	fmt.Println("Reversing above linked list...")
	mylistnode2 := mylistnode.ReverseLinkedList()
	mylistnode2.PrintLinkedList()
}

