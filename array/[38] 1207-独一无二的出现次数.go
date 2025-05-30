package array

// 给你一个整数数组 arr，如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
//
// 示例 1：
//
// 输入：arr = [1,2,2,1,1,3]
// 输出：true
// 解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
//
// 示例 2：
//
// 输入：arr = [1,2]
// 输出：false
//
// 示例 3：
//
// 输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
// 输出：true
//
// 提示：
//
// 1 <= arr.length <= 1000
// -1000 <= arr[i] <= 1000
func uniqueOccurrences(arr []int) bool {
	drift := 1000
	counter, occurred := make([]int, 2001), make([]bool, 1001)
	for _, num := range arr {
		counter[num+drift]++
	}
	for _, count := range counter {
		if count == 0 {
			continue
		}
		if !occurred[count] {
			occurred[count] = true
		} else {
			return false
		}
	}
	return true
}

/**
思路：
先遍历数组，统计一遍各元素的出现次数。
由于题目已经给定数据范围，所以可以用数组代替哈希表进行统计。注意数组无法表达负数，所以需要整体向右进行偏移。
对于初始化的长度，需要代入计算一下。
*/
