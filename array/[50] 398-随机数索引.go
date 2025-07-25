package array

import (
	"math/rand"
)

// 给你一个可能含有 重复元素 的整数数组 nums ，请你随机输出给定的目标数字 target 的索引。你可以假设给定的数字一定存在于数组中。
//
// 实现 Solution 类：
//
// Solution(int[] nums) 用数组 nums 初始化对象。
// int pick(int target) 从 nums 中选出一个满足 nums[i] == target 的随机索引 i 。如果存在多个有效的索引，则每
// 个索引的返回概率应当相等。
//
// 示例：
//
// 输入
// ["Solution", "pick", "pick", "pick"]
// [[[1, 2, 3, 3, 3]], [3], [1], [3]]
// 输出
// [null, 4, 0, 2]
//
// 解释
// Solution solution = new Solution([1, 2, 3, 3, 3]);
// solution.pick(3); // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
// solution.pick(1); // 返回 0 。因为只有 nums[0] 等于 1 。
// solution.pick(3); // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
//
// 提示：
//
// 1 <= nums.length <= 2 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// target 是 nums 中的一个整数
// 最多调用 pick 函数 10⁴ 次

type SolutionReservoir struct {
	nums []int
}

func ConstructorReservoir(nums []int) SolutionReservoir {
	return SolutionReservoir{nums: nums}
}

func (this *SolutionReservoir) PickFromReservoir(target int) int {
	count, res := 0, 0
	for index, num := range this.nums {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				res = index
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

/**
水塘算法：
对于一个不知道长度的序列，需要等概率获取其中的元素。
维护一个大小为k的水塘，将序列中的元素逐个加入。
通过随机抽取+替换，使得每次获取的元素概率都相等。

对于本题：
在序列中获取元素target时，需遍历一次序列，通过O(n)的时间复杂度等概率抽取一个元素。
维护两个辅助变量：
- count：迄今为止遍历到的等于target的元素数量，即水塘；
- res：目前抽取的等于target的元素索引；
步骤：
- 遍历序列，若此时 num == target：
	- count++；
	- 在[0, count)之中生成一个随机数，如果该随机数 == 0，则用当前元素的索引作为结果：
		- 如果当前水塘中元素个数为1，rand.Intn(1) == 0，即抽取该元素的概率 == 100%，res = 0；
		- 如果当前再次加入一个1，rand.Intn(2) IN {0, 1}，如果此时为0，则替换res = 1，使得两个1的抽取概率均为50%。
		- 如果当前再次加入一个1，则要根据重新抛硬币决定是保留第二轮的抽取，还是替换为本轮新加入的元素；
		- 以此类推……
*/
