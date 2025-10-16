package binarysearch

import "slices"

// 给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[
// j] 表示第 j 瓶药水的能量强度。
//
// 同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。
//
// 请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。
//
// 示例 1：
//
// 输入：spells = [5,1,3], potions = [1,2,3,4,5], success = 7
// 输出：[4,0,3]
// 解释：
// - 第 0 个咒语：5 * [1,2,3,4,5] = [5,10,15,20,25] 。总共 4 个成功组合。
// - 第 1 个咒语：1 * [1,2,3,4,5] = [1,2,3,4,5] 。总共 0 个成功组合。
// - 第 2 个咒语：3 * [1,2,3,4,5] = [3,6,9,12,15] 。总共 3 个成功组合。
// 所以返回 [4,0,3] 。
//
// 示例 2：
//
// 输入：spells = [3,1,2], potions = [8,5,8], success = 16
// 输出：[2,0,2]
// 解释：
// - 第 0 个咒语：3 * [8,5,8] = [24,15,24] 。总共 2 个成功组合。
// - 第 1 个咒语：1 * [8,5,8] = [8,5,8] 。总共 0 个成功组合。
// - 第 2 个咒语：2 * [8,5,8] = [16,10,16] 。总共 2 个成功组合。
// 所以返回 [2,0,2] 。
//
// 提示：
//
// n == spells.length
// m == potions.length
// 1 <= n, m <= 10⁵
// 1 <= spells[i], potions[i] <= 10⁵
// 1 <= success <= 10¹⁰
func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	n, m := len(spells), len(potions)
	res := make([]int, n)
	var find func(spell int) int
	find = func(spell int) int {
		left, right, mid := 0, m-1, 0
		for left <= right {
			mid = left + (right-left)>>1
			if int64(spell)*int64(potions[mid]) >= success {
				if mid == 0 || int64(spell)*int64(potions[mid-1]) < success {
					return m - mid
				}
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return 0
	}
	for i := 0; i < n; i++ {
		res[i] = find(spells[i])
	}
	return res
}

/**
思路：
先对potions进行排序，再在其上进行二分搜索。
本质上是找到potions的左边界，令其后面的potions与当前spell相乘均 >= success。

注意，这里涉及int和int64的运算和比较，需要先将操作数转换为int64，再进行相乘，
而不是相乘后再转为int64，这样会有计算结果溢出的风险。
*/
