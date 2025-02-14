package array

// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
// 示例 1：
//
// 输入：nums = [3,2,3]
// 输出：[3]
//
// 示例 2：
//
// 输入：nums = [1]
// 输出：[1]
//
// 示例 3：
//
// 输入：nums = [1,2]
// 输出：[1,2]
//
// 提示：
//
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
func majorityElementIIByHashMap(nums []int) []int {
	cnt := make(map[int]int)
	var res []int
	for _, num := range nums {
		cnt[num]++
	}
	for num, count := range cnt {
		if count > len(nums)/3 {
			res = append(res, num)
		}
	}
	return res
}

func majorityElementII(nums []int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
		if len(cnt) <= 2 {
			continue
		}
		for k := range cnt {
			cnt[k]--
			if cnt[k] == 0 {
				delete(cnt, k)
			}
		}
	}
	var res []int
	for k := range cnt {
		var count int
		for _, num := range nums {
			if num == k {
				count++
			}
		}
		if count > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}

/**
思路1：
使用map记录每个元素出现的频率。这题就不能边计数边获取结果了，参照[2, 2]，会得到[2, 2]作为结果。

思路2：
因为出现次数要求超过[n/3]，对于一个数组来说，这样的数最多只会有2个。使用容量为2的map来存放元素和出现频率。
当map长度大于容量时，将map中所有元素的出现频率都减去1，同时移除为0的元素，因为它们已经不符合条件了。
最后map中剩余的元素，也不一定满足条件，如原始数组[1, 1, 2, 2, 3, 3, 4]，最后剩下的4不能作为结果返回。
所以还要整体遍历一遍原始数组，统计一下实际出现的次数。
整体时间复杂度O(n)，空间复杂度O(1)。
*/
