package greedy

// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
// 你需要按照以下要求，给这些孩子分发糖果：
//
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
//
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
//
// 示例 1：
//
// 输入：ratings = [1,0,2]
// 输出：5
// 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
//
// 示例 2：
//
// 输入：ratings = [1,2,2]
// 输出：4
// 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
//
//	第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
//
// 提示：
//
// n == ratings.length
// 1 <= n <= 2 * 10⁴
// 0 <= ratings[i] <= 2 * 10⁴
func candy(ratings []int) int {
	allocation := make([]int, len(ratings))
	for i := 0; i < len(allocation); i++ {
		allocation[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			allocation[i] = allocation[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			allocation[i] = max(allocation[i], allocation[i+1]+1)
		}
	}
	var count int
	for _, num := range allocation {
		count += num
	}
	return count
}

/**
因为需要兼顾左侧和右侧的分配情况，所以需要两轮分配来保证：

先从左往右遍历：处理右侧，即【右孩子分数】>【左孩子分数】的情况；
- 其中，需要先将每个人的糖果分配情况数组初始化，因每人至少有一颗，故全初始化为1；
- 最开始的做法，只初始化了第一个元素为1，后续元素的赋值都依赖于前一个元素——不是局部最优(每个孩子拿尽可能少的糖果)；

再从右往左遍历：处理左侧，即【左孩子分数】>【右孩子分数】的情况；
- 为什么需要反向遍历：如allocation[4]和allocation[5]的比较情况，实际上依赖于allocation[5]和allocation[6]的情况，所以要从右侧开始处理；
- allocation[i] = max(allocation[i], allocation[i+1]+1)：需要取较大值，才能保证左侧和右侧均满足题意。
*/
