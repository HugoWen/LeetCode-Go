package main

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var queue MyQueue
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ret := 0
	l := len(this.stack1)
	for i := l - 1; i >= 0; i-- {
		ret = this.stack1[i]
		if i > 0 {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = this.stack1[:i]
	}
	l2 := len(this.stack2)
	if l2 == 0 {
		return ret
	}
	for i := l2 - 1; i >= 0; i-- {
		this.stack1 = append(this.stack1, this.stack2[i])
		this.stack2 = this.stack2[:i]
	}
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	ret := 0
	l := len(this.stack1)
	for i := l - 1; i >= 0; i-- {
		ret = this.stack1[i]
	}
	return ret
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
