package array

// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
// 示例 1：
//
// 输入：nums = [3,2,3]
// 输出：[3]
//
// 示例 2：
//
// 输入：nums = [1]
// 输出：[1]
//
// 示例 3：
//
// 输入：nums = [1,2]
// 输出：[1,2]
//
// 提示：
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
func majorityElementIIByHashMap(nums []int) []int {
	cnt := make(map[int]int)
	var res []int
	for _, num := range nums {
		cnt[num]++
	}
	for num, count := range cnt {
		if count > len(nums)/3 {
			res = append(res, num)
		}
	}
	return res
}

/**
这题就不能边计数边获取结果了，参照[2, 2]
*/
