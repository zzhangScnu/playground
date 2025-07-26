package array

import "math"

// 集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有
// 一个数字重复 。
//
// 给定一个数组 nums 代表了集合 S 发生错误后的结果。
//
// 请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
//
// 示例 1：
//
// 输入：nums = [1,2,2,4]
// 输出：[2,3]
//
// 示例 2：
//
// 输入：nums = [1,1]
// 输出：[1,2]
//
// 提示：
//
// 2 <= nums.length <= 10⁴
// 1 <= nums[i] <= 10⁴
func findErrorNums(nums []int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}
	res := make([]int, 2)
	for num := 1; num <= len(nums); num++ {
		count := counter[num]
		if count == 2 {
			res[0] = num
		} else if count == 0 {
			res[1] = num
		}
	}
	return res
}

func findErrorNumsII(nums []int) []int {
	res := make([]int, 2)
	for _, num := range nums {
		index := int(math.Abs(float64(num))) - 1
		if nums[index] < 0 {
			res[0] = index + 1
		}
		nums[index] = -nums[index]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && i+1 != res[0] {
			res[1] = i + 1
		}
	}
	return res
}

func findErrorNumsIII(nums []int) []int {
	res := make([]int, 2)
	for _, num := range nums {
		index := int(math.Abs(float64(num))) - 1
		if nums[index] < 0 {
			res[0] = index + 1
		} else {
			nums[index] = -nums[index]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res[1] = i + 1
		}
	}
	return res
}
