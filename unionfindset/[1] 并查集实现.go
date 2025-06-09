package unionfindset

type UnionFindSet struct {
	parent []int
	count  int
}

func NewUnionFindSet(nodes []int) *UnionFindSet {
	parent := make([]int, len(nodes))
	for _, node := range nodes {
		parent[node] = node
	}
	return &UnionFindSet{
		parent: parent,
		count:  len(nodes),
	}
}

func (u *UnionFindSet) union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return
	}
	u.parent[rootX] = rootY
	u.count--
}

func (u *UnionFindSet) isConnected(x, y int) bool {
	rootX, rootY := u.find(x), u.find(y)
	return rootX == rootY
}

func (u *UnionFindSet) find(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}
