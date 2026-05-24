package solution

import "math"

type num struct {
	data int
	preMin int
}

type MinStack []num

func Constructor() MinStack {
	return MinStack{{0, math.MaxInt}}
}

func (this *MinStack) Push(val int) {
	*this = append(*this, num{val, min(val, this.GetMin())})
}

func (this *MinStack) Pop() {
	*this = (*this)[:len(*this)-1]
}

func (this *MinStack) Top() int {
	return (*this)[len(*this)-1].data
}

func (this *MinStack) GetMin() int {
	return (*this)[len(*this)-1].preMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
