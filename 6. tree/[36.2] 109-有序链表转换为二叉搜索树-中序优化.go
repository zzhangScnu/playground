package tree

// 给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为 平衡 二叉搜索树。
//
// 示例 1:
//
// 输入: head = [-10,-3,0,5,9]
// 输出: [0,-3,9,-10,null,5]
// 解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
//
// 示例 2:
//
// 输入: head = []
// 输出: []
//
// 提示:
//
// head 中的节点数在[0, 2 * 10⁴] 范围内
// -10⁵ <= Node.val <= 10⁵
func sortedListToBSTInorder(head *ListNode) *TreeNode {
	var curListNode *ListNode
	var traverse func(start, end int) *TreeNode
	traverse = func(start, end int) *TreeNode {
		if head == nil || start > end {
			return nil
		}
		mid := start + (end-start)/2
		node := &TreeNode{}
		node.Left = traverse(start, mid-1)
		node.Val = curListNode.Val
		curListNode = curListNode.Next
		node.Right = traverse(mid+1, end)
		return node
	}
	curListNode = head
	return traverse(0, getLen(head)-1)
}

func getLen(head *ListNode) int {
	var length int
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	return length
}

func sortedListToBSTInorderByPointer(head *ListNode) *TreeNode {
	return doSortedListToBSTInorderByPointer(&head, 0, getLen(head)-1)
}

func doSortedListToBSTInorderByPointer(head **ListNode, start, end int) *TreeNode {
	if head == nil || start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{}
	left := doSortedListToBSTInorderByPointer(head, start, mid-1)
	node.Val = (*head).Val
	*head = (*head).Next
	right := doSortedListToBSTInorderByPointer(head, mid+1, end)
	node.Left, node.Right = left, right
	return node
}
