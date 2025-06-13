package unionfindset

type UnionFindSet struct {
	parent []int
	count  int
}

func NewUnionFindSet(n int) *UnionFindSet {
	parent := make([]int, n)
	for node := 0; node < n; node++ {
		parent[node] = node
	}
	return &UnionFindSet{
		parent: parent,
		count:  n,
	}
}

func (u *UnionFindSet) Union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return
	}
	u.parent[rootX] = rootY
	u.count--
}

func (u *UnionFindSet) IsConnected(x, y int) bool {
	rootX, rootY := u.find(x), u.find(y)
	return rootX == rootY
}

/*
通过递归方式在查找节点x时，将[根节点 -> x节点]路径上的所有节点，
都挂在根节点下。
使得树的高度维持在常数，操作的时间复杂度降至O(1)。

parent[x]表示x节点指向的父节点，初始化时parent[x] = x，即每个节点自身是独立的联通分量。此时联通分量数量 == 节点数量。

if x == parent[x]，即x的父节点为自己，
则表示x仍是初始化时的独立联通分量，直接返回parent[x]。

if x != parent[x]，即x的父节点是其他节点，表示x已经挂载在其他联通分量下了，
then 递归向上寻找，直到 x == parent[x]，即到达根节点。
此时触底反弹，层层返回这个根节点，在每一层用parent[x]接收，即将每层节点直接指向这个根节点。

注意，第一次写的时候将if写成for，揉杂了递归和迭代。
*/
func (u *UnionFindSet) find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

/*
通过迭代方式，实现跟递归效果相同的路径压缩。

但过程稍有不同：

if x == parent[x]，即x的父节点为自己，
则表示x仍是初始化时的独立联通分量，直接返回parent[x]。

if x != parent[x]，即x的父节点是其他节点，表示x已经挂载在其他联通分量下了，
then 令x的父节点 = x父节点的父节点 -> 此时将x的层级压缩，向上提了一层。
循环此操作，给x赋值[x节点 -> 根节点]路径上的每层节点，将每个x都向上提一层。
直到 x == parent[x]，即到达根节点。
最终返回根节点x。
*/
func (u *UnionFindSet) findIteratively(x int) int {
	for x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UnionFindSet) Size() int {
	return u.count
}
