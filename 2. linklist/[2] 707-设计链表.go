package linklist

// MyLinkedList 你可以选择使用单链表或者双链表，设计并实现自己的链表。
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
type MyLinkedList struct {
	Size      int
	DummyHead *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		DummyHead: &ListNode{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}
	p := this.DummyHead.Next
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this == nil {
		return
	}
	head := &ListNode{
		Val:  val,
		Next: this.DummyHead.Next,
	}
	this.DummyHead.Next = head
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	p := this.DummyHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{
		Val: val,
	}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this == nil || index < 0 || index > this.Size {
		return
	}
	p := this.DummyHead
	for i := 0; i < index; i++ {
		if p != nil {
			p = p.Next
		}
	}
	newNode := &ListNode{
		Val:  val,
		Next: p.Next,
	}
	p.Next = newNode
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this == nil || index < 0 || index >= this.Size {
		return
	}
	p := this.DummyHead
	for i := 0; i < index; i++ {
		if p != nil {
			p = p.Next
		}
	}
	p.Next = p.Next.Next
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
单链表的设计
1. 维护一个链表长度，可以规避访问时的越界问题，也可以更前置地判断访问的下标是否合法。
   如果合法，在遍历时就无需每次都额外判断节点是否为nil；
2. 注意题目， AddAtIndex在下标 == 链表长度时，需要追加到链表末尾，所以【index == this.Size】时也合法；
3. 遍历时都从虚拟头节点开始，可以规避链表为空时的越界问题。
    对于AddAtIndex和DeleteAtIndex来说，需要找到目标位置index的前一个节点进行操作，理应走index-1步，
	但由于是从DummyHead出发的，需要多走一步，即index步。
*/
