package linklist

import "container/list"

type EmbeddedLRUCache struct {
	capacity int
	values   *list.List
	mapping  map[int]*list.Element
}

func EmbeddedLRUCacheConstructor(capacity int) EmbeddedLRUCache {
	return EmbeddedLRUCache{
		capacity: capacity,
		values:   list.New(),
		mapping:  make(map[int]*list.Element),
	}
}

func (this *EmbeddedLRUCache) Get(key int) int {
	if node, ok := this.mapping[key]; ok {
		this.values.MoveToFront(node)
		return node.Value.([]int)[1]
	}
	return -1
}

func (this *EmbeddedLRUCache) Put(key int, value int) {
	if node, ok := this.mapping[key]; ok {
		node.Value = []int{key, value}
		this.values.MoveToFront(node)
		return
	}
	if this.values.Len() == this.capacity {
		node := this.values.Back()
		this.values.Remove(node)
		delete(this.mapping, node.Value.([]int)[0])
	}
	node := this.values.PushFront([]int{key, value})
	this.mapping[key] = node
}

/**
List represents a doubly linked list.

注意，node 中必须存储完整键值对，否则逐出时无法关联至 map 中应逐出位置。
*/
