package graph

// 基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
//
// 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
//
// 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
//
// 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
//
// 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成
// 此基因变化，返回 -1 。
//
// 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
//
// 示例 1：
//
// 输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
// 输出：1
//
// 示例 2：
//
// 输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA",
// "AAACGGTA"]
// 输出：2
//
// 示例 3：
//
// 输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC",
// "AACCCCCC"]
// 输出：3
//
// 提示：
//
// start.length == 8
// end.length == 8
// 0 <= bank.length <= 10
// bank[i].length == 8
// start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成

type Gene struct {
	gene string
	step int
}

func NewGene(gene string, step int) Gene {
	return Gene{
		gene: gene,
		step: step,
	}
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	movements := []uint8{'A', 'T', 'C', 'G'}
	valid := make(map[string]interface{})
	for _, b := range bank {
		valid[b] = struct{}{}
	}
	if _, ok := valid[endGene]; !ok {
		return -1
	}
	queue := []Gene{NewGene(startGene, 0)}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.gene == endGene {
			return cur.step
		}
		gene := []byte(cur.gene)
		for i := 0; i < len(gene); i++ {
			ch := gene[i]
			for _, newCh := range movements {
				if newCh == ch {
					continue
				}
				gene[i] = newCh
				if _, ok := valid[string(gene)]; ok {
					queue = append(queue, NewGene(string(gene), cur.step+1))
					delete(valid, string(gene))
				}
			}
			gene[i] = ch
		}
	}
	return -1
}

/**
思路：
通用的广度优先搜索方式。

注意：
- 合并【合法字符串集合】和【已访问字符串集合】。
*/
