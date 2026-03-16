package graph

var visited []bool
var path []bool
var hasCircle bool

func detectCircle(graph [][]int) bool {
	n := len(graph)
	hasCircle = false
	visited, path = make([]bool, n), make([]bool, n)
	for i := 0; i < n; i++ { // 这里是为了处理多联通子图，防止有子图未遍历到而导致遗漏
		traverse(graph, i)
	}
	return hasCircle
}

func traverse(graph [][]int, vertex int) {
	if path[vertex] {
		hasCircle = true
		return
	}
	if visited[vertex] {
		return
	}
	visited[vertex] = true
	path[vertex] = true
	for _, neighbor := range graph[vertex] {
		traverse(graph, neighbor)
	}
	path[vertex] = false // 已遍历完当前节点的所有邻居，即所有子路径的可能性，需回退到上一个节点，探索其他路径
}

/**
思路：
DFS：路径遍历，递归+回溯，检查是否有同一个节点在同一条路径中被重复访问；
- visited数组：标记节点是否已访问，避免重复访问导致死循环（地图上 “已经探索过的区域”，无需重复探索）；
- path数组：记录当前路径上已访问的节点，即单次递归的堆栈中的节点。在递归遍历某个节点的邻居时，发现这个邻居节点已经被标记在「当前递归路径（path）」中，则表示有环（当前 “走的这条小路”，是否绕回了自己的脚印）。
  当单次遍历路径已扩展到无下一个可用节点，则需要进行回溯，回到先前的节点，进行其他节点的选择和访问。

注意：
1. 因为图可能是非联通图，所以需要对图中的每一个节点进行遍历，防止遗漏；
2. 因为将辅助数组和结果变量作为全局变量维护，所以在每次函数调用入口都需要初始化一遍，否则调用间会相互影响；
3. 判断环的时机：
之前的写法是：
func traverse(graph [][]int, vertex int) {
	if visited[vertex] {
		return
	}
	visited[vertex] = true
	if path[vertex] {
		hasCircle = true
		return
	}
	path[vertex] = true
	for _, neighbor := range graph[vertex] {
		traverse(graph, neighbor)
	}
	path[vertex] = false
}
这种写法无法检测环，因为对于环来说，同一个节点一定会被访问多次。
如果一开始先检测重复访问，则下面的环检测会被阻断。
所以应先处理环检测，如果无环再判断是否重复访问。
*/
