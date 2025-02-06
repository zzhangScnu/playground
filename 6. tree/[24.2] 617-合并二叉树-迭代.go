package tree

// 给你两棵二叉树： root1 和 root2 。
//
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠
// ，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
//
// 返回合并后的二叉树。
//
// 注意: 合并过程必须从两个树的根节点开始。
//
// 示例 1：
//
// 输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// 输出：[3,4,5,5,4,null,7]
//
// 示例 2：
//
// 输入：root1 = [1], root2 = [1,2]
// 输出：[2,2]
//
// 提示：
//
// 两棵树中的节点数目在范围 [0, 2000] 内
// -10⁴ <= Node.val <= 10⁴
func mergeTreesIteratively(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := NewQueue()
	queue.Push(root1)
	queue.Push(root2)
	for !queue.IsEmpty() {
		p, q := queue.Pop(), queue.Pop()
		p.Val += q.Val
		if p.Left != nil && q.Left != nil {
			queue.Push(p.Left)
			queue.Push(q.Left)
		}
		if p.Right != nil && q.Right != nil {
			queue.Push(p.Right)
			queue.Push(q.Right)
		}
		if p.Left == nil && q.Left != nil {
			p.Left = q.Left
		}
		if p.Right == nil && q.Right != nil {
			p.Right = q.Right
		}
	}
	return root1
}

/**
base case：
	if node1 == nil {
		return node2 // node2为nil也可以
	}
	if node2 == nil {
		return node1 // node1为nil也可以
	}
主要保证入列的都是非空节点。
每次从队列中弹出2个节点进行处理：
- 由于入队时保证节点非空，表示当前是2个需要叠加的节点，直接在第一棵树上累加值；
- 然后处理左右孩子：
  1. 情况一：两棵树的左/右孩子均不为空，则将它们并列推入队列，以便在下一层中弹出处理；
  2. 情况二：两棵树的左/右孩子其一为空：
    - 若树1的左孩子为空，树2的左孩子不为空，则直接将树1的左孩子指向树2的左孩子。即将左子树直接赋予过去。右孩子也一样；
	- 若树1的左孩子不为空，树2的左孩子为空，则不处理，树1的左孩子保持不变。右孩子也一样；
  3. 情况三：两棵树的左/右孩子均为空，即当前为叶子节点，不处理。

需要注意的是，需要先将都非空的左右孩子入队，再处理直接赋予的情况。
否则会变成，举个例子，原本没有右孩子的树1，在继承了树2的右孩子后，因为树1和树2都有了右孩子，导致入队了，在下一层错误累加计算。
*/
