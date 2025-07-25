package array

import "math"

// rectangles，⾥⾯装着若⼲四元组
// (x1,y1,x2,y2)，每个四元组就是记录⼀个矩形的左下⻆和右上⻆坐标。
func isRectangleCover(rectangles [][]int) bool {
	X1, Y1, X2, Y2 := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	actualArea := 0
	points := make(map[[2]int]interface{})
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[2]
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
