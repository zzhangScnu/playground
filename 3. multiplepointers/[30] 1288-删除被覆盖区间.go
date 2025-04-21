package multiplepointers

import "sort"

// 给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
//
// 只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
//
// 在完成所有删除操作后，请你返回列表中剩余区间的数目。
//
// 示例：
//
// 输入：intervals = [[1,4],[3,6],[2,8]]
// 输出：2
// 解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了。
//
// 提示：
//
// 1 <= intervals.length <= 1000
// 0 <= intervals[i][0] < intervals[i][1] <= 10^5
// 对于所有的 i != j：intervals[i] != intervals[j]
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] ||
			intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1]
	})
	covered, curEnd := 0, -1
	for _, interval := range intervals {
		if interval[1] <= curEnd {
			covered++
		} else {
			curEnd = interval[1]
		}
	}
	return len(intervals) - covered
}

/**
思路：
先对区间列表进行排序。
首先对起点升序排序；
其次起点相同的情况下，对终点降序排序。
——因为本题需要删除被覆盖区间，所以这样排序之后，对于起点相同的区间，一定是较长的区间靠前，
则可以用它的终点作为边界条件curEnd，来判断【起点相同】的区间是否被覆盖；
且随着区间的遍历，滚动更新curEnd。

若区间未被覆盖，即当前区间的interval[1] > curEnd，则仅更新curEnd。

最终用区间列表长度减去被覆盖的区间列表长度，则为未被覆盖的区间列表长度。
*/

/**
为什么不能用双指针法？
因为双指针处理区间问题，多用于两个列表的相对位置关系，需要同步推进；
而本题求的是同一个列表的相互包含关系。
*/
