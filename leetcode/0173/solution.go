package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.inorder(root)
	return it
}

func (it *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}

	it.inorder(node.Left)
	it.nums = append(it.nums, node.Val)
	it.inorder(node.Right)
}

func (this *BSTIterator) Next() int {
	val := this.nums[0]
	this.nums = this.nums[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
