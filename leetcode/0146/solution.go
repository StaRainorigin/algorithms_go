package solution

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	cap int
	head *Node
	mp map[int]*Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.prev = dummy
	dummy.next = dummy
	return LRUCache{
		cap:  capacity,
		head:     dummy,
		mp: map[int]*Node{},
	}
}

func remove(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) insertHead(node *Node) {
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.mp[key]; ok {
        remove(node)
        this.insertHead(node)
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.mp[key]; ok {
        node.value = value
        remove(node)
        this.insertHead(node)
        return
    }
    if len(this.mp) >= this.cap {
        delete(this.mp, this.head.prev.key)
        remove(this.head.prev)
    }
    node := &Node{key: key, value: value}
    this.insertHead(node)
    this.mp[key] = node
}


/*
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
