package graph

import (
	"fmt"
	"strconv"
	"strings"
)

/**
题目描述
给定一个由 1（陆地）和 0（水）组成的矩阵，岛屿指的是由水平或垂直方向上相邻的陆地单元格组成的区域，且完全被水域单元格包围。孤岛是那些位于矩阵内部、所有单元格都不接触边缘的岛屿。

现在你需要计算所有孤岛的总面积，岛屿面积的计算方式为组成岛屿的陆地的总数。

输入描述
第一行包含两个整数 N, M，表示矩阵的行数和列数。之后 N 行，每行包含 M 个数字，数字为 1 或者 0。

输出描述
输出一个整数，表示所有孤岛的总面积，如果不存在孤岛，则输出 0。

输入示例
4 5
1 1 0 0 0
1 1 0 0 0
0 0 1 0 0
0 0 0 1 1

输出示例
1
*/

func main() {
	movements := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	graph, visited := initiateGraph()
	m, n := len(graph), len(graph[0])
	var flood func(graph [][]string, visited [][]bool, x, y int)
	flood = func(graph [][]string, visited [][]bool, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if visited[x][y] {
			return
		}
		if graph[x][y] == "0" {
			return
		}
		visited[x][y] = true
		graph[x][y] = "0"
		for _, movement := range movements {
			flood(graph, visited, x+movement[0], y+movement[1])
		}
	}
	for j := 0; j < n; j++ {
		flood(graph, visited, 0, j)
		flood(graph, visited, m-1, j)
	}
	for i := 0; i < m; i++ {
		flood(graph, visited, i, 0)
		flood(graph, visited, i, n-1)
	}
	var res int
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if graph[i][j] == "1" {
				res++
			}
		}
	}
	fmt.Print(res)
}

func initiateGraph() ([][]string, [][]bool) {
	var size string
	_, _ = fmt.Scan(&size)
	arr := strings.Split(size, " ")
	m, _ := strconv.Atoi(arr[0])
	n, _ := strconv.Atoi(arr[1])
	graph, visited := make([][]string, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		var content string
		_, _ = fmt.Scan(&content)
		values := strings.Split(content, " ")
		graph[i] = make([]string, n)
		for j := 0; j < n; j++ {
			graph[i][j] = values[j]
		}
		visited[i] = make([]bool, n)
	}
	return graph, visited
}
