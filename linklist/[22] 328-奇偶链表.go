package linklist

// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别分组，保持它们原有的相对顺序，然后把偶数索引节点分组连接到奇数索引节点分组之后，
// 返回重新排序的链表。
//
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
//
// 示例 1:
//
// 输入: head = [1,2,3,4,5]
// 输出:[1,3,5,2,4]
//
// 示例 2:
//
// 输入: head = [2,1,3,5,6,4,7]
// 输出: [2,3,6,7,1,5,4]
//
// 提示:
//
// n == 链表中的节点数
// 0 <= n <= 10⁴
// -10⁶ <= Node.val <= 10⁶
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, even, evenHead := head, head.Next, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// todo：pre 的更新逻辑错误，本质是 混淆了「循环中处理的当前奇数节点」和「整个链表的最后一个奇数节点」。循环结束后，odd 才是真正的最后一个奇数节点（无论链表长度是奇数还是偶数），而 pre 只是循环中最后处理的那个奇数节点，并非最终的尾部。
//因此，修正的核心是让 pre 最终指向 odd（循环结束时的 odd），而不是在循环中跟随 odd 更新。
