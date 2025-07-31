package greedy

import "sort"

// 你将会获得一系列视频片段，这些片段来自于一项持续时长为 time 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
//
// 使用数组 clips 描述所有的视频片段，其中 clips[i] = [starti, endi] 表示：某个视频片段开始于 starti 并于
// endi 结束。
//
// 甚至可以对这些片段自由地再剪辑：
//
// 例如，片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。
//
// 我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, time]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1
// 。
//
// 示例 1：
//
// 输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], time = 10
// 输出：3
// 解释：
// 选中 [0,2], [8,10], [1,9] 这三个片段。
// 然后，按下面的方案重制比赛片段：
// 将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
// 现在手上的片段为 [0,2] + [2,8] + [8,10]，而这些覆盖了整场比赛 [0, 10]。
//
// 示例 2：
//
// 输入：clips = [[0,1],[1,2]], time = 5
// 输出：-1
// 解释：
// 无法只用 [0,1] 和 [1,2] 覆盖 [0,5] 的整个过程。
//
// 示例 3：
//
// 输入：clips = [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],
// [2,6],[3,4],[4,5],[5,7],[6,9]], time = 9
// 输出：3
// 解释：
// 选取片段 [0,4], [4,7] 和 [6,9] 。
//
// 提示：
//
// 1 <= clips.length <= 100
// 0 <= starti <= endi <= 100
// 1 <= time <= 100
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	if clips[0][0] > 0 {
		return -1
	}
	count, curEnd, nextEnd, i := 0, 0, 0, 0
	for i < len(clips) && clips[i][0] <= curEnd {
		for i < len(clips) && clips[i][0] <= curEnd {
			nextEnd = max(nextEnd, clips[i][1])
			i++
		}
		count++
		if nextEnd >= time {
			return count
		}
		curEnd = nextEnd
	}
	return -1
}

/**
思路：
遇到区间问题，第一反应就是先排序。
一般是先根据起点排序，起点相同的再根据终点排序。

本题需要寻找几个片段，可以完整覆盖从起点到终点即[0, time]区间。
所以使用贪心思路：
1. 先锚定一个区间终点curEnd，寻找所有起点 < curEnd的片段。目的是为了无缝衔接，连成更长的片段；
2. 在这些符合条件的片段中，使用贪心思想，寻找终点最远的片段，使得覆盖长度尽可能延伸，从而达到片段数量尽可能少的目标；
3. 在寻找完成之后，相当于选择了衔接的片段，此时count++且curEnd = nextEnd；如果此时已经到达终点，则返回片段数量。
*/

/**
sort.Slice(clips, func(i, j int) bool {
	if clips[i][0] == clips[j][0] {
		return clips[i][1] > clips[j][1]
	}
	return clips[i][0] < clips[j][0]
})

sort.Slice(clips, func(i, j int) bool {
	return clips[i][0] == clips[j][0] && clips[i][1] > clips[j][1] ||
		clips[i][0] < clips[j][0]
})

代码逻辑等价。
*/
