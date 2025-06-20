package unionfindset

// 给你两个 m x n 的二进制矩阵 grid1 和 grid2 ，它们只包含 0 （表示水域）和 1 （表示陆地）。一个 岛屿 是由 四个方向 （水平或者竖
// 直）上相邻的 1 组成的区域。任何矩阵以外的区域都视为水域。
//
// 如果 grid2 的一个岛屿，被 grid1 的一个岛屿 完全 包含，也就是说 grid2 中该岛屿的每一个格子都被 grid1 中同一个岛屿完全包含，那
// 么我们称 grid2 中的这个岛屿为 子岛屿 。
//
// 请你返回 grid2 中 子岛屿 的 数目 。
//
// 示例 1：
// 输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]],
// grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
// 输出：3
// 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
// grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。
//
// 示例 2：
// 输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]],
// grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
// 输出：2
// 解释：如上图所示，左边为 grid1 ，右边为 grid2 。
// grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。
//
// 提示：
//
// m == grid1.length == grid2.length
// n == grid1[i].length == grid2[i].length
// 1 <= m, n <= 500
// grid1[i][j] 和 grid2[i][j] 都要么是 0 要么是 1 。
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid1), len(grid1[0])
	unionFindSet := NewUnionFindSet(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				for _, movement := range movements {
					x, y := i+movement[0], j+movement[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid2[x][y] == 1 {
						unionFindSet.Union(i*n+j, x*n+y)
					}
				}
			}
		}
	}
	var count int
	visited := make(map[int]bool)
	isSub := true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 {
				continue
			}
			root := unionFindSet.find(i*n + j)
			if visited[root] {
				continue
			}
			visited[root] = true
			isSub = true
			for x := 0; x < m && isSub; x++ {
				for y := 0; y < n && isSub; y++ {
					if unionFindSet.IsConnected(root, x*n+y) && grid2[x][y] == 1 && grid1[x][y] == 0 {
						isSub = false
					}
				}
			}
			if isSub {
				count++
			}
		}
	}
	return count
}

// 本题对于大数据集会超时，扩展下思路就好

/**
思路：并查集

如果b是a的子岛屿，则b陆地是a陆地的子集，能被a陆地完全覆盖，
即b中的每一块陆地在a中必然是陆地。

可以理解为i&j和x&y都是在遍历grid2，
i&j：定位grid2中的陆地网格，通过grid[i][j]找到该陆地网格对应的岛屿在并查集中的根节点root；
x&y：遍历grid2中的网格，试图将root下该岛屿的所有陆地网格筛选出来。
1. 将grid2中的岛屿都加入并查集，每一个岛屿最终会形成一个连通分量；
2. 对grid2中的每一个网格进行判断：
	- 如果该网格(i, j)是陆地；
	- 如果grid1该位置对应的网格(x, y)是海水；
	- 如果(x, y)与(i, j)属于同一个联通分量，即(x, y)有可能是(i, j)父岛屿的一部分；
   则不满足子岛屿条件。
3. 性能优化
   因为对于grid2中的某个网格来说，通过其能推断出其从属的一整片岛屿是否符合要求。
   当该岛屿的判断结果更新后，其下的其他网格无需进行重复判断。
   实现方式：
   对于每一个root，在访问时录入visited，即对每个岛屿仅访问一次，下次通过其他陆地网格关联到相同岛屿的root时，则直接跳过。
   维护isSub标志位，初始化为true；
   将isSub嵌入for循环执行的判断条件，当grid2中的岛屿不是grid1中的子岛屿时，isSub被置false时，跳出循环，不再继续执行判断。
   如果遍历完x&y后，isSub仍为true，则说明grid2中以root为根节点的岛屿找到了其在grid1中的父岛屿，count自增。

注意：
1. 预处理并查集时，对陆地网格四个方向的遍历，需注意增加越界控制；
2. isSub不能嵌套到最外层对i&j的遍历，因为grid2需整体遍历一次，才能覆盖到所有岛屿。
*/
