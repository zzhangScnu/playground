package array

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]
//排序后，数组变为 [0,1,9,16,100]
//
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]

// 提示：
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 已按 非递减顺序 排序
//
// 进阶：
//
// 请你设计时间复杂度为 O(n) 的算法解决本问题
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	low, high := 0, len(nums)-1
	i := len(nums) - 1
	for low <= high {
		if nums[low]*nums[low] > nums[high]*nums[high] {
			res[i] = nums[low] * nums[low]
			low++
		} else {
			res[i] = nums[high] * nums[high]
			high--
		}
		i--
	}
	return res
}

/**
- 一开始的思路是从中间到两边——先找到第一个正数，将其作为分隔，双指针向两边扩展。但这样的问题是需要处理数组越界的边界场景；
- 后来的思路是以结果数组的长度作为遍历条件。但涉及到双指针的题解，最好统一以两个指针的相对位置作为控制条件；
- make([]int, 0, len(nums))是创建一个长度为0的数组，直接下标访问会越界。正确写法是make([]int, len(nums))
*/
