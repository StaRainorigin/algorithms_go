package solution

import "math/rand"

type RandomizedSet struct {
	nums []int
	maps map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.maps[val]; ok {
		return false
	}
	this.maps[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.maps[val]
	if !ok {
		return false
	}
	n := len(this.nums)
	this.nums[index] = this.nums[n - 1]
	this.nums = this.nums[:n - 1]
	delete(this.maps, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
