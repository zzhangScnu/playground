package N数之和

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
// 你可以按任意顺序返回答案。
//
// 示例 1：
//
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
// 示例 2：
//
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
//
// 示例 3：
//
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
//
// 提示：
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// 只会存在一个有效答案
func twoSum(nums []int, target int) []int {
	numIdxMap := make(map[int]*int, len(nums))
	for i, num := range nums {
		idx := i
		if numIdxMap[target-num] != nil {
			return []int{*numIdxMap[target-num], idx}
		}
		numIdxMap[num] = &idx
	}
	return []int{}
}

/**
做多少次都永远无法一次过的第一题……
这次掉的坑，是for里面的i和num，实际上是固定内存地址的。
如果不做额外赋值 idx := i，而是直接做 numIdxMap[num] = &i，
会导致最终numIdxMap里面的value全都是同一个值，而且是for最后循环到的那个值。

其实更好的写法，我最开始想到了，但没有实现，就是直接用go map的ok特性。

注意要先判断是否有和等于target的情况，再将本次循环到的数塞进map，否则会导致我等于我自己-。-
*/
