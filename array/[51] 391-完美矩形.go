package array

import "math"

// 给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是
// (xi, yi) ，右上顶点是 (ai, bi) 。
//
// 如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
// 输出：true
// 解释：5 个矩形一起可以精确地覆盖一个矩形区域。
//
// 示例 2：
//
// 输入：rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
// 输出：false
// 解释：两个矩形之间有间隔，无法覆盖成一个矩形。
//
// 示例 3：
//
// 输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
// 输出：false
// 解释：因为中间有相交区域，虽然形成了矩形，但不是精确覆盖。
//
// 提示：
//
// 1 <= rectangles.length <= 2 * 10⁴
// rectangles[i].length == 4
// -10⁵ <= xi < ai <= 10⁵
// -10⁵ <= yi < bi <= 10⁵
func isRectangleCover(rectangles [][]int) bool {
	X1, Y1, X2, Y2 := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	actualArea := 0
	points := make(map[[2]int]interface{})
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		X1, Y1 = min(X1, x1), min(Y1, y1)
		X2, Y2 = max(X2, x2), max(Y2, y2)
		actualArea += (x2 - x1) * (y2 - y1)
		p1, p2, p3, p4 := [2]int{x1, y2}, [2]int{x2, y2}, [2]int{x1, y1}, [2]int{x2, y1}
		for _, point := range [][2]int{p1, p2, p3, p4} {
			if _, ok := points[point]; ok {
				delete(points, point)
			} else {
				points[point] = struct{}{}
			}
		}
	}
	area := (X2 - X1) * (Y2 - Y1)
	if area != actualArea {
		return false
	}
	if len(points) != 4 {
		return false
	}
	if _, ok := points[[2]int{X1, Y2}]; !ok {
		return false
	}
	if _, ok := points[[2]int{X2, Y2}]; !ok {
		return false
	}
	if _, ok := points[[2]int{X1, Y1}]; !ok {
		return false
	}
	if _, ok := points[[2]int{X2, Y1}]; !ok {
		return false
	}
	return true
}
