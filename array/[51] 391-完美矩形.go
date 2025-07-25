package array

// rectangles，⾥⾯装着若⼲四元组
// (x1,y1,x2,y2)，每个四元组就是记录⼀个矩形的左下⻆和右上⻆坐标。
func isRectangleCover(rectangles [][]int) bool {
	X1, Y1, X2, Y2 := 0, 0, 0, 0
	actualArea := 0
	points := make(map[int]interface{})
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[2]
		X1, Y1 = min(X1, x1), max(Y1, y1)
		X2, Y2 = max(X2, x2), min(Y2, y2)
		actualArea += (x2 - x1) * (y2 - y1)
		for _, point := range []int{x1, x2, y1, y2} {
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
	for _, point := range points {
		if point != X1 && point != X2 && point != Y1 && point != Y2 {
			return false
		}
	}
	return true
}
