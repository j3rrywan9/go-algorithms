package main

import (
	"fmt"
)

type Queue struct {
	stack1 []int
	stack2 []int
	size   int
}

func (this *Queue) Push(x int) {
	this.stack1 = append(this.stack1, x)
	this.size += 1
}

func (this *Queue) Pop() {
	if this.stack2 != nil {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	} else {
		for i := 0; i < this.size-1; i++ {
			x := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, x)
		}
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
	this.size -= 1
}

func (this *Queue) Peek() int {
	if this.stack2 != nil {
		return this.stack2[len(this.stack2)-1]
	} else {
		for i := 0; i < this.size; i++ {
			x := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, x)
		}
		return this.stack2[len(this.stack2)-1]
	}
}

func (this *Queue) Empty() bool {
	return this.size == 0
}

func (this *Queue) Print() {
	fmt.Println("stack2:", this.stack2)
	fmt.Println("stack1:", this.stack1)
	fmt.Println("queue:", append(this.stack2, this.stack1...))
}

func main() {
	myqueue := new(Queue)
	myqueue.Push(1)
	myqueue.Push(2)
	myqueue.Push(3)
	myqueue.Push(4)
	myqueue.Push(5)
	myqueue.Pop()
	test := myqueue.Peek()
	fmt.Printf("%d\n", test)
	myqueue.Pop()
	myqueue.Print()
	fmt.Printf("Is queue empty? %t\n", myqueue.Empty())
}
