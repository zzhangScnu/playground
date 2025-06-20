package graph

import (
	"fmt"
	"strings"
)

// 给你一个由'1'(陆地)和'0'(水)组成的的二维网格,请你你计算网格中岛屿的数量。
// 岛屿总是被水包围,并且每座岛屿只能由水平方向和/或竖直直方向上相邻的陆地连接形成。此外,你可以假设
// 该网格的四条边均被水包围。
func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make(map[string]struct{})
	sb := strings.Builder{}
	var traverse func(grid [][]int, x, y int, direction string)
	traverse = func(grid [][]int, x, y int, direction string) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		sb.WriteString(direction)
		traverse(grid, x-1, y, "1")
		traverse(grid, x+1, y, "2")
		traverse(grid, x, y-1, "3")
		traverse(grid, x, y+1, "4")
		sb.WriteString(fmt.Sprintf("-%s", direction))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sb.Reset()
				traverse(grid, i, j, "-1")
				memo[sb.String()] = struct{}{}
			}
		}
	}
	return len(memo)
}

/**
思路：
在遍历岛屿过程中，对岛屿进行序列化并记录到集合中，最终集合大小即为不同岛屿数量。
跟二叉树的序列化不同，二叉树每个节点都有自己的值，直接记录即可。
而本题二维矩阵中，不仅陆地的值全为1，且因为需要识别岛屿的形状，所以需要生成其遍历路径并记录。

需要记录：
进入陆地的方向 + 离开陆地的方向。
即在前序位置收集 + 在后序位置收集。
同时需要额外通过参数告诉递归函数，当前执行的是哪个方向的遍历。
这就导致本题用枚举四个方向、写四次函数调用的形式，代替for循环遍历方向数组。
（其实for循环也可以，在movements数组中为每个方向增加一个方向标记）

如果只收集进入方向，无法唯一标识岛屿形状。
如进入路径都是下+右的岛屿，可以是正方形的左上角或左下角的形状。
*/
