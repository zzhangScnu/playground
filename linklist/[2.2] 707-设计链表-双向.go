package linklist

// MyTwoWayLinkedList 你可以选择使用单链表或者双链表，设计并实现自己的链表。
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
type MyTwoWayLinkedList struct {
	DummyHead *TwoWayListNode
	Size      int
}

func TwoWayConstructor() MyTwoWayLinkedList {
	return MyTwoWayLinkedList{
		DummyHead: &TwoWayListNode{},
	}
}

func (this *MyTwoWayLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	cur := this.DummyHead
	for i := 0; i < index+1; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyTwoWayLinkedList) AddAtHead(val int) {
	if this == nil {
		return
	}
	newNode := &TwoWayListNode{
		Val:  val,
		Next: this.DummyHead.Next,
		Pre:  this.DummyHead,
	}
	if this.DummyHead.Next != nil {
		this.DummyHead.Next.Pre = newNode
	}
	this.DummyHead.Next = newNode
	this.Size++
}

func (this *MyTwoWayLinkedList) AddAtTail(val int) {
	if this == nil {
		return
	}
	cur := this.DummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &TwoWayListNode{
		Val: val,
		Pre: cur,
	}
	this.Size++
}

func (this *MyTwoWayLinkedList) AddAtIndex(index int, val int) {
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
	if cur.Next != nil {
		cur.Next.Pre = newNode
	}
	cur.Next = newNode
	this.Size++
}

func (this *MyTwoWayLinkedList) DeleteAtIndex(index int) {
	if this == nil || index < 0 || index >= this.Size {
		return
	}
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil && cur.Next.Next != nil {
		cur.Next.Next.Pre = cur
	}
	cur.Next = cur.Next.Next
	this.Size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

/**
1. Get的时候要注意，从DummyHead和DummyHead.Next开始遍历，该走的步数是不一样的；
2. 相比起单向链表，双向在删除/新增节点的时候，更要注意越界问题：
	if cur.Next != nil {...}
	if cur.Next != nil && cur.Next.Next != nil {...}
3. 对于双向链表的删除/新增，需要先处理index处节点的后半截，再处理前半截，否则指向关系会丢失。
*/
