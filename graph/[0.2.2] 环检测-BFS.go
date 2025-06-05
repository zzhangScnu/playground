package graph

func detectCircleBFS(graph [][]int) bool {
	n := len(graph)
	if n == 0 {
		return false
	}
	inDegree := make([]int, n)
	var queue []int
	for _, tos := range graph {
		for _, to := range tos {
			inDegree[to]++
		}
	}
	for vertex, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}
	var count int
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range graph[cur] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return count != n
}

/**
思路：
BFS的做法像是在图中找到起点，从起点开始抽丝剥茧，顺藤摸瓜，遍历所有节点。
就像一棵二叉树，先找到根节点，从其开始进行自上而下的遍历，最终到达叶子节点，完成图的整体遍历。

1. 遍历图，统计每个节点的入度；
2. 遍历各节点的入度，如果入度 == 0，表示该节点没有被其他节点指向，就像栈中的栈口元素一样，其上没有其他元素压制，可优先直接访问。
	此时将其入列；
3. 维护一个节点count，表示已遍历的节点数量；
4. 将队列中的节点逐一出列，对count++，且将节点的相邻节点们的入度--。
	对于入度变为0的节点，需入列等待访问；
5. 循环直至队列为空；
6. 最后判断count和图中的节点数量是否相等，相等则表示每个节点恰好被访问1次，即图中无环。
	若不相等，可能是当队列为空时，图中仅剩入度非0的节点，没有产生可以入列的节点，算法终止。此时图中仍有节点未被访问，即存在环。

本题无需维护visited数组，因为BFS向外层层扩展的特性，且限制了只有入度为0的节点才能入列，
天然不会走回头路，不存在重复访问节点的情况。

注意，本方法是检测是否有环，所以最终返回的结果是count != n。
*/
