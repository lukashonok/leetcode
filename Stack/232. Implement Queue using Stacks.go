package main

type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor3() MyQueue {
	return MyQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.s1) != 0 {
		top := this.s1[len(this.s1)-1]
		this.s2 = append(this.s2, top)
		this.s1 = this.s1[:len(this.s1)-1]
	}
	this.s1 = append(this.s1, x)
	for len(this.s2) != 0 {
		top := this.s2[len(this.s2)-1]
		this.s1 = append(this.s1, top)
		this.s2 = this.s2[:len(this.s2)-1]
	}
}

func (this *MyQueue) Pop() int {
	top := this.s1[len(this.s1)-1]
	this.s1 = this.s1[:len(this.s1)-1]
	return top
}

func (this *MyQueue) Peek() int {
	return this.s1[len(this.s1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}
