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
//		class Node {
//		   public boolean val;
//
//	  public boolean isLeaf;
//	  public Node topLeft;
//	  public Node topRight;
//	  public Node bottomLeft;
//	  public Node bottomRight;
//	}
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
func constructII(grid [][]int) *QuadTreeNode {
	n := len(grid)
	var newPrefixSum func() [][]int
	newPrefixSum = func() [][]int {
		prefixSum := make([][]int, n+1)
		for i := 0; i <= n; i++ {
			prefixSum[i] = make([]int, n+1)
		}
		for i := 0; i < n; i++ {
			rowSum := 0
			for j := 0; j < n; j++ {
				rowSum += grid[i][j]
				prefixSum[i+1][j+1] = prefixSum[i][j+1] + rowSum
			}
		}
		return prefixSum
	}
	prefixSum := newPrefixSum()
	var isUniform func(i, j int, length int) bool
	isUniform = func(x, y int, length int) bool {
		sum := prefixSum[x+length][y+length] + prefixSum[x][y] - prefixSum[x+length][y] - prefixSum[x][y+length]
		return sum == 0 || sum == length*length
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
在原始做法中，每次都需要遍历当前区域，从而判定是否叶子节点，时间复杂度为O(k^2)，k为区域边长。
可以采用前缀和做法，用O(n^2)的整体空间，换每次判定的O(1)时间。

前缀和矩阵定义：
- prefixSum大小：O((n + 1) * (n + 1))
- prefixSum[i][j]含义：表示从(0, 0)到(i - 1, j - 1)形成的区域中元素的和；
为什么不是初始化为n * n的大小？为什么prefixSum[i][j]不是表示从(0, 0)到(i, j)形成的区域中元素的和？
因为由前缀和计算通用思想可知，当前区域前缀和依赖于前一区域前缀和推导而来，如果不在空间上冗余，就需要在代码中显式处理边界逻辑。

前缀和矩阵初始化：
固定第i行：
	遍历当前行中的每一列：
		1. 累加当前行的元素和，即rowSum += grid[i][0...j]
		2. 对于当前行 * 列，即grid[i][j]，从(0, 0)到(i, j)形成的区域中元素的和表示为prefixSum[i + 1][j + 1]，
           故prefixSum[i + 1][j + 1] = prefixSum[i][j + 1] + rowSum
		   即逐步通过扩展行，来完成前缀和矩阵的计算。
注意行和列遍历的索引范围都是[0, n)，即n不包含。这样才能兼容prefixSum[i + 1][j + 1]可能存在的数组越界异常。

前缀和矩阵使用：
从(i, j)到(i + length - 1, j + length - 1)形成的区域中元素的和 =
	prefixSum[i + length][j + length] // 大矩形
	- prefixSum[i + length][j] // 左下矩形
	- prefixSum[i][j + length] // 右上矩形
	+ prefixSum[i][j] // 左下矩形 * 右上矩形中重叠的小矩形
*/
