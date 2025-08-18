package graph

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
//
// 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
//
//	class Node {
//	   public int val;
//	   public List<Node> neighbors;
//	}
//
// 测试用例格式：
//
// 简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻
// 接列表表示。
//
// 邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
//
// 给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。
//
// 示例 1：
//
// 输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
// 输出：[[2,4],[1,3],[2,4],[1,3]]
// 解释：
// 图中有 4 个节点。
// 节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
// 节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
// 节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
// 节点 4 的值是 4，它有两个邻居：节点 1 和 3 。
//
// 示例 2：
//
// 输入：adjList = [[]]
// 输出：[[]]
// 解释：输入包含一个空列表。该图仅仅只有一个值为 1 的节点，它没有任何邻居。
//
// 示例 3：
//
// 输入：adjList = []
// 输出：[]
// 解释：这个图是空的，它不含任何节点。
//
// 提示：
//
// 这张图中的节点数在 [0, 100] 之间。
// 1 <= Node.val <= 100
// 每个节点值 Node.val 都是唯一的，
// 图中没有重复的边，也没有自环。
// 图是连通图，你可以从给定节点访问到所有节点。
func cloneGraphII(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := []*Node{node}
	mapping := make(map[*Node]*Node)
	mapping[node] = &Node{Val: node.Val}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, neighbor := range cur.Neighbors {
			if _, ok := mapping[neighbor]; !ok {
				clone := &Node{Val: neighbor.Val}
				mapping[neighbor] = clone
				queue = append(queue, neighbor)
			}
			mapping[cur].Neighbors = append(mapping[cur].Neighbors, mapping[neighbor])
		}
	}
	return mapping[node]
}

/**
思路：
迭代处理图的克隆
通过队列+哈希表去重的方式，避免重复创建节点&陷入死循环。
注意最外层先处理第一个节点，再开启循环，处理队列中的邻接节点。
因为图是联通图，一个节点必然是其他节点的邻接节点，所以可以保证所有节点都在循环中被处理。

为什么要在队列外提前处理“根节点”？
为了统一队列中的处理逻辑。

对于当前节点的每一个邻接节点neighbor：
1. 需判断neighbor是否曾被克隆过，如果有，则跳过此次克隆；否则，维护【原节点-克隆节点】即【neighbor-clone】映射，并将neighbor加入队列，以便后续扩散继续处理neighbor的下一层邻接节点；
2. 无论neighbor先前是否曾被克隆过，都应该处理cur的克隆节点的邻接节点列表，mapping[cur].Neighbors。即将已有的或本次新建的neighbor的clone节点，mapping[neighbor]纳入该列表中。
*/
