package graph

import (
	"strconv"
	"strings"
)

// 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示。一次 移动 定义为选择 0 与一个相邻的数字（
// 上下左右）进行交换.
//
// 最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
//
// 给出一个谜板的初始状态 board ，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
//
// 示例 1：
//
// 输入：board = [[1,2,3],[4,0,5]]
// 输出：1
// 解释：交换 0 和 5 ，1 步完成
//
// 示例 2:
//
// 输入：board = [[1,2,3],[5,4,0]]
// 输出：-1
// 解释：没有办法完成谜板
//
// 示例 3:
//
// 输入：board = [[4,1,2],[5,0,3]]
// 输出：5
// 解释：
// 最少完成谜板的最少移动次数是 5 ，
// 一种移动路径:
// 尚未移动: [[4,1,2],[5,0,3]]
// 移动 1 次: [[4,1,2],[0,5,3]]
// 移动 2 次: [[0,1,2],[4,5,3]]
// 移动 3 次: [[1,0,2],[4,5,3]]
// 移动 4 次: [[1,2,0],[4,5,3]]
// 移动 5 次: [[1,2,3],[4,5,0]]
//
// 提示：
//
// board.length == 2
// board[i].length == 3
// 0 <= board[i][j] <= 5
// board[i][j] 中每个值都 不同
func slidingPuzzle(board [][]int) int {
	movements := [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	target := "123450"
	start := serialize(board)
	if target == start {
		return 0
	}
	queue := []string{start}
	visited := make(map[string]struct{})
	visited[start] = struct{}{}
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			for j := 0; j < len(cur); j++ {
				if cur[j] != '0' {
					continue
				}
				for _, k := range movements[j] {
					next := swap(cur, j, k)
					if _, ok := visited[next]; ok {
						continue
					}
					visited[next] = struct{}{}
					queue = append(queue, next)
				}
			}
		}
		step++
	}
	return -1
}

func swap(elements string, i, j int) string {
	bytes := []byte(elements)
	bytes[i], bytes[j] = bytes[j], bytes[i]
	return string(bytes)
}

func serialize(board [][]int) string {
	m, n := len(board), len(board[0])
	var elements []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			elements = append(elements, board[i][j])
		}
	}
	var sb strings.Builder
	for _, element := range elements {
		sb.WriteString(strconv.Itoa(element))
	}
	return sb.String()
}
