package main

type MyStack struct {
	q []int
}

func Constructor5() MyStack {
	return MyStack{q: []int{}}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	for i := 0; i < len(this.q)-1; i++ {
		top := this.q[0]
		this.q = this.q[1:]
		this.q = append(this.q, top)
	}
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.q = this.q[1:]
	return top
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

// func main() {
// 	obj := Constructor()
// 	obj.Push(1)
// 	obj.Push(2)
// 	obj.Push(3)
// 	fmt.Println(obj.Pop())
// 	fmt.Println(obj.Top())
// 	fmt.Println(obj.Empty())
// }
