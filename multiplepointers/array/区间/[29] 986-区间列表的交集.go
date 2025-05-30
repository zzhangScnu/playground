package 区间

// 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而
// secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
//
// 返回这 两个区间列表的交集 。
//
// 形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
//
// 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。
//
// 示例 1：
//
// 输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,
// 24],[25,26]]
// 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
//
// 示例 2：
//
// 输入：firstList = [[1,3],[5,9]], secondList = []
// 输出：[]
//
// 示例 3：
//
// 输入：firstList = [], secondList = [[4,8],[10,12]]
// 输出：[]
//
// 示例 4：
//
// 输入：firstList = [[1,7]], secondList = [[3,10]]
// 输出：[[3,7]]
//
// 提示：
//
// 0 <= firstList.length, secondList.length <= 1000
// firstList.length + secondList.length >= 1
// 0 <= starti < endi <= 10⁹
// endi < starti+1
// 0 <= startj < endj <= 10⁹
// endj < startj+1
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	f, s := 0, 0
	for f < len(firstList) && s < len(secondList) {
		if secondList[s][0] <= firstList[f][1] && firstList[f][0] <= secondList[s][1] {
			res = append(res, []int{max(firstList[f][0], secondList[s][0]), min(firstList[f][1], secondList[s][1])})
		}
		if firstList[f][1] < secondList[s][1] {
			f++
		} else {
			s++
		}
	}
	return res
}

/**
思路：双指针
用两个指针指向不同区间：
1. 判断两个区间是否有交集。若有交集，获取交集区间，即[max(左边界), min(右边界)]；
2. 推进区间。将身位靠后的区间进行推进。
    举例，若firstList[f][1] < secondList[s][1]，则f++，secondList不动，
    以便下一个firstList可以再和当前的secondList进行比对，寻求剩余可能的交集。
	注意，判断条件是firstList[f][1] < secondList[s][1]，而不是f < s。

为什么不需要排序？
——注意，输入已经有序，否则需要提前排序。
*/
