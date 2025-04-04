package tree

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//
// 每条从根节点到叶节点的路径都代表一个数字：
//
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//
// 计算从根节点到叶节点生成的 所有数字之和 。
//
// 叶节点 是指没有子节点的节点。
//
// 示例 1：
//
// 输入：root = [1,2,3]
// 输出：25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25
//
// 示例 2：
//
// 输入：root = [4,9,0,5,1]
// 输出：1026
// 解释：
// 从根到叶子节点路径 4->9->5 代表数字 495
// 从根到叶子节点路径 4->9->1 代表数字 491
// 从根到叶子节点路径 4->0 代表数字 40
// 因此，数字总和 = 495 + 491 + 40 = 1026
//
// 提示：
//
// 树中节点的数目在范围 [1, 1000] 内
// 0 <= Node.val <= 9
// 树的深度不超过 10
func sumNumbers(root *TreeNode) int {
	var res int
	var path []int
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			res += slice2int(path)
			return
		}
		if node.Left != nil {
			traverse(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			traverse(node.Right)
			path = path[:len(path)-1]
		}
	}
	traverse(root)
	return res
}

func slice2int(slice []int) int {
	var res int
	for _, num := range slice {
		res = res*10 + num
	}
	return res
}

func sumNumbersII(root *TreeNode) int {
	var res int
	path := []int{root.Val}
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			res += slice2int(path)
			return
		}
		if node.Left != nil {
			path = append(path, node.Left.Val)
			traverse(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			path = append(path, node.Right.Val)
			traverse(node.Right)
			path = path[:len(path)-1]
		}
	}
	traverse(root)
	return res
}

/**
核心区别在于：
I：
进入递归时，记录本层值；
进入下一层递归处理左右孩子；
处理完左孩子时，撤销左孩子的路径选择，右孩子同理。
即，本层记录，返回上层时由上层撤销。

以三级树结构为例：
	   A
	  / \
	 B   C
	/
   D
执行流程：
1. 进入A层递归 -> append(A) -> path=[A]
2. 处理B子树 -> 进入B层递归 -> append(B) -> path=[A,B]
3. 处理D子树 -> 进入D层递归 -> append(D) -> path=[A,B,D] -> 叶子节点计算路径 -> 回溯到B层 -> path=[A,B]
4. B层递归结束 -> 回到A层 -> path截断B -> path=[A]
5. 处理C子树 -> 进入C层递归 -> append(C) -> path=[A,C] -> 叶子节点计算路径 -> 回溯到A层 -> path截断C -> path=[A]
6. A层递归结束 -> 由于"进入递归时，记录本层值"的逻辑，递归结束时，path=[根节点]

II：
递归前先记录根节点；
在每层递归中，对称记录/撤销对左孩子/右孩子的路径选择。
*/
