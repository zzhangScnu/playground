package greedy

import "strings"

// Dota2 的世界里有两个阵营：Radiant（天辉）和 Dire（夜魇）
//
// Dota2 参议院由来自两派的参议员组成。现在参议院希望对一个 Dota2 游戏里的改变作出决定。他们以一个基于轮为过程的投票进行。在每一轮中，每一位参议
// 员都可以行使两项权利中的 一 项：
//
// 禁止一名参议员的权利：参议员可以让另一位参议员在这一轮和随后的几轮中丧失 所有的权利 。
// 宣布胜利：如果参议员发现有权利投票的参议员都是 同一个阵营的 ，他可以宣布胜利并决定在游戏中的有关变化。
//
// 给你一个字符串 senate 代表每个参议员的阵营。字母 'R' 和 'D'分别代表了 Radiant（天辉）和 Dire（夜魇）。然后，如果有 n 个参
// 议员，给定字符串的大小将是 n。
//
// 以轮为基础的过程从给定顺序的第一个参议员开始到最后一个参议员结束。这一过程将持续到投票结束。所有失去权利的参议员将在过程中被跳过。
//
// 假设每一位参议员都足够聪明，会为自己的政党做出最好的策略，你需要预测哪一方最终会宣布胜利并在 Dota2 游戏中决定改变。输出应该是 "Radiant"
// 或 "Dire" 。
//
// 示例 1：
//
// 输入：senate = "RD"
// 输出："Radiant"
// 解释：
// 第 1 轮时，第一个参议员来自 Radiant 阵营，他可以使用第一项权利让第二个参议员失去所有权利。
// 这一轮中，第二个参议员将会被跳过，因为他的权利被禁止了。
// 第 2 轮时，第一个参议员可以宣布胜利，因为他是唯一一个有投票权的人。
//
// 示例 2：
//
// 输入：senate = "RDD"
// 输出："Dire"
// 解释：
// 第 1 轮时，第一个来自 Radiant 阵营的参议员可以使用第一项权利禁止第二个参议员的权利。
// 这一轮中，第二个来自 Dire 阵营的参议员会将被跳过，因为他的权利被禁止了。
// 这一轮中，第三个来自 Dire 阵营的参议员可以使用他的第一项权利禁止第一个参议员的权利。
// 因此在第二轮只剩下第三个参议员拥有投票的权利,于是他可以宣布胜利
//
// 提示：
//
// n == senate.length
// 1 <= n <= 10⁴
// senate[i] 为 'R' 或 'D'
func predictPartyVictory(senate string) string {
	rc, dc := 0, 0
	for _, s := range senate {
		if s == 'R' {
			rc++
		} else {
			dc++
		}
	}
	senates := strings.Split(senate, "")
	for {
		for i, senate := range senates {
			if senate == "R" {
				if dc == 0 {
					return "Radiant"
				}
				dc = banAfter(senates, "D", dc, i)
			} else if senate == "D" {
				if rc == 0 {
					return "Dire"
				}
				rc = banAfter(senates, "R", rc, i)
			}
		}
	}
}

func banAfter(senates []string, senate string, senateCount int, position int) int {
	n := len(senates)
	var index int
	for offset := 1; offset <= n; offset++ {
		index = (position + offset) % n
		if senates[index] == senate {
			senates[index] = ""
			return senateCount - 1
		}
	}
	return senateCount
}

/**
思路：
对每一个议员，模拟禁用其后对立阵营议员的权利。
即对每一个R / D，消除其后的D / R。
首先计算R和D的数量，如果轮到R / D时，发现对立阵营人数归零，则可宣布胜利；
否则禁用对立阵营议员，且将对立阵营人数-1。
注意这里应判断，如果禁用操作发动成功，才需减去对立阵营人数，
否则可能会出现在场均为自己人，已经无人可禁，但还是做了减法，导致出现负数。

注意存在环形场景，即遍历范围应为[0, i-1], [i+1, n)。
可以写2个for循环来覆盖，
也可以采用取模方式，偏移量范围为[1, n]，即从下一个开始，向后覆盖n个人，绕一圈回到自身。
for offset := 1; offset <= n; offset++ {
	index = (position + offset) % n
	// ...
}
*/
