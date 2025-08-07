package hashmap

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1：
//
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
// 示例 3：
//
// 输入：nums = [1,0,1,2]
// 输出：3
//
// 提示：
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
func longestConsecutive(nums []int) int {
	mapping := make(map[int]interface{})
	for _, num := range nums {
		mapping[num] = struct{}{}
	}
	var res int
	for num := range mapping {
		if _, ok := mapping[num-1]; ok {
			continue
		}
		curNum, maxLen, flag := num, 1, false
		for {
			if _, flag = mapping[curNum+1]; flag {
				maxLen += 1
				curNum += 1
			}
			if !flag {
				break
			}
		}
		res = max(res, maxLen)
	}
	return res
}

/**
思路：
维护一个HashMap，用途有二：
1. 对原数组去重；
2. 建立【元素 -> 是否存在】的快速映射，用O(n)空间换O(1)时间，避免每次都需要O(n)遍历数组。

优化：
在遍历到元素 num 时，先判断 num - 1 是否存在。
若存在，则 num 一定不是最长连续序列的起点，跳过即可。
一定会在 num - 1 或更先前的元素开始，作为最长连续序列的起点。
有了这个优化，虽然代码中有双重 for 循环，但每个元素都只会被遍历一次。
*/
