package binarysearch

// 传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
//
// 传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最
// 大运载重量。
//
// 返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。
//
// 示例 1：
//
// 输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
// 输出：15
// 解释：
// 船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
// 第 1 天：1, 2, 3, 4, 5
// 第 2 天：6, 7
// 第 3 天：8
// 第 4 天：9
// 第 5 天：10
//
// 请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (1
// 0) 是不允许的。
//
// 示例 2：
//
// 输入：weights = [3,2,2,4,1,4], days = 3
// 输出：6
// 解释：
// 船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
// 第 1 天：3, 2
// 第 2 天：2, 4
// 第 3 天：1, 4
//
// 示例 3：
//
// 输入：weights = [1,2,3,1,1], days = 4
// 输出：3
// 解释：
// 第 1 天：1
// 第 2 天：2
// 第 3 天：3
// 第 4 天：1, 1
//
// 提示：
//
// 1 <= days <= weights.length <= 5 * 10⁴
// 1 <= weights[i] <= 500
func shipWithinDays(weights []int, days int) int {
	var left, right int
	for _, weight := range weights {
		if weight > left {
			left = weight
		}
		right += weight
	}
	for left <= right {
		mid := left + (right-left)/2
		d := day(weights, mid)
		if d == days {
			right = mid - 1
		} else if d < days {
			right = mid - 1
		} else if d > days {
			left = mid + 1
		}
	}
	return left
}

func day(weights []int, capability int) int {
	var d int
	for i := 0; i < len(weights); {
		remainCapability := capability
		for i < len(weights) {
			if remainCapability >= weights[i] {
				remainCapability -= weights[i]
				i++
			} else {
				break
			}
		}
		d++
	}
	return d
}

/**
关于day的计算：
跟875-珂珂吃香蕉有不同。
875：每次只处理一个数据；
1011：只要未超限，每次可处理多个数据。

本题f(x)的计算，更像是双指针的思想：
- 外层指针指向本次处理的起始点；
- 内层指针在限制范围内向前推进，直到无法再推进。

三部曲：
1. 画出函数在二维坐标上的图像，明确 x、f(x)、target，并实现函数 f；
	- x：所求的船只运载能力；
	- f(x)：将传送带上的所有包裹送达的天数；运载能力↑ -> 所需天数↓，存在单调递减关系。
	- target：days。


2. 明确 x 的取值范围，作为二分搜索的搜索区间，初始化left和right变量；
	- left：最大货物的重量，即每次仅能运走一件货物；
	- right：所有货物的总重，即一次性能运走所有货物。

3. 根据题意明确使用搜索左侧 / 右侧的二分搜索算法，写出解法代码。
	在f(x) == target的约束下，求x的左侧边界。
*/
