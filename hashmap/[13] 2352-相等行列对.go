package hashmap

import (
	"strconv"
	"strings"
)

// 给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。
//
// 如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。
//
// 示例 1：
//
// 输入：grid = [[3,2,1],[1,7,6],[2,7,7]]
// 输出：1
// 解释：存在一对相等行列对：
// - (第 2 行，第 1 列)：[2,7,7]
//
// 示例 2：
//
// 输入：grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
// 输出：3
// 解释：存在三对相等行列对：
// - (第 0 行，第 0 列)：[3,1,2,2]
// - (第 2 行, 第 2 列)：[2,4,2,2]
// - (第 3 行, 第 2 列)：[2,4,2,2]
//
// 提示：
//
// n == grid.length == grid[i].length
// 1 <= n <= 200
// 1 <= grid[i][j] <= 10⁵
func equalPairs(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	counter := make(map[string]int)
	var sb strings.Builder
	for _, row := range grid {
		for _, num := range row {
			sb.WriteString(strconv.Itoa(num))
			sb.WriteString(",")
		}
		counter[sb.String()]++
		sb.Reset()
	}
	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sb.WriteString(strconv.Itoa(grid[j][i]))
			sb.WriteString(",")
		}
		res += counter[sb.String()]
		sb.Reset()
	}
	return res
}

/**
思路：
使用哈希表，先将行值进行拼接及计数。
再遍历列值，若有与行值重复部分，则累计到结果值。

注意：
- StringBuilder 可代替 += 拼接，因字符不可变，若直接拼接会产生大量临时中间值。
- 数字间需增加分隔符，否则无法识别准确数字。
*/
