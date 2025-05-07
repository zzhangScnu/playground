package linklist

import "math"

// LFUCache 请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
// 实现 LFUCache 类：
//
// LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
// int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
// void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
// capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。
//
// 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
//
// 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
// 示例：
//
// 输入：
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
// "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
// 输出：
// [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
// 解释：
// // cnt(x) = 键 x 的使用计数
// // cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
// LFUCache lfu = new LFUCache(2);
// lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
// lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
// lfu.get(1);      // 返回 1
//
//	// cache=[1,2], cnt(2)=1, cnt(1)=2
//
// lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
//
//	// cache=[3,1], cnt(3)=1, cnt(1)=2
//
// lfu.get(2);      // 返回 -1（未找到）
// lfu.get(3);      // 返回 3
//
//	// cache=[3,1], cnt(3)=2, cnt(1)=2
//
// lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
//
//	// cache=[4,3], cnt(4)=1, cnt(3)=2
//
// lfu.get(1);      // 返回 -1（未找到）
// lfu.get(3);      // 返回 3
//
//	// cache=[3,4], cnt(4)=1, cnt(3)=3
//
// lfu.get(4);      // 返回 4
//
//	// cache=[3,4], cnt(4)=2, cnt(3)=3
//
// 提示：
//
// 1 <= capacity <= 10⁴
// 0 <= key <= 10⁵
// 0 <= value <= 10⁹
// 最多调用 2 * 10⁵ 次 get 和 put 方法
type LFUCache struct {
	KV           map[int]*TwoWayListNode
	KF           map[int]int
	FK           map[int]*MyCircularTwoWayLinkedList
	MinFrequency int
	Capacity     int
}

func LFUCacheConstructor(capacity int) LFUCache {
	return LFUCache{
		KV:           make(map[int]*TwoWayListNode),
		KF:           make(map[int]int),
		FK:           make(map[int]*MyCircularTwoWayLinkedList), // frequency -> linkedList
		MinFrequency: math.MaxInt,
		Capacity:     capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	node := this.KV[key]
	if node == nil {
		return -1
	}
	this.IncrementFrequency(key)
	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	if _, ok := this.KV[key]; ok {
		this.KV[key].Val = value
		this.IncrementFrequency(key)
		return
	}
	if len(this.KV) == this.Capacity {
		node := this.FK[this.MinFrequency].GetNode(0)
		delete(this.KV, node.Key)
		delete(this.KF, node.Key)
		if _, ok := this.FK[this.MinFrequency]; ok {
			this.FK[this.MinFrequency].RemoveNode(node)
		}
	}
	node := &TwoWayListNode{Key: key, Val: value}
	this.KV[key] = node
	currentFrequency := 1
	this.KF[key] = currentFrequency
	this.MinFrequency = currentFrequency
	this.AddNodeToFK(currentFrequency, node)
}

func (this *LFUCache) IncrementFrequency(key int) {
	currentFrequency := this.KF[key]
	updatedFrequency := currentFrequency + 1
	this.KF[key] = updatedFrequency
	node := this.KV[key]
	this.FK[currentFrequency].RemoveNode(node)
	if this.MinFrequency == currentFrequency && this.FK[currentFrequency].Size == 0 {
		this.MinFrequency = updatedFrequency
	}
	this.AddNodeToFK(updatedFrequency, node)
}

func (this *LFUCache) AddNodeToFK(frequency int, node *TwoWayListNode) {
	if _, ok := this.FK[frequency]; !ok {
		linkedList := CircularTwoWayConstructor()
		this.FK[frequency] = &linkedList
	}
	this.FK[frequency].AddToTail(node)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
思路：
在LRU缓存的基础上，增加访问频次的管理。
跟LRU相似，首先需要的数据结构：
1. KV：与LRU相同。key -> 链表节点node，便于通过key快速定位到链表的某个位置；
2. FK：与LRU相似。在LRU中，缓存中的数据使用双向链表进行维护；
		   在LFU中，因为有访问频次维度，将双向链表根据节点访问频次的不同，分割为不同的子链表，维护在哈希表中。
		   即frequency -> 双向链表doubleLinkedList，便于定位到某个node后，获取其前驱和后继节点，对该node进行删除。
3. KF：比LRU新增。key -> frequency，需要通过key获取访问频次，从而对其进行管理；
4. MinFrequency：比LRU新增。维护一个全局最小的访问频次，便于缓存满时快速定位到需处理的子链表，避免遍历寻找。

LFU要求：
1. O(1)时间复杂度通过key获取value；
2. 当缓存满，需要淘汰访问频次最小 + 访问时间最旧的数据。

实现：
1. Get方法：
- KV：O(1)地通过key获取链表节点node及其value，若不存在返回-1；
- KF：累加key对应的frequency；
- FK：根据KF获取oldFrequency，自增得到updatedFrequency，将node从oldFrequency子链表移动到updatedFrequency子链表尾部（注意不存在时需初始化）；
- MinFrequency：如果MinFrequency恰好等于oldFrequency，且对应的子链表长度为0，即该访问频次下已经没有节点了，则更新MinFrequency为updatedFrequency。

2. Put方法：
- 若key已存在，则更新对应node的value，并增加访问频次；
- 若key不存在，
	- 若缓存已满：通过MinFrequency + FK获取最小频次的子链表，获取头节点，并将其在KV、KF和FK中移除；
						  注意，从FK中移除时，因为是对子链表进行操作，需要前置判空。
						  避免MinFrequency == math.MaxInt及初始值时，此时子链表为空。
	- 新建node，插入KV、KF和FK，同时将MinFrequency置为1。
*/
