package linklist

// MyCircularTwoWayLinkedList 你可以选择使用单链表或者双链表，设计并实现自己的链表。
//
// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
//
// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
//
// 实现 MyLinkedList 类：
//
// MyLinkedList() 初始化 MyLinkedList 对象。
// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
// void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果
// index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
//
// 示例：
//
// 输入
// ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get",
// "deleteAtIndex", "get"]
// [[], [1], [3], [1, 2], [1], [1], [1]]
// 输出
// [null, null, null, null, 2, null, 3]
//
// 解释
// MyLinkedList myLinkedList = new MyLinkedList();
// myLinkedList.addAtHead(1);
// myLinkedList.addAtTail(3);
// myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
// myLinkedList.get(1);              // 返回 2
// myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
// myLinkedList.get(1);              // 返回 3
//
// 提示：
//
// 0 <= index, val <= 1000
// 请不要使用内置的 LinkedList 库。
// 调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。
type MyCircularTwoWayLinkedList struct {
	DummyHead *TwoWayListNode
	Size      int
}

func CircularTwoWayConstructor() MyCircularTwoWayLinkedList {
	dummyHead := &TwoWayListNode{}
	dummyHead.Pre = dummyHead
	dummyHead.Next = dummyHead
	return MyCircularTwoWayLinkedList{
		DummyHead: dummyHead,
	}
}

func (this *MyCircularTwoWayLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	cur := this.DummyHead
	for i := 0; i < index+1; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyCircularTwoWayLinkedList) AddAtHead(val int) {
	if this == nil {
		return
	}
	newNode := &TwoWayListNode{
		Val:  val,
		Next: this.DummyHead.Next,
		Pre:  this.DummyHead,
	}
	this.DummyHead.Next.Pre = newNode
	this.DummyHead.Next = newNode
	this.Size++
}

func (this *MyCircularTwoWayLinkedList) AddAtTail(val int) {
	if this == nil {
		return
	}
	newNode := &TwoWayListNode{
		Val:  val,
		Pre:  this.DummyHead.Pre,
		Next: this.DummyHead,
	}
	this.DummyHead.Pre.Next = newNode
	this.DummyHead.Pre = newNode
	this.Size++
}

func (this *MyCircularTwoWayLinkedList) AddAtIndex(index int, val int) {
	if this == nil || index < 0 || index > this.Size {
		return
	}
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode := &TwoWayListNode{
		Val:  val,
		Next: cur.Next,
		Pre:  cur,
	}
	cur.Next.Pre = newNode
	cur.Next = newNode
	this.Size++
}

func (this *MyCircularTwoWayLinkedList) DeleteAtIndex(index int) {
	if this == nil || index < 0 || index >= this.Size {
		return
	}
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next.Next.Pre = cur
	cur.Next = cur.Next.Next
	this.Size--
}

/**
函数实现基本跟双向链表一致，不同点在于：
1. 初始化的时候，DummyHead的Pre和Next都指向自身，形成一个环；
2. 在查找、操作时，无需再判断指针是否为空，因为在环状链表中，永远不会为空。
*/
