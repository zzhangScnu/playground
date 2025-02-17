package greedy

import (
	"math"
	"sort"
)

// 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//
// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//
// 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//
// 以这种方式修改数组后，返回数组 可能的最大和 。
//
// 示例 1：
//
// 输入：nums = [4,2,3], k = 1
// 输出：5
// 解释：选择下标 1 ，nums 变为 [4,-2,3] 。
//
// 示例 2：
//
// 输入：nums = [3,-1,0,2], k = 3
// 输出：6
// 解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
//
// 示例 3：
//
// 输入：nums = [2,-3,-1,5,-4], k = 2
// 输出：13
// 解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
//
// 提示：
//
// 1 <= nums.length <= 10⁴
// -100 <= nums[i] <= 100
// 1 <= k <= 10⁴

type Num []int

func (n Num) Len() int {
	return len(n)
}

func (n Num) Less(i, j int) bool {
	return math.Abs(float64(n[i])) >= math.Abs(float64(n[j]))
}

func (n Num) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Sort(Num(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	if k > 0 && k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	var res int
	for _, num := range nums {
		res += num
	}
	return res
}

/**
局部最优：尽可能将【绝对值大的负数】和【绝对值小的正数】取反；
全局最优：数组和最大。

实现sort.Interface后，可以使用sort.Sort(Num(nums))按自定义规则排序。

这种做法实际上无法更新num的值，因为for range会持有一个固定地址的局部遍历num，在每轮遍历开始的时候，将nums[i]赋值到num上。
for _, num := range nums {
	if num < 0 && k > 0 {
		num *= -1
		k--
	}
}

如果处理完所有负数后，k的余额有剩，则反复操作最小的正数，nums[len(nums)-1]*(-1)*k 。
若k -> 偶数：结果仍等于原值；
若k -> 奇数：结果等于原值*(-1)。
if k > 0 && k%2 == 1 {
	nums[len(nums)-1] *= -1
}
*/
