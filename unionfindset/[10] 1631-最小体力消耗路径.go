package unionfindset

import (
	"container/heap"
	"math"
)

// 你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row,
// col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你
// 每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
//
// 一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
//
// 请你返回从左上角走到右下角的最小 体力消耗值 。
//
// 示例 1：
//
// 输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
// 输出：2
// 解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
// 这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
//
// 示例 2：
//
// 输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
// 输出：1
// 解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
//
// 示例 3：
//
// 输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
// 输出：0
// 解释：上图所示路径不需要消耗任何体力。
//
// 提示：
//
// rows == heights.length
// columns == heights[i].length
// 1 <= rows, columns <= 100
// 1 <= heights[i][j] <= 10⁶

type Vertex2D struct {
	x, y   int
	effort int
}

type Vertex2DMinHeap []*Vertex2D

func (v Vertex2DMinHeap) Len() int {
	return len(v)
}

func (v Vertex2DMinHeap) Less(i, j int) bool {
	return v[i].effort < v[j].effort
}

func (v Vertex2DMinHeap) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v *Vertex2DMinHeap) Push(x any) {
	*v = append(*v, x.(*Vertex2D))
}

func (v *Vertex2DMinHeap) Pop() any {
	arr := *v
	n := len(arr)
	x := arr[n-1]
	*v = arr[0 : n-1]
	return x
}

var movements = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	efforts := make([][]int, m)
	for i := 0; i < m; i++ {
		efforts[i] = make([]int, n)
		for j := 0; j < n; j++ {
			efforts[i][j] = math.MaxInt
		}
	}
	efforts[0][0] = 0
	minHeap := &Vertex2DMinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, &Vertex2D{
		x:      0,
		y:      0,
		effort: 0,
	})
	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(*Vertex2D)
		if cur.x == m-1 && cur.y == n-1 {
			return cur.effort
		}
		if cur.effort > efforts[cur.x][cur.y] {
			continue
		}
		for _, movement := range movements {
			x, y := cur.x+movement[0], cur.y+movement[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			effort := max(efforts[cur.x][cur.y],
				int(math.Abs(float64(heights[x][y]-heights[cur.x][cur.y]))))
			if effort < efforts[x][y] {
				efforts[x][y] = effort
				heap.Push(minHeap, &Vertex2D{
					x:      x,
					y:      y,
					effort: effort,
				})
			}
		}
	}
	return 0
}

//在最小体力消耗路径问题中，取max值的原因如下：
//
//1. 问题定义 ：我们需要找到从起点到终点的路径，使得路径上的最大高度差最小。因此需要跟踪路径上的最大高度差。
//2. 路径性质 ：一条路径的体力消耗由该路径上的最大高度差决定，而不是所有高度差的总和。
//3. 算法逻辑 ：
//
//   - efforts[cur.x][cur.y] 表示到达当前节点的路径上的最大高度差
//   - math.Abs(heights[x][y]-heights[cur.x][cur.y]) 是当前移动的高度差
//   - 取两者的max值，确保记录的是新路径上的最大高度差
//4. 与Dijkstra区别 ：
//
//   - 传统Dijkstra是求路径和的最小值
//   - 这里求的是路径上最大值的最小值，因此需要取max而不是sum
