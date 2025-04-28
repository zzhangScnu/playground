package linklist

// LRUCache 请你设计并实现一个满足
// LRU (最近最少使用) 缓存 约束的数据结构。
//
// 实现
// LRUCache 类：
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
// key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 示例：
//
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4
//
// 提示：
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
type LRUCache struct {
	cap      int
	data     MyCircularTwoWayLinkedList
	location map[int]*TwoWayListNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{
		cap:      capacity,
		data:     CircularTwoWayConstructor(),
		location: make(map[int]*TwoWayListNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.location[key]
	if !ok {
		return -1
	}
	val := node.Val
	this.data.RemoveNode(node)
	this.data.AddToTail(node)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.location[key]; ok {
		node.Val = value
		this.data.RemoveNode(node)
		this.data.AddToTail(node)
		return
	}
	if this.cap == this.data.Size {
		evictedNode := this.data.GetNode(0)
		this.data.DeleteAtIndex(0)
		delete(this.location, evictedNode.Key)
	}
	this.data.AddAtTail(key, value)
	this.location[key] = this.data.DummyHead.Pre
}

/**
诉求：
1. 缓存Put平均时间复杂度O(1) -> 需快速插入元素 -> 链表；
2. 缓存Get平均时间复杂度O(1) -> 需快速查找key对应的value -> 哈希表；
3. 缓存中的元素需要有时序，区分最近使用 / 久未使用，当容量满时需删除最久未使用的元素；
4. 每次访问元素，需将其变为最近使用，即需支持任意位置快速插入和删除元素 -> 链表。

用散列表+循环双向链表实现。
散列表：降低遍历链表找到对应节点的O(n)时间复杂度；
循环双向链表：
- 方便通过散列表找到节点后，快速定位到其前驱和后驱节点；
- 方便通过DummyHead实现快速插入队列末尾；
其实就是Java中LinkedHashMap的实现原理。

思路：
- 链表分为冷热区间，前面冷后面热。
最新访问的数据移动到热区间；
当插入数据时，若key已存在，则更新value且将节点移动到热区间；
当插入数据时，若key不存在且已满，则从冷区间驱逐第一个节点，再插入新节点。

map的使用：
当删除map中的一对键值对时，若m[key] = nil，不会真正删除，在下次访问if _, ok := this.location[key]; ok时，ok仍为true。
正确删除方式：delete(m, key)
*/
