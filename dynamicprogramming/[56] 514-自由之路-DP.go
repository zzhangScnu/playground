package dynamicprogramming

import "math"

// 电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
//
// 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
//
// 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按
// 钮，以此逐个拼写完 key 中的所有字符。
//
// 旋转 ring 拼出 key 字符 key[i] 的阶段中：
//
// 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于
// 字符 key[i] 。
// 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段
// ）, 直至完成所有拼写。
//
// 示例 1：
//
// 输入: ring = "godding", key = "gd"
// 输出: 4
// 解释:
// 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
// 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
// 当然, 我们还需要1步进行拼写。
// 因此最终的输出是 4。
//
// 示例 2:
//
// 输入: ring = "godding", key = "godding"
// 输出: 13
//
// 提示：
//
// 1 <= ring.length, key.length <= 100
// ring 和 key 只包含小写英文字母
// 保证 字符串 key 一定可以由字符串 ring 旋转拼出
func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	location := make(map[uint8][]int)
	for i := 0; i < m; i++ {
		location[ring[i]] = append(location[ring[i]], i)
	}
	memo := make(map[[2]int]int)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if j == n {
			return 0
		}
		if step, ok := memo[[2]int{i, j}]; ok {
			return step
		}
		res := math.MaxInt
		target := key[j]
		for _, index := range location[target] {
			clockwiseTimes := int(math.Abs(float64(index - i)))
			step := min(clockwiseTimes, m-clockwiseTimes)
			subStep := dp(index, j+1)
			res = min(res, step+subStep+1)
		}
		memo[[2]int{i, j}] = res
		return res
	}
	return dp(0, 0)
}

/**
本题用自顶向下的方式较为合适，将原问题拆解为子问题。
想要将ring当前指针指向的第i位，旋转到key的第j位，需考虑：
- 目标位置：ring转盘中所有的key[j]字符及其位置；
- 旋转方向：顺时针/逆时针旋转。

DP函数及参数含义：
- i：ring转盘当前指针指向的位置；
- j：key当前需匹配的位置；
- dp[i][j]：将ring[i]旋转到key[j]的最小步数。

递推公式：
for ring中每一个key[j]的位置index
	dp(i, j) = min(
		子问题从ring[index]到key[j+1]的最小步数dp(index, j+1)
			+ 本次从i到index的旋转步数
			+ 本次按下按钮的次数1
	)

旋转步数：
顺时针：｜index - i｜
逆时针：len(ring) - ｜index - i｜

base case：
当j == n时，不在合法范围中，此时返回0。
*/
