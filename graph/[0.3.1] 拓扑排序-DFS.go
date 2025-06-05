package graph

import (
	"slices"
)

var res []int

func topologicalSort(graph [][]int) []int {
	n := len(graph)
	hasCircle = false
	visited, path, res = make([]bool, n), make([]bool, n), make([]int, 0, n)
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

注意，
0. 如果图中有环，则显然无法完成拓扑排序，此时应显式返回空结果集；
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
4. 切片初始化的两种方式：
- make([]int, 0, n)：分配一个长度 == 0，容量 == n的切片，append正常，但索引访问会导致空指针异常；
- make([]int, n)：分配一个长度 == n，容量 == n的切片，其中每个元素均为0，append会导致扩容且前序元素均为0，索引访问正常。
*/
