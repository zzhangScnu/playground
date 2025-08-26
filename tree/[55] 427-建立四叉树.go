package tree

// 给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。
//
// 你需要返回能表示矩阵 grid 的 四叉树 的根结点。
//
// 四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：
//
// val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False。注意，当 isLeaf 为 False 时，你可以把 True 或者
// False 赋值给节点，两种值都会被判题机制 接受 。
// isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。
//
//	class Node {
//	   public boolean val;
//
//   public boolean isLeaf;
//   public Node topLeft;
//   public Node topRight;
//   public Node bottomLeft;
//   public Node bottomRight;
// }
//
// 我们可以按以下步骤为二维区域构建四叉树：
//
// 如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后
// 停止。
// 如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
// 使用适当的子网格递归每个子节点。
//
// 如果你想了解更多关于四叉树的内容，可以参考 百科 。
//
// 四叉树格式：
//
// 你不需要阅读本节来解决这个问题。只有当你想了解输出格式时才会这样做。输出为使用层序遍历后四叉树的序列化形式，其中 null 表示路径终止符，其下面不存在节
// 点。
//
// 它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 [isLeaf, val] 。
//
// 如果 isLeaf 或者 val 的值为 True ，则表示它在列表 [isLeaf, val] 中的值为 1 ；如果 isLeaf 或者 val 的值为
// False ，则表示值为 0 。
//
// 示例 1：
//
// 输入：grid = [[0,1],[1,0]]
// 输出：[[0,1],[1,0],[1,1],[1,1],[1,0]]
// 解释：此示例的解释如下：
// 请注意，在下面四叉树的图示中，0 表示 false，1 表示 True 。
//
// 示例 2：
//
// 输入：grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,
// 1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
// 输出：[[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
//
// 解释：网格中的所有值都不相同。我们将网格划分为四个子网格。
// topLeft，bottomLeft 和 bottomRight 均具有相同的值。
// topRight 具有不同的值，因此我们将其再分为 4 个子网格，这样每个子网格都具有相同的值。
// 解释如下图所示：
//
// 提示：
//
// n == grid.length == grid[i].length
// n == 2ˣ 其中 0 <= x <= 6

type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func construct(grid [][]int) *QuadTreeNode {
	n := len(grid)
	var isUniform func(i, j int, length int) bool
	isUniform = func(x, y int, length int) bool {
		val := grid[x][y]
		for i := x; i < x+length; i++ {
			for j := y; j < y+length; j++ {
				if grid[i][j] != val {
					return false
				}
			}
		}
		return true
	}
	var traverse func(i, j int, length int) *QuadTreeNode
	traverse = func(i, j int, length int) *QuadTreeNode {
		if isUniform(i, j, length) {
			return &QuadTreeNode{
				Val:    grid[i][j] == 1,
				IsLeaf: true,
			}
		}
		nextLength := length / 2
		return &QuadTreeNode{
			IsLeaf:      false,
			TopLeft:     traverse(i, j, nextLength),
			TopRight:    traverse(i, j+nextLength, nextLength),
			BottomLeft:  traverse(i+nextLength, j, nextLength),
			BottomRight: traverse(i+nextLength, j+nextLength, nextLength),
		}
	}
	return traverse(0, 0, n)
}

/**
思路：
采用分治-递归的方法，每次将区域拆解为4个子区域；
直至触底反弹，从叶子节点开始构建子树，最终得到根节点。

在递归方法的入口处先调用【是否区域中所有元素均相等】的判断，
实际上隐含了base case，即当区域中元素个数为1的情况，此时会判定为叶子节点且返回。
*/
