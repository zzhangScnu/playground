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
func sortedListToBST(head *ListNode) *TreeNode {
	return doSortedListToBST(head, nil)
}

func doSortedListToBST(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}
	mid := findMiddle(head, tail)
	return &TreeNode{
		Val:   mid.Val,
		Left:  doSortedListToBST(head, mid),
		Right: doSortedListToBST(mid.Next, tail),
	}
}

func findMiddle(head, tail *ListNode) *ListNode {
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

/**
分治法。
因为二叉搜索树的中序遍历结果就是有序序列，所以可以：
每次先找到链路中点，作为根节点，再用[head, mid)和[mid+1, tail)递归构建左右子树。

所以
- 主方法中tail传入的是nil，表示链表结尾的后一个节点；
- 构建子树时，当head == tail即返回nil，因为对于左开右闭区间来说，此时区间内没有可用元素。
*/
