package array

import "math"

// 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
//
// 示例 1：
//
// 输入：points = [[1,1],[2,2],[3,3]]
// 输出：3
//
// 示例 2：
//
// 输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// 输出：4
//
// 提示：
//
// 1 <= points.length <= 300
// points[i].length == 2
// -10⁴ <= xi, yi <= 10⁴
// points 中的所有点 互不相同
func maxPoints(points [][]int) int {
	res := 1
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		count := make(map[float64]int)
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			slope := math.MaxFloat64
			if p1[0] != p2[0] {
				slope = float64(p1[1]-p2[1]) / float64(p1[0]-p2[0])
			}
			count[slope]++
			res = max(res, count[slope]+1)
		}
	}
	return res
}
