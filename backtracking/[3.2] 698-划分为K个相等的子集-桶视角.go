package backtracking

// 给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
// 示例 1：
//
// 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
// 输出： True
// 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//
// 示例 2:
//
// 输入: nums = [1,2,3,4], k = 3
// 输出: false
//
// 提示：
//
// 1 <= k <= len(nums) <= 16
// 0 < nums[i] < 10000
// 每个元素的频率在 [1,4] 范围内
func canPartitionKSubsetsII(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	var used int
	memo := make(map[int]bool)
	var traverse func(nums []int, start int, remainBucketNum int, remainTarget int) bool
	traverse = func(nums []int, start int, remainBucketNum int, remainTarget int) bool {
		if remainBucketNum == 0 {
			return true
		}
		if remainTarget < 0 {
			return false
		}
		if remainTarget == 0 {
			flag := traverse(nums, 0, remainBucketNum-1, target)
			memo[used] = flag
			return flag
		}
		if flag, ok := memo[used]; ok {
			return flag
		}
		for i := start; i < len(nums); i++ {
			if (used>>i)&1 == 1 {
				continue
			}
			used = used | 1<<i
			if traverse(nums, i+1, remainBucketNum, remainTarget-nums[i]) {
				return true
			}
			used = used ^ 1<<i
			if i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
		return false
	}
	return traverse(nums, 0, k, target)
}

/**
以桶的视角：
对于每个桶来说，对每一个数字做选择，每个数字只有2种状态，放入/不放入。
当K个桶都放入了相等的target且数字恰好用完时，表示数组可被划分为K个相等的子集。

仅用for循环实现：
for k > 0 {
	bucket := 0 // 当前桶的和
	for i := 0; i < len(nums); i++ {
		if canAdd(bucket, nums[i]) {
			bucket += nums[i]
		}
		if bucket == target {
			k-- // 装满一个桶，继续下一个
			break
		}
	}
}

改写为递归的实现：
1. 设有K个桶，每个桶的容量为target。维护待塞满的桶的数量，bucketRemain；
2. 维护当前桶剩余多少才能塞满，targetRemain；
3. 维护当前可用数字范围，已经入桶的元素不可重复使用；
4. base case：
	- 若targetRemain为负，则当前分配方式会导致某个桶中的总和超出target；
	- 若targetRemain为零，则当前桶负载符合要求，对下一个桶进行分配；
	- 若bucketRemain为零，且没有命中上述不合法的base case，则找到了一种符合要求的分配方式；
5. 对当前桶尝试进行放入&取出；
6. 如果遍历完所有数字后，依然没有找到符合要求的分配方式，则返回false。

时间复杂度：
设 N = len(nums)，K = 桶的数量
每个桶要遍历N个数字，每个数字有装入/不装入两种选择，组合的结果即每个桶有2^N种选择，则K个桶为：
O(K * 2^N)
*/

/**
性能优化：
1. 可以先对nums排序，优先处理大数，快速失败。
2. 对排序后的nums，相同分支所生成的子树必然重复，可以通过剪枝减小回溯树的大小；
3. 备忘录，如下：

原始做法，纯暴力：
func canPartitionKSubsetsII(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used := make([]bool, len(nums))
	var traverse func(nums []int, remainBucketNum int, remainTarget int) bool
	traverse = func(nums []int, remainBucketNum int, remainTarget int) bool {
		if remainBucketNum == 0 {
			return true
		}
		if remainTarget < 0 {
			return false
		}
		if remainTarget == 0 {
			return traverse(nums, remainBucketNum-1, target)
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			if traverse(nums, remainBucketNum, remainTarget-nums[i]) {
				return true
			}
			used[i] = false
		}
		return false
	}
	return traverse(nums, k, target)
}
会存在大量重叠子问题，重复计算导致效率低下。


优化版本一：备忘录
func canPartitionKSubsetsII(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used := make([]bool, len(nums))
	memo := make(map[string]bool)
	var traverse func(nums []int, remainBucketNum int, remainTarget int) bool
	traverse = func(nums []int, remainBucketNum int, remainTarget int) bool {
		if remainBucketNum == 0 {
			return true
		}
		if remainTarget < 0 {
			return false
		}
		if remainTarget == 0 {
			flag := traverse(nums, remainBucketNum-1, target)
			memo[strings.Join(used, "")] = flag
			return flag
		}
		if flag, ok := memo[strings.Join(used)]; ok {
			return flag
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			if traverse(nums, remainBucketNum, remainTarget-nums[i]) {
				return true
			}
			used[i] = false
		}
		return false
	}
	return traverse(nums, k, target)
}
要注意虽然是桶视角，但是每个元素也只能使用一次，不能重复选取；
所以需要同时控制当前桶和当前可选元素范围，作为函数参数。

仅缓存桶装满时元素的使用情况。
因为桶不满时元素使用分布的情况非常多，有多种路径到达；
而桶装满时元素是相对固定的，枚举记录的成本较低，且具有明确的复用价值。

举例说明：nums = [1, 2, 4, 3, ...], target = 5
第一次尝试：
bucket1 = 1, 4，此时记录used = [1, 0, 1, 0, ...]
bucket2 = 2, 3，此时记录used = [1, 1, 1, 1, ...]
...
bucketN = 无法满足分配要求
此时开始逐层回溯：
1. bucket2使用的2和3释放
2. bucket1使用的1和4释放
回到第一层后，因为nums[0]行不通，所以从nums[1]开始第二次尝试，
bucket1 = 2, 3，此时used = [0, 1, 0, 1, ...]，
bucket2 = 1, 4，此时used = [1, 1, 1, 1, ...]，命中备忘录，直接返回，避免了重复计算。
这种情况下，只是将桶和桶之间装的元素换个位置，肯定是凑不出和为target的K个子集的。不需要再重复穷举，直接返回即可。

为什么仅需关注nums的使用情况，而无需关注具体分配到桶的情况？
因为算法并不关心哪个桶被填满，而只关注剩余的数字能否组成剩余的桶。
如果是一种合法的分配，当完成一个桶即子集时，剩下的数字必须能够组成其他子集，此时used包含了已使用的所有数字。
所以used状态相同 -> 剩下的数字集合相同，转化为更小规模的相同问题，跟此前的桶分配情况无关 -> 剩余所能组成的子集相同

缺点：对used数组反复序列化，会有性能损耗。

优化版本二：位运算备忘录
1 <= k <= len(nums) <= 16，所以用一个整形每一位的0/1就可以表示所有元素的使用情况。
1. 判断第i位是否使用： if (used>>i)&1 == 1
2. 将第i位置1：used = used | 1<<i
   1<<i，将1左移i位，再与used做位或，used的第i位会变成1，其余位保持不变；
3. 将第i位置0：used = used ^ 1<<i
   1<<1，将1左移i位，再与used做位异或(相同为0，不同为1)。
   按位异或通常用于翻转特定位，由于在递归时将第i位置为1，此时翻转为0，即
   used的第i位会变成0，其余位保持不变。
*/
