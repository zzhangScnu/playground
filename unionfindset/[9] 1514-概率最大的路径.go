package unionfindset

import "container/heap"

// 给你一个由 n 个节点（下标从 0 开始）组成的无向加权图，该图由一个描述边的列表组成，其中 edges[i] = [a, b] 表示连接节点 a 和 b
// 的一条无向边，且该边遍历成功的概率为 succProb[i] 。
//
// 指定两个节点分别作为起点 start 和终点 end ，请你找出从起点到终点成功概率最大的路径，并返回其成功概率。
//
// 如果不存在从 start 到 end 的路径，请 返回 0 。只要答案与标准答案的误差不超过 1e-5 ，就会被视作正确答案。
//
// 示例 1：
//
// 输入：n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0,
// end = 2
// 输出：0.25000
// 解释：从起点到终点有两条路径，其中一条的成功概率为 0.2 ，而另一条为 0.5 * 0.5 = 0.25
//
// 示例 2：
//
// 输入：n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0,
// end = 2
// 输出：0.30000
//
// 示例 3：
//
// 输入：n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
// 输出：0.00000
// 解释：节点 0 和 节点 2 之间不存在路径
//
// 提示：
//
// 2 <= n <= 10^4
// 0 <= start, end < n
// start != end
// 0 <= a, b < n
// a != b
// 0 <= succProb.length == edges.length <= 2*10^4
// 0 <= succProb[i] <= 1
// 每两个节点之间最多有一条边

type Vertex struct {
	node        int
	probability float64
}

type MaxHeap []*Vertex

func (h MaxHeap) Less(i, j int) bool {
	return h[i].probability > h[j].probability
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(*Vertex))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	graph := make([][][]float64, n)
	for i, edge := range edges {
		v, w := edge[0], edge[1]
		probability := succProb[i]
		graph[v] = append(graph[v], []float64{float64(w), probability})
		graph[w] = append(graph[w], []float64{float64(v), probability})
	}
	probabilities := make([]float64, n)
	probabilities[startNode] = 1
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, &Vertex{
		node:        startNode,
		probability: 1,
	})
	for maxHeap.Len() > 0 {
		cur := heap.Pop(maxHeap).(*Vertex)
		if cur.node == endNode {
			return cur.probability
		}
		if cur.probability < probabilities[cur.node] {
			continue
		}
		for _, to := range graph[cur.node] {
			if cur.probability*to[1] > probabilities[int(to[0])] {
				probabilities[int(to[0])] = cur.probability * to[1]
				heap.Push(maxHeap, &Vertex{
					node:        int(to[0]),
					probability: probabilities[int(to[0])],
				})
			}
		}
	}
	return 0
}

// 在最大概率路径问题中，初始概率设置为1而不是0的原因如下：
//
//1. 概率乘法性质 ：概率是相乘的关系，1是乘法的单位元。从起点到起点的概率应该是1（100%），因为这是确定的。
//2. 算法逻辑 ：后续计算中，路径概率是通过当前节点概率乘以边概率得到的。如果初始为0，所有计算结果都会是0，导致算法失效。
//3. 数学定义 ：在概率论中，一个事件自身的概率总是1，这是概率的基本性质。
//4. 与Dijkstra区别 ：在最短路径问题中，距离初始化为0是因为距离是相加的关系，0是加法的单位元。但在概率问题中，乘法关系决定了1才是合适的初始值。
