package main

import "fmt"

type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		queue: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	elem := this.queue[0]
	this.queue = this.queue[1:]
	return elem
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

func main() {
	/**
	 * Your MyQueue object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Push(x);
	 * param_2 := obj.Pop();
	 * param_3 := obj.Peek();
	 * param_4 := obj.Empty();
	 */

	obj := Constructor()
	obj.Push(1)
	param_2 := obj.Pop()
	param_5 := obj.Empty()

	fmt.Print(obj, param_2, param_5)
}
