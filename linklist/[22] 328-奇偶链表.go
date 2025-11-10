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

/**
思路：
最开始的做法是另起一个链表，连接偶数位置的节点；原链表则保留奇数位置的节点。
遍历结束后，意味着奇偶链表分离完成，此时将进行头尾连接即可。
但是这种做法比较冗余。

实际上使用两个不同的头节点（奇数链表head & 偶数链表evenHead），使用两个不同的游标（奇数odd & 偶数even）进行串联即可。
最后进行拼接（遍历完成后，odd指向奇数链表的最后一个节点，直接指向evenHead即可）。
返回奇数链表头节点head。
*/
