package unionfindset

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
// 示例 3：
//
// 输入：nums = [1,0,1,2]
// 输出：3
//
// 提示：
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹

type UnionFindSet128 struct {
	parent map[int]int
	size   map[int]int
}

func NewUnionFindSet128(nums map[int]interface{}) *UnionFindSet128 {
	parent, size := make(map[int]int, len(nums)), make(map[int]int, len(nums))
	for num := range nums {
		parent[num] = num
		size[num] = 1
	}
	return &UnionFindSet128{
		parent: parent,
		size:   size,
	}
}

func (u *UnionFindSet128) isConnected(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *UnionFindSet128) find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UnionFindSet128) connect(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return
	}
	u.parent[rootX] = rootY
	u.size[rootY] += u.size[rootX]
}

func longestConsecutive(nums []int) int {
	set := make(map[int]interface{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	unionFindSet := NewUnionFindSet128(set)
	for num := range set {
		if _, ok := set[num+1]; ok {
			unionFindSet.connect(num, num+1)
		}
	}
	var maxLen int
	for _, size := range unionFindSet.size {
		maxLen = max(maxLen, size)
	}
	return maxLen
}
