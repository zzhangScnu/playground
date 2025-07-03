package dynamicprogramming

func robII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	res1 := robRange(nums, 0, n-2)
	res2 := robRange(nums, 1, n-1)
	return max(res1, res2)
}

func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start], dp[start+1] = nums[start], max(nums[start], nums[start+1])
	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[end]
}

/**
本题相当于成环，环相接的地方就是相邻的房屋，不能连续偷窃。
所以问题可以转化为，
【考虑偷第一间 -> 如果偷第一间，则最后一间不能偷；如果不偷第一间，则最后一间能偷】
【考虑偷最后一间 -> 如果偷最后一间，则第一间不能偷；如果不偷最后一间，则第一间能偷】
这两种情况的最小值。
为什么无需考虑只偷中间，不偷首尾的情况？因为已经被上两种场景包含了。
且这两种情况选择的范围比只偷中间大，均为n-1，而只偷中间仅为n-2，房子里都是钱，所以取最大值时考虑两边一定会比只考虑中间优。

在robRange方法中，因为有start+1索引取值的操作，所以需要保证end > start才能往下走；
其次因为限制了本次处理的是nums数组的[start, end]区间，所以虽然dp初始化的长度为len(nums)，但是在其上的初始化、赋值、返回值等操作都是跟随start和end的。
*/
