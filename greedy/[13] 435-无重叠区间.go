package greedy

import "sort"

// 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重
// 叠 。
//
// 注意 只在一点上接触的区间是 不重叠的。例如 [1, 2] 和 [2, 3] 是不重叠的。
//
// 示例 1:
//
// 输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
// 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。
//
// 示例 2:
//
// 输入: intervals = [ [1,2], [1,2], [1,2] ]
// 输出: 2
// 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//
// 示例 3:
//
// 输入: intervals = [ [1,2], [2,3] ]
// 输出: 0
// 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
// 提示:
//
// 1 <= intervals.length <= 10⁵
// intervals[i].length == 2
// -5 * 10⁴ <= starti < endi <= 5 * 10⁴
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var count int
	preEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if preEnd <= intervals[i][0] {
			preEnd = intervals[i][1]
		} else {
			count++
			preEnd = min(preEnd, intervals[i][1])
		}
	}
	return count
}

/**
跟引爆气球思路相似。
先按左区间升序排序。
固定一段区间，跟下一段区间判断是否重叠。
如果不重叠，就固定下一段区间，继续向后判断；
如果重叠，就移除两个区间中的较长者，保留较短者。贪心就体现在这里，给后续的区间多留空间。

不需要先排左区间、再排右区间，因为如果区间重叠(甚至起点相等)的话，也会取终点较小者保留。
*/
