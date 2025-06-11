package unionfindset

// 给定从0到n-1标号的n个结点,和一个无向边列表(每条边以结点对来表示),请编写一个函数用来判断
// 这些边是否能够形成一个合法有效的树结构。

// 输入:n=5,边列表edges=[[0,1],[0,1],[0,2],[0,3],[1,4]]
// 输出:true

func validTree(n int, edges [][]int) bool {
	unionFindSet := NewUnionFindSet(n)
	for _, edge := range edges {
		if unionFindSet.isConnected(edge[0], edge[1]) {
			return false
		}
		unionFindSet.union(edge[0], edge[1])
	}
	return unionFindSet.size() == 1
}
