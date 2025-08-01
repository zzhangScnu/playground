package array

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
// 说明：
//
// 为什么返回数值是整数，但输出的答案是数组呢？
//
// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
// int len = removeDuplicates(nums);
//
// // 在函数里修改输入数组对于调用者是可见的。
// // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
//
//	for (int i = 0; i < len; i++) {
//	 print(nums[i]);
//	}
//
// 示例 1：
//
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
//
// 示例 2：
//
// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
// 解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素
// 。
//
// 提示：
//
// 1 <= nums.length <= 3 * 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 已按升序排列
func removeDuplicatesII(nums []int) int {
	slow, fast, count := 1, 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 1
		}
		if count <= 2 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/**
思路：
元素最多只能出现一次 -> 通过比较 nums[slow] 与 nums[fast] -> 不相等则将 nums[fast] 赋值到 slow 位置并推进 slow
元素最多只能出现两次 -> 通过比较 nums[fast - 1] 与 nums[fast] -> 不相等则重置重复元素出现次数count / 相等则累加count -> 若count未超限则将 nums[fast - 1] 赋值到 slow 位置并推进 slow

相同点是：每轮fast指针都需要移动；而slow指针仅在结果数组发生变更时移动。
*/

func removeDuplicatesIII(nums []int) int {
	return doRemoveDuplicates(nums, 2)
}

func doRemoveDuplicates(nums []int, maxDuplicateTimes int) int {
	if len(nums) < maxDuplicateTimes {
		return len(nums)
	}
	slow, fast := maxDuplicateTimes, maxDuplicateTimes
	for ; fast < len(nums); fast++ {
		if nums[slow-maxDuplicateTimes] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/**
思路：
可以将问题扩展为：最多重复 maxDuplicateTimes 次(k次)。前提是数组有序。
对于最多可重复 k 次的场景，slow 指针指向的是下一个合法&可被接纳的结果元素，slow 指针前的[slow - k, slow - 1]区间最多包含 k 个相同元素。
- 如果此时 nums[slow - k] == nums[fast]，说明如果要接纳 nums[fast]，则当前重复元素会超出 k 次限制。此时不能接纳 nums[fast]，所以仅需要推进 fast 指针，进行下一轮判断；
- 否则可以接纳 nums[fast]，需要同时推进 fast 和 slow 指针。
*/
