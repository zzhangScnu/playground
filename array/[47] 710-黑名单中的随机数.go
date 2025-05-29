package array

import (
	"math/rand"
)

//给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入 黑名单
//blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。
//
// 优化你的算法，使它最小化调用语言 内置 随机函数的次数。
//
// 实现 Solution 类:
//
//
// Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
// int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
//
//
//
//
// 示例 1：
//
//
//输入
//["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
//[[7, [2, 3, 5]], [], [], [], [], [], [], []]
//输出
//[null, 0, 4, 1, 6, 1, 0, 4]
//
//解释
//Solution solution = new Solution(7, [2, 3, 5]);
//solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
//                 // 0、1、4和6的返回概率必须相等(即概率为1/4)。
//solution.pick(); // 返回 4
//solution.pick(); // 返回 1
//solution.pick(); // 返回 6
//solution.pick(); // 返回 1
//solution.pick(); // 返回 0
//solution.pick(); // 返回 4
//
//
//
//
// 提示:
//
//
// 1 <= n <= 10⁹
// 0 <= blacklist.length <= min(10⁵, n - 1)
// 0 <= blacklist[i] < n
// blacklist 中所有值都 不同
// pick 最多被调用 2 * 10⁴ 次

type Solution struct {
	blackNumRelocation map[int]int
	blackNumThreshold  int
}

func SolutionConstructor(n int, blacklist []int) Solution {
	blackNumThreshold, blackNumRelocation := n-len(blacklist), make(map[int]int)
	for _, blackNum := range blacklist {
		blackNumRelocation[blackNum] = -1
	}
	relocation := n - 1
	for _, blackNum := range blacklist {
		if blackNum >= blackNumThreshold {
			continue
		}
		for blackNumRelocation[relocation] == -1 {
			relocation--
		}
		blackNumRelocation[blackNum] = relocation
		relocation--
	}
	return Solution{
		blackNumRelocation: blackNumRelocation,
		blackNumThreshold:  blackNumThreshold,
	}
}

func (this *Solution) Pick() int {
	num := rand.Intn(this.blackNumThreshold)
	if whiteNum, ok := this.blackNumRelocation[num]; ok {
		return whiteNum
	}
	return num
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */

/**
最暴力的方法：
生成[0, n - 1]随机数，如果其在blacklist中，重新生成并检查，直到获取非黑名单元素。


最直接的想法：
type Solution struct {
	data []int
}

func SolutionConstructor(n int, blacklist []int) Solution {
	rand.Seed(time.Now().UnixNano())
	blackNumLocation := make(map[int]interface{})
	for _, blackNum := range blacklist {
		blackNumLocation[blackNum] = true
	}
	var data []int
	for num := 0; num < n; num++ {
		if _, ok := blackNumLocation[num]; ok {
			continue
		}
		data = append(data, num)
	}
	return Solution{data: data}
}

func (this *Solution) Pick() int {
	index := rand.Intn(len(this.data))
	return this.data[index]
}
维护一个合法元素列表，初始化时，若元素在黑名单中，则跳过不处理。
这种方法会在数据量大时超时。


最终解决思路：
虚拟一个元素全集列表，其中有一条分界线blackNumThreshold，前面是合法元素，后面是黑名单元素。
随机取数时，从前面部分获取即可。

在实际实现时，无需真实维护列表，也无需将黑名单元素搬移到列表末尾，而是：
维护一个【黑名单元素 -> 合法元素】映射blackNumRelocation，如果随机到黑名单元素，直接返回其对应的合法元素即可。
处理步骤：
1. 如果黑名单元素本身就在blackNumThreshold之后，即处于虚拟列表的黑名单区域中，则无需额外维护其映射关系；
2. 如果黑名单元素在blackNumThreshold之前，即处于虚拟列表的合法区域中，则从黑名单区域中寻找一个合法元素，与其建联。
	这个合法元素，可以从虚拟列表末尾开始往前找，并在被使用后动态更新。
    如何表明这个元素是合法元素呢？：
	将blackNumRelocation的值初始化为-1，如果寻找合法元素时，检查其在blackNumRelocation中的值为-1，则表示是黑名单元素，需要继续往前找。
*/
