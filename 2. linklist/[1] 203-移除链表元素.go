package linklist

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
// 示例 1：
//
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
//
// 示例 2：
//
// 输入：head = [], val = 1
// 输出：[]
//
// 示例 3：
//
// 输入：head = [7,7,7,7], val = 7
// 输出：[]
//
// 提示：
//
// 列表中的节点数目在范围 [0, 10⁴] 内
// 1 <= Node.val <= 50
// 0 <= val <= 50
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	h, p := dummyHead, dummyHead
	for p != nil {
		if p.Next != nil && p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return h.Next
}

/**
1. 虚拟头节点：
- 可以规避空链表带来的访问越界问题；
- 可以规避处理头节点时的特殊逻辑。例如常规删除节点，将其前置节点指向其后置节点，即跳过该节点本身；而删除头节点时，直接将新的头节点指向后一个节点；
- 注意返回链表时，也需要返回虚拟头节点的下一个节点，即真正的链表本身；
2. 注意访问节点时，要先判空。
*/
