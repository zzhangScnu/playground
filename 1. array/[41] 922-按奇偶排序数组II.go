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

/**
思路：
双指针，odd指向奇数索引位置，even指向偶数索引位置，共同向前推进，每次走2步。
每轮只处理一种情况：
1. 若nums[odd]为奇数，合法 -> 推进odd，进入下一轮循环；
2. 若nums[even]为偶数，合法 -> 推进even，进入下一轮循环；
最终一定会找到一对不合法的odd & even组合，此时：
3. 若nums[odd]为偶数，nums[even]为奇数，不合法 -> 交换nums[odd]和nums[even]，进入下一轮循环。


之前的处理是写了个if & else if & else，会导致遗漏元素。
for even, odd := 0, 1; even < n && odd < n; {
	if nums[even]%2 == 0 {
		even += 2
	}
	if nums[odd]%2 == 1 {
		odd += 2
	}
	if nums[even]%2 == 1 && nums[odd]%2 == 0 {
		nums[even], nums[odd] = nums[odd], nums[even]
		even += 2
		odd += 2
	}
}
问题是，每轮循环可能处理多个场景，导致混乱。
只有在even和odd都不合法时，才会进入交换逻辑。只有一个不合法的场景会被遗漏。
*/
