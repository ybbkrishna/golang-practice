package main

import "math"

// https://leetcode.com/problems/min-stack/
// Start Jun 18 2021 1:40pm
// End Jun 18 2021 1:56pm
type MinStack struct {
	Stack []int
	Min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		[]int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	currMin := math.MaxInt64
	if len(this.Min) > 0 {
		currMin = this.Min[len(this.Min)-1]
	}
	if val < currMin {
		currMin = val
	}
	this.Min = append(this.Min, currMin)
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]
	return
}

func (this *MinStack) Top() int {
	val := this.Stack[len(this.Stack)-1]
	return val
}

func (this *MinStack) GetMin() int {
	min := this.Min[len(this.Min)-1]
	return min
}
