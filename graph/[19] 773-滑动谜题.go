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

// todo：基于-1，1，-n，n的坐标计算解法

/**
思路：
求解最短路径，就要想要BFS算法。
本题是通过相邻元素间的互相交换，来构建最短路径，从起始二维数组转移到目标二维数组。

因为BFS中的节点扩散、重复访问控制等都基于字符串，所以本题还涉及二维数组board的序列化，
以及序列化后，对单个节点(值为序列化的board)向外扩散至下一步节点(值为执行了一次相邻的两个数字的交换后的序列化的board)时，字符串元素间的交换操作。

1. 将board序列化并标记为已访问，入列；
2. 不断从队列中取队口节点，基于节点所表示的board现状，对数字0尝试进行四个方向的交换，即扩散出下一步节点，将其标记已访问且入列。
   其中如何对字符串即打平后的一维数组，执行原二维数组中的数字交换？
   因为本题是2 * 3的矩阵，可以直接枚举【一维数组中每个数字】在【原二维数组中的相邻数字在一维数组中对应的坐标】；
3. 在队列中，每次都需要按队列当前的大小size进行一圈扩散，确保当前队列中所有的节点都向外齐头迈进一步，
   在全体迈步之后才对step自增，确保寻找到target时的步数最少。
*/
