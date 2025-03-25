package array

// 给定一个非负整数数组 nums， nums 中一半整数是 奇数 ，一半整数是 偶数 。
//
// 对数组进行排序，以便当 nums[i] 为奇数时，i 也是 奇数 ；当 nums[i] 为偶数时， i 也是 偶数 。
//
// 你可以返回 任何满足上述条件的数组作为答案 。
//
// 示例 1：
//
// 输入：nums = [4,2,5,7]
// 输出：[4,5,2,7]
// 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
//
// 示例 2：
//
// 输入：nums = [2,3]
// 输出：[2,3]
//
// 提示：
//
// 2 <= nums.length <= 2 * 10⁴
// nums.length 是偶数
// nums 中一半是偶数
// 0 <= nums[i] <= 1000
//
// 进阶：可以不使用额外空间解决问题吗？
func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	for even, odd := 0, 1; even < n && odd < n; {
		if nums[even]%2 == 0 {
			even += 2
		} else if nums[odd]%2 == 1 {
			odd += 2
		} else {
			nums[even], nums[odd] = nums[odd], nums[even]
			even += 2
			odd += 2
		}
	}
	return nums
}

// 每次只处理一个分支的好处；同时处理带来的混乱：

/**
The user's current code has a loop where they check even and odd separately, then swap if both are wrong. But the issue is that after incrementing even or odd in the first two if statements, the third condition might not trigger when only one is wrong. For example, if even is correct (even index has even number), so even increments, but odd is incorrect. Then, the third condition might not be met because even is already moved. This could leave some elements unswapped.
*/
