package diffarray

// 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
//
// 给定整数 capacity 和一个数组 trips , trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有
// numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
//
// 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
//
// 示例 1：
//
// 输入：trips = [[2,1,5],[3,3,7]], capacity = 4
// 输出：false
//
// 示例 2：
//
// 输入：trips = [[2,1,5],[3,3,7]], capacity = 5
// 输出：true
//
// 提示：
//
// 1 <= trips.length <= 1000
// trips[i].length == 3
// 1 <= numPassengersi <= 100
// 0 <= fromi < toi <= 1000
// 1 <= capacity <= 10⁵
func carPooling(trips [][]int, capacity int) bool {
	tripsLen := 1001
	difference := NewDifference(tripsLen)
	for _, trip := range trips {
		difference.Update(trip[1], trip[2]-1, trip[0])
	}
	sizes := difference.GetResult()
	for _, size := range sizes {
		if size > capacity {
			return false
		}
	}
	return true
}

/**
思路：
利用差分数组，将每段路程期望携带的乘客数进行累加，最终通过差分数组还原出每个站点期望携带的乘客数。
再判断每个站点是否超员，即超出空座位大小capacity。

注意，因为trip[i] = [numPassengersi, fromi, toi]，toi表示乘客的下车站点，即numPassengersi在数组索引fromi增加，在toi减少。
即numPassengersi影响的数组索引区间为[fromi, toi)。
*/
