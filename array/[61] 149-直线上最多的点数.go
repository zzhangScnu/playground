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

/**
思路：


前置知识：
斜率 = (y1 - y0) / (x1 - x0)
在同一条直线上 -> 斜率相等


解法：
哈希表：【斜率 -> 该直线上的点数】

固定每个点作为线段起点，与其他点共同计算斜率，并收集至哈希表中。
即对每一个点，需重新初始化用于计数的哈希表。
再维护一个全局变量res，收集全局最大的直线上的点的个数。

注意，如果x0 == x1，即线段垂直于x轴，则斜率计算公式的被除数会为0，导致结果异常。此时赋一个默认值即可规避。

结果为int：slope = (p1[0] - p2[0]) / (p1[1] - p2[1])
结果为float64：slope = float64(p1[1]-p2[1]) / float64(p1[0]-p2[0])
*/
