package graph

import (
	"slices"
)

var res []int

func topologicalSort(graph [][]int) []int {
	n := len(graph)
	visited, res = make([]bool, n), make([]int, n)
	for v := 0; v < n; v++ {
		traverseInTopologicalSort(graph, v)
		if hasCircle {
			return []int{}
		}
	}
	slices.Reverse(res)
	return res
}

func traverseInTopologicalSort(graph [][]int, vertex int) {
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
		traverseInTopologicalSort(graph, neighbor)
	}
	res = append(res, vertex)
	path[vertex] = false
}

/**
思路：
在环检测算法的基础上，在后序位置增加节点访问结果收集，
最终再将结果集反转，即为拓扑排序结果。

1. 为什么是后序？
图的DFS遍历其实就是二叉树DFS遍历的扩展。
以二叉树为例，将二叉树看作任务间的依赖关系图，根节点依赖左右孩子。
二叉树后序位置收集结果，会在收集完左子树+右子树结果后，再收集根节点结果，
这一特点强要求一个任务必须等到其依赖的所有任务都完成后才能开始执行。
后序遍历结束后，二叉树[1, 2, 3]的结果为[2, 3, 1]。

2. 为什么需要反转结果？
如上述二叉树例子，反转结果后得到[1, 3, 2]，符合要求，即从1出发，依次遍历相邻节点。

注意，如果图中有环，则显然无法完成拓扑排序，此时应显式返回空结果集。
*/
