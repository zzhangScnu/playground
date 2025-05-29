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

/**
若需对二叉搜索树进行频繁删除、插入和检索，可以通过在节点添加额外信息，记录以自身为根节点的树的大小，
从而快速计算出当前节点在整棵树中的位置，从而使用类二分搜索的思路，将定位时间减半。

- 增/删：与二叉搜索树相同，只是在递归操作中、每层结束后，需更新当前节点记录的树的大小；
- 查：通过K与当前节点为根节点的树的大小的比较，判断目标节点在左子树还是右子树，递归查找。

注意：
- 抽出辅助方法，用节点作为入参，供结构体的方法调用；
- 这样的实现是查找第K小的元素，即顺序查找，所以与左子树的大小进行比较；若需查找第K大的元素，则为逆序查找，与右子树的大小进行比较；
- 可以用switch-case+表达式代替if-else分支。
*/
