package dynamicprogramming

import "math"

// 有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城
// 市 fromi 开始，以价格 pricei 抵达 toi。
//
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便
// 宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
// 示例 1：
//
// 输入:
// n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0,
// dst = 3, k = 1
// 输出: 700
// 解释: 城市航班图如上
// 从城市 0 到城市 3 经过最多 1 站的最佳路径用红色标记，费用为 100 + 600 = 700。
// 请注意，通过城市 [0, 1, 2, 3] 的路径更便宜，但无效，因为它经过了 2 站。
//
// 示例 2：
//
// 输入:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
// 输出: 200
// 解释:
// 城市航班图如上
// 从城市 0 到城市 2 经过最多 1 站的最佳路径标记为红色，费用为 100 + 100 = 200。
//
// 示例 3：
//
// 输入：n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
// 输出：500
// 解释：
// 城市航班图如上
// 从城市 0 到城市 2 不经过站点的最佳路径标记为红色，费用为 500。
//
// 提示：
//
// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 10⁴
// 航班没有重复，且不存在自环
// 0 <= src, dst, k < n
// src != dst
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := buildGraph(n, flights)
	memo := make(map[[2]int]int)
	var dp func(dst, k int) int
	dp = func(dst, k int) int {
		if dst == src {
			return 0
		}
		if k < 0 {
			return -1
		}
		if res, ok := memo[[2]int{dst, k}]; ok {
			return res
		}
		res := math.MaxInt
		for _, prev := range graph[dst] {
			from, price := prev[0], prev[1]
			sub := dp(from, k-1)
			if sub == -1 {
				continue
			}
			res = min(res, sub+price)
		}
		if res == math.MaxInt {
			res = -1
		}
		memo[[2]int{dst, k}] = res
		return res
	}
	return dp(dst, k)
}

func buildGraph(n int, flights [][]int) [][][]int {
	graph := make([][][]int, n)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		graph[to] = append(graph[to], []int{from, price})
	}
	return graph
}

/**
解法一：
从终点倒推起点
src -> node[...] -> dst
for 每个node -> dst中的node(from)
	dp(src -> dst, k) =
		min(dp(src -> node, k - 1) + weight[node][dst])

- 从递推公式可以看出，变化的变量只有dst和k；
- k - 1是因为子问题中的中转次数比当前问题少1次；
- base case：当k < 0时，即中转次数使用完毕，此时返回特殊标识-1，表示不可达。
*/

func findCheapestPriceFromSrc(n int, flights [][]int, src int, dst int, k int) int {
	graph := buildGraphFrom(n, flights)
	memo := make(map[[2]int]int)
	var dp func(src, k int) int
	dp = func(src, k int) int {
		if src == dst {
			return 0
		}
		if k < 0 {
			return -1
		}
		if res, ok := memo[[2]int{src, k}]; ok {
			return res
		}
		res := math.MaxInt
		for _, next := range graph[src] {
			to, price := next[0], next[1]
			sub := dp(to, k-1)
			if sub == -1 {
				continue
			}
			res = min(res, sub+price)
		}
		if res == math.MaxInt {
			res = -1
		}
		memo[[2]int{src, k}] = res
		return res
	}
	return dp(src, k)
}

func buildGraphFrom(n int, flights [][]int) [][][]int {
	graph := make([][][]int, n)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		graph[from] = append(graph[from], []int{to, price})
	}
	return graph
}

/**
解法二：
从起点正推终点
src -> node[...] -> dst
for 每个src -> node中的node(to)
	dp(src -> dst, k) =
		min(dp(node -> dst, k - 1) + weight[src][node])

- 从递推公式可以看出，变化的变量只有src和k；
- k - 1是因为子问题中的中转次数比当前问题少1次；
- base case：当k < 0时，即中转次数使用完毕，此时返回特殊标识-1，表示不可达。
*/
