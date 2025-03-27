package graph

import (
	"fmt"
	"strings"
)

/**
题目描述：

给定一个由 1（陆地）和 0（水）组成的矩阵，岛屿指的是由水平或垂直方向上相邻的陆地单元格组成的区域，且完全被水域单元格包围。
孤岛是那些位于矩阵内部、所有单元格都不接触边缘的岛屿。
现在你需要将所有孤岛“沉没”，即将孤岛中的所有陆地单元格（1）转变为水域单元格（0）。

输入描述：
第一行包含两个整数 N, M，表示矩阵的行数和列数。
之后 N 行，每行包含 M 个数字，数字为 1 或者 0，表示岛屿的单元格。

输出描述
输出将孤岛“沉没”之后的岛屿矩阵。

输入示例：
4 5
1 1 0 0 0
1 1 0 0 0
0 0 1 0 0
0 0 0 1 1

输出示例：
1 1 0 0 0
1 1 0 0 0
0 0 0 0 0
0 0 0 1 1

数据范围：
1 <= M, N <= 50
*/

func main() {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	graph, _ := initiateGraph()
	m, n := len(graph), len(graph[0])
	var traverse func(graph [][]string, x, y int)
	traverse = func(graph [][]string, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if graph[x][y] == "0" || graph[x][y] == "-1" {
			return
		}
		graph[x][y] = "-1"
		for _, movement := range movements {
			traverse(graph, x+movement[0], y+movement[1])
		}
	}
	for i := 0; i < m; i++ {
		traverse(graph, i, 0)
		traverse(graph, i, n-1)
	}
	for j := 0; j < n; j++ {
		traverse(graph, 0, j)
		traverse(graph, m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch graph[i][j] {
			case "1":
				graph[i][j] = "0"
			case "-1":
				graph[i][j] = "1"

			}
		}
		fmt.Printf("%s\n", strings.Join(graph[i], " "))
	}
}

/**
思路：
需要沉没孤岛，可以转换一下思路：
将周边的岛屿标记为特殊的值 -> 矩阵中的陆地仅剩孤岛
	-> 遍历矩阵 + 孤岛【1 -> 0】 + 周边岛屿【-1 -> 1】。

注意，
if graph[x][y] == "0" || graph[x][y] == "-1" {
	return
}
不要忘了后半截判断条件，否则会导致重复访问进而无限循环。
*/
