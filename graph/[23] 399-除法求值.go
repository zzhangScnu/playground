package graph

// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和
// values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
//
// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj =
// ? 的结果作为答案。
//
// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替
// 代这个答案。
//
// 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
//
// 注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。
//
// 示例 1：
//
// 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"]
// ,["b","a"],["a","e"],["a","a"],["x","x"]]
// 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
// 解释：
// 条件：a / b = 2.0, b / c = 3.0
// 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
// 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
// 注意：x 是未定义的 => -1.0
//
// 示例 2：
//
// 输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0],
// queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
// 输出：[3.75000,0.40000,5.00000,0.20000]
//
// 示例 3：
//
// 输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],[
// "a","c"],["x","y"]]
// 输出：[0.50000,2.00000,-1.00000,-1.00000]
//
// 提示：
//
// 1 <= equations.length <= 20
// equations[i].length == 2
// 1 <= Ai.length, Bi.length <= 5
// values.length == equations.length
// 0.0 < values[i] <= 20.0
// 1 <= queries.length <= 20
// queries[i].length == 2
// 1 <= Cj.length, Dj.length <= 5
// Ai, Bi, Cj, Dj 由小写英文字母与数字组成
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := buildGraph(equations, values)
	var traverse func(from, to string, visited map[string]bool) float64
	traverse = func(from, to string, visited map[string]bool) float64 {
		if graph[from] == nil || graph[to] == nil {
			return -1.0
		}
		if visited[from] {
			return -1.0
		}
		visited[from] = true
		defer delete(visited, from)
		if value, ok := graph[from][to]; ok {
			return value
		}
		for nextFrom, nextValue := range graph[from] {
			subValue := traverse(nextFrom, to, visited)
			if subValue != -1 {
				return nextValue * subValue
			}
		}
		return -1.0
	}
	var res []float64
	for _, query := range queries {
		from, to := query[0], query[1]
		res = append(res, traverse(from, to, make(map[string]bool)))
	}
	return res
}

func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	n := len(values)
	graph := make(map[string]map[string]float64)
	for i := 0; i < n; i++ {
		equation, value := equations[i], values[i]
		from, to := equation[0], equation[1]
		if graph[from] == nil {
			graph[from] = make(map[string]float64)
		}
		if graph[to] == nil {
			graph[to] = make(map[string]float64)
		}
		graph[from][to], graph[to][from], graph[from][from], graph[to][to] = value, 1/value, 1.0, 1.0
	}
	return graph
}

/**
思路：
表达式求值 ->
	操作数 -> 节点
	值 -> 边

注意：
访问过的节点返回 - 1，是针对【当前搜索路径】的限制，防止陷入无限循环。因为在处理逻辑中，只有子问题返回非-1时才会继续深入搜索。
回溯操作即清除标记，保证了节点可以在其他路径中被重新使用。
注意visited需在对每个表达式求值时重新初始化，避免结果相互污染。
这种设计既避免了无限递归，又不会错过正确结果，是图遍历的标准处理方式。

为什么是乘法处理nextValue * subValue？
假设有如下关系：
a/b = 2 （即 a = 2b）
b/c = 3 （即 b = 3c）
求 a/c 的值：
根据传递性：a = 2b = 2 × (3c) = 6c → a/c = 6
对应代码逻辑：nextValue=2（a→b），subValue=3（b→c），结果为 2×3=6
*/

/**
visited[from] = true
defer delete(visited, from)
// 处理逻辑

等同于
visited[from] = true
// 处理逻辑
visited[from] = false
*/
