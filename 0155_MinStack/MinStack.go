package main

import "fmt"

type MinStack struct {
	elements, min []int
	len           int
}

func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		min:      []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	if this.len == 0 {
		this.min = append(this.min, val)
	} else {
		m := this.GetMin()
		if val < m {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, m)
		}
	}
	this.len++
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:this.len-1]
	this.min = this.min[:this.len-1]
	this.len--
}

func (this *MinStack) Top() int {
	return this.elements[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.len-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())

	// minStack.Push(-3)
	// fmt.Println(minStack.GetMin())
}
