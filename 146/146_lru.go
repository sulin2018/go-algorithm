package main

type LRUCache struct {
	size       int
	capacity   int
	data       map[int]*DataNode
	head, tail *DataNode
}

type DataNode struct {
	key, value int
	next, pre  *DataNode
}

func initNode(key, value int) *DataNode {
	return &DataNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
		data:     map[int]*DataNode{},
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.data[key]; !ok {
		return -1
	}
	node := this.data[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		// 存在就更新 并移动到最前面
		node := this.data[key]
		node.value = value
		this.moveToHead(node)
	} else {
		tempNode := initNode(key, value)
		this.data[key] = tempNode
		this.addHead(tempNode)
		this.size++
		// 否则 判断size是否与capacity相等, 相等移除末尾再添加
		if this.size > this.capacity {
			removeNode := this.removeTail()
			delete(this.data, removeNode.key)
			this.size--
		}
	}
}

func (s *LRUCache) removeNode(tempNode *DataNode) {
	tempNode.pre.next = tempNode.next
	tempNode.next.pre = tempNode.pre
}

func (s *LRUCache) addHead(tempNode *DataNode) {
	tempNode.pre = s.head
	tempNode.next = s.head.next
	s.head.next.pre = tempNode
	s.head.next = tempNode
}

func (s *LRUCache) removeTail() *DataNode {
	node := s.tail.pre
	s.removeNode(node)
	return node
}

func (s *LRUCache) moveToHead(tempNode *DataNode) {
	s.removeNode(tempNode)
	s.addHead(tempNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
