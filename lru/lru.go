package main

/*
方法1：有序字典
方法2：哈希表+双向链表
本题是练习面向对象语法的好题，用于我学习golang语法，方法并不难，双链表，加上 hash表，本题我加了两个虚拟节点Head,Tail方便操作
*/
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	limit int
	hash  map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	h := &Node{-1, -1, nil, nil}
	t := &Node{-1, -1, nil, nil}
	h.Next = t
	t.Prev = h
	hash := make(map[int]*Node, capacity)
	cache := LRUCache{hash: hash, limit: capacity, Head: h, Tail: t}
	return cache
}

func (this *LRUCache) insert(node *Node) {
	t := this.Tail
	node.Prev = t.Prev
	t.Prev.Next = node
	node.Next = t
	t.Prev = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		v.Val = value
	} else {
		if len(this.hash) >= this.limit {
			h := this.Head.Next
			this.remove(h)
			delete(this.hash, h.Key)
		}
		node := &Node{key, value, nil, nil}
		this.hash[key] = node
		this.insert(node)
	}
}
