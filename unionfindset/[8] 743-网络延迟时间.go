package unionfindset

import (
	"container/heap"
	"math"
)

// 有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
// wi 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
// 示例 1：
//
// 输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
// 输出：2
//
// 示例 2：
//
// 输入：times = [[1,2,1]], n = 2, k = 1
// 输出：1
//
// 示例 3：
//
// 输入：times = [[1,2,1]], n = 2, k = 2
// 输出：-1
//
// 提示：
//
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][]int, n+1)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], []int{time[1], time[2]})
	}
	distance := make([]int, n+1)
	for i := 0; i <= n; i++ {
		distance[i] = math.MaxInt
	}
	distance[k] = 0
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, &Node{
		node:     k,
		distance: 0,
	})
	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(*Node)
		if cur.distance > distance[cur.node] {
			continue
		}
		for _, to := range graph[cur.node] {
			if distance[cur.node]+to[1] >= distance[to[0]] {
				continue
			}
			distance[to[0]] = distance[cur.node] + to[1]
			heap.Push(minHeap, &Node{
				node:     to[0],
				distance: distance[to[0]],
			})
		}
	}
	res := -1
	for i := 1; i <= n; i++ {
		if distance[i] == math.MaxInt {
			return -1
		}
		res = max(res, distance[i])
	}
	return res
}
