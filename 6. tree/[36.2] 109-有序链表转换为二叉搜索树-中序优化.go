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

/**
结合【二叉搜索树中序遍历 -> 有序序列】的特性，可以一遍进行中序遍历，一遍构建二叉搜索树。
构造一个值为空的节点 -> 递归处理左子树 -> 回到根节点 -> 递归处理右子树，
每处理一个节点，链表的指针就向前移动一位，那么在处理树的中节点的时候，链表指针一定指向有序序列中该中节点的值。

注意：
对于Go来说，参数的传递均为值传递。
所以sortedListToBSTInorder(head *ListNode)，传入的head是链表指针head的地址的拷贝。在方法中修改head，实际上不会影响原本的head。
解决方案：
1. 全局变量：sortedListToBSTInorder
2. 传入指针地址的地址：doSortedListToBSTInorderByPointer(head **ListNode, start, end int)
*/

/**
结合【二叉搜索树中序遍历 -> 有序序列】的特性，可以一遍进行中序遍历，一遍构建二叉搜索树。
构造一个值为空的节点 -> 递归处理左子树 -> 回到根节点 -> 递归处理右子树，
每处理一个节点，链表的指针就向前移动一位，那么在处理树的中节点的时候，链表指针一定指向有序序列中该中节点的值。

注意：
对于Go来说，参数的传递均为值传递。
所以sortedListToBSTInorder(head *ListNode)，传入的head是链表指针head的地址的拷贝。在方法中修改head，实际上不会影响原本的head。
解决方案：
1. 全局变量：sortedListToBSTInorder
2. 传入指针地址的地址：doSortedListToBSTInorderByPointer(head **ListNode, start, end int)
*/
