package greedy

import (
	"container/list"
	"sort"
)

// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第 i
// 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
//
// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj] 是队列中第
// j 个人的属性（queue[0] 是排在队列前面的人）。
//
// 示例 1：
//
// 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
// 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
// 解释：
// 编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
// 编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
// 编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
// 编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
// 编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
// 编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
// 因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
//
// 示例 2：
//
// 输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
// 输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
//
// 提示：
//
// 1 <= people.length <= 2000
// 0 <= hi <= 10⁶
// 0 <= ki < people.length
// 题目数据确保队列可以被重建
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	head := list.New()
	head.PushBack([]int{})
	for i := 0; i < len(people); i++ {
		node := head.PushBack(people[i])
		position := people[i][1]
		cur := head.Front()
		for position > 0 {
			position--
			cur = cur.Next()
		}
		head.MoveAfter(node, cur)
	}
	var res [][]int
	for cur := head.Front().Next(); cur != nil; cur = cur.Next() {
		res = append(res, cur.Value.([]int))
	}
	return res
}

/**
思路跟分发糖果一样，当有两个维度时，先固定下来一个，处理完成后再考虑另一个。
本题中，
- 先固定相对顺序：因为身高未知，所以明确不下来；
- 先固定身高：在身高有序的情况下，再根据相对顺序进行调整，从而使序列满足题意。
  因为相对顺序的定义是"前面正好有ki个身高>=hi的人"，所以身高按从高到低降序先排序。身高相同的情况下，再按相对顺序从少到多排。
固定身高后，再根据相对顺序从头开始插队。又因为身高已经降序排下来了，往前插队一定能满足"前面都是比自己高的人"，也不会影响先于自己完成插队的节点。

如果用数组去实现插队，每次都要往后移动元素再插入到正确位置，时间复杂度高，这种场景应该考虑使用链表。

代码注意点：
1. sort.Slice的应用；
2. container.list的应用；
3. 虚拟头节点再次登场，head.PushBack([]int{}) 和 for cur := head.Front().Next(); cur != nil; cur = cur.Next()
   保证兼容直接插入最前面和插入中间位置的场景。
*/
