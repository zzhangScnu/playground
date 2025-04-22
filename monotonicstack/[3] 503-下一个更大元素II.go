package monotonicstack

// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素
// 。
//
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
// 。
//
// 示例 1:
//
// 输入: nums = [1,2,1]
// 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数；
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
// 示例 2:
//
// 输入: nums = [1,2,3,4,3]
// 输出: [2,3,4,-1,4]
//
// 提示:
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
func nextGreaterElements(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = -1
	}
	st := []int{0}
	for i := 1; i < size*2; i++ {
		if nums[i%size] <= nums[st[len(st)-1]] {
			st = append(st, i%size)
		} else {
			for len(st) > 0 && nums[i%size] > nums[st[len(st)-1]] {
				res[st[len(st)-1]] = nums[i%size]
				st = st[:len(st)-1]
			}
			st = append(st, i%size)
		}
	}
	return res
}

/**
思路：

第一种方法：
将nums拼接在nums后面，将环形问题拆解为线性问题。
记录结果的数组随之扩大2倍，但返回时只取前面一半的值。


第二种方法：
不改变nums，直接使用取模的方式模拟回环。
记录结果的数组跟nums的大小相同。
只要遇到环形数组，都可以用取模思路解决。

第二种方法中，res的大小为nums的大小，而栈会顺序装下2倍nums的元素。
但res中存放的结果并不会被覆盖。
*/
