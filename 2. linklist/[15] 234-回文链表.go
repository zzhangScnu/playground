package linklist

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：head = [1,2,2,1]
// 输出：true
//
// 示例 2：
//
// 输入：head = [1,2]
// 输出：false
//
// 提示：
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	p, q := head, reverse(slow)
	for p != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
