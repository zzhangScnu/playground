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

/**
思路
思路一：
O(n)时间复杂度 & O(n)空间复杂度
1. 第一次遍历：用HashMap维护nums中元素的出现次数
2. 第二次遍历：对HashMap中的元素进行判断
	- 出现次数为2 -> 重复的元素
	- 出现次数为0 -> 缺失的元素
	- 出现此时为1 -> 正常的元素

思路二：
O(n)时间复杂度 & O(1)空间复杂度
索引范围[0, n - 1]
元素范围[1, n]
正常情况下，元素和索引是一一匹配的，元素nums[i] <-> 索引i + 1。
但因为存在某个元素重复，导致2个不同索引指向同一个元素。所以可以利用这个特性，将重复元素挑选出来；然后再基于此将缺失元素也挑选出来。
步骤：
遍历数组，在每一轮处理中：
1. 将元素nums[i]映射为索引index，即index = nums[i] - 1；
2. 将index位置的元素设置为负数，即nums[index] = -nums[index]；
3. 如果在遍历中遇到已经是负数的nums[index]，说明是重复的元素，将其加入结果集。
再次遍历数组，寻找一个正数nums[i]：
- 如果nums[i]不等于结果集中重复的元素，则为缺失的元素，将【索引i对应的缺失元素i + 1】加入结果集——这种对应的是第一次遍历中无脑取负的处理方式——即第一次遍历结束后，数组中还有两个正数；
- 直接将【索引i对应的缺失元素i + 1】加入结果集——这种对应的是第一次遍历中，如果遍历到负数，就不再取负的处理方式——即第一次遍历结束后，数组中只剩一个正数。

注意：
- 因为会将元素映射为索引，而元素会被置为负数，所以需要取绝对值避免数组越界问题。
*/
