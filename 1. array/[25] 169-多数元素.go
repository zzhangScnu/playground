package array

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
// 示例 1：
//
// 输入：nums = [3,2,3]
// 输出：3
//
// 示例 2：
//
// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
//
// 提示：
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
func majorityElementByHashMap(nums []int) int {
	res, maxCnt := 0, 0
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
		if cnt[num] > maxCnt {
			res = num
			maxCnt = cnt[num]
		}
	}
	return res
}

func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			res = num
		}
		if num == res {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
