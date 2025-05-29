package tree

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
//
// 示例 1：
//
// 输入：root = [3,1,4,null,2], k = 1
// 输出：1
//
// 示例 2：
//
// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3
//
// 提示：
//
// 树中的节点数为 n 。
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

type SizedTreeNode struct {
	Val   int
	Left  *SizedTreeNode
	Right *SizedTreeNode
	Size  int // 子树节点总数
}

type BinarySearchTreeOptimizer struct {
	root *SizedTreeNode
}

func NewBinarySearchTreeOptimizer(val int) *BinarySearchTreeOptimizer {
	return &BinarySearchTreeOptimizer{
		root: &SizedTreeNode{
			Val:  val,
			Size: 1,
		},
	}
}

func BinarySearchTreeOptimizerConstructor(root *SizedTreeNode) {
	UpdateSize(root)
}

func (b *BinarySearchTreeOptimizer) findK(k int) int {
	if b.root == nil {
		return -1
	}
	return doFindK(b.root, k)
}

// todo：如果是第k大，则需要从右边开始算起，用右子树大小做比较
func doFindK(node *SizedTreeNode, k int) int {
	if node == nil {
		return -1
	}
	leftSize := getTreeSize(node)
	switch {
	case leftSize == k-1:
		return node.Val
	case leftSize >= k:
		return doFindK(node.Left, k)
	default:
		return doFindK(node.Right, k-leftSize-1)
	}
}

func getTreeSize(node *SizedTreeNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}

func (b *BinarySearchTreeOptimizer) Add(val int) {
	b.root = doAdd(b.root, val)
}

func doAdd(node *SizedTreeNode, val int) *SizedTreeNode {
	if node == nil {
		return &SizedTreeNode{Val: val, Size: 1}
	}
	if node.Val > val {
		node.Left = doAdd(node.Left, val)
	} else {
		node.Right = doAdd(node.Right, val)
	}
	node.Size = 1 + getTreeSize(node.Left) + getTreeSize(node.Right)
	return node
}

func (b *BinarySearchTreeOptimizer) Delete(val int) {
	b.root = doDelete(b.root, val)
}

func doDelete(node *SizedTreeNode, val int) *SizedTreeNode {
	if node == nil {
		return nil
	}
	if node.Val > val {
		node.Left = doDelete(node.Left, val)
	} else if node.Val < val {
		node.Right = doDelete(node.Right, val)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		cur := node.Right
		for cur != nil && cur.Left != nil {
			cur = cur.Left
		}
		node.Val = cur.Val
		node.Right = doDelete(node.Right, cur.Val)
	}
	node.Size = 1 + getTreeSize(node.Left) + getTreeSize(node.Right)
	return node
}

func UpdateSize(node *SizedTreeNode) int {
	if node == nil {
		return 0
	}
	node.Size = 1 + UpdateSize(node.Left) + UpdateSize(node.Right)
	return node.Size
}
