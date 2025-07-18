package greedy

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
// 0 <= j <= nums[i]
// i + j < n
//
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
//
// 示例 1:
//
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//
//	从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
//
// 示例 2:
//
// 输入: nums = [2,3,0,1,4]
// 输出: 2
//
// 提示:
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
// 题目保证可以到达 nums[n-1]
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step, curCovered, nextCovered := 0, 0, 0
	for i := 0; i <= curCovered; i++ {
		if i+nums[i] > nextCovered {
			nextCovered = i + nums[i]
		}
		if i == curCovered {
			step++
			curCovered = nextCovered
			if curCovered >= len(nums)-1 {
				break
			}
		}
	}
	return step
}

/**
思路：
每个单元格，都有最远覆盖范围，记为curCovered。
遍历该范围，在范围中计算、择优出下一步可到达的最远范围。
循环往复，直到到达终点。

当数组长度为1时，起点即为终点，原地到达。
否则，在当前所覆盖的每个下标上，计算下一步能覆盖的范围，并记录最大值(最远点)，即选择最有潜力以最少步数到达终点的选择。
如果这一步跳完(i == curCovered)，先累加步数，再获取接下来的覆盖范围(curCovered = nextCovered)，
如果到达目的地，就结束处理，并返回最终步数。
*/
