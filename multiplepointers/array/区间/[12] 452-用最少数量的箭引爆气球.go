package 区间

import "sort"

// 有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示
// 水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
//
// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足
// xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
//
// 给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。
//
// 示例 1：
//
// 输入：points = [[10,16],[2,8],[1,6],[7,12]]
// 输出：2
// 解释：气球可以用2支箭来爆破:
// -在x = 6处射出箭，击破气球[2,8]和[1,6]。
// -在x = 11处发射箭，击破气球[10,16]和[7,12]。
//
// 示例 2：
//
// 输入：points = [[1,2],[3,4],[5,6],[7,8]]
// 输出：4
// 解释：每个气球需要射出一支箭，总共需要4支箭。
//
// 示例 3：
//
// 输入：points = [[1,2],[2,3],[3,4],[4,5]]
// 输出：2
// 解释：气球可以用2支箭来爆破:
// - 在x = 2处发射箭，击破气球[1,2]和[2,3]。
// - 在x = 4处射出箭，击破气球[3,4]和[4,5]。
//
// 提示:
//
// 1 <= points.length <= 10⁵
// points[i].length == 2
// -2³¹ <= xstart < xend <= 2³¹ - 1
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	count := 1
	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			count++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	return count
}

/**
思路：
先按气球的起始位置进行排序。
再不断往后判断，设：
x = 可以被一箭射爆的一系列气球[0, i-1]的右边界的最小值；
y = 下一个气球i的左边界的值。
当x >= y时，意味着[0, i]区间内的所有气球都是有重叠的，可以被一箭带走；否则，所需箭数需要累加。

其实将区间气球的右边界最小值更新为i-1个气球的右边界太tricky，可以用一个变量来维护会更直观。
func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })
    res := 1
    right := points[0][1]
    for _, point := range points[1:] {
        if point[0] > right {
            res++
            right = point[1]
        } else {
            right = min(right, point[1])
        }
    }
    return res
}
*/
