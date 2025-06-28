package graph

import "container/list"

// 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9
// ' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
//
// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
//
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
//
// 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
//
// 示例 1:
//
// 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
// 输出：6
// 解释：
// 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
// 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
// 因为当拨动到 "0102" 时这个锁就会被锁定。
//
// 示例 2:
//
// 输入: deadends = ["8888"], target = "0009"
// 输出：1
// 解释：把最后一位反向旋转一次即可 "0000" -> "0009"。
//
// 示例 3:
//
// 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
// target = "8888"
// 输出：-1
// 解释：无法旋转到目标数字且不被锁定。
//
// 提示：
//
// 1 <= deadends.length <= 500
// deadends[i].length == 4
// target.length == 4
// target 不在 deadends 之中
// target 和 deadends[i] 仅由若干位数字组成
func openLock(deadends []string, target string) int {
	visited := make(map[string]struct{})
	for _, dead := range deadends {
		visited[dead] = struct{}{}
	}
	start := "0000"
	if _, ok := visited[start]; ok {
		return -1
	}
	if start == target {
		return 0
	}
	queue := list.New()
	queue.PushBack(start)
	visited[start] = struct{}{}
	depth := 0
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			cur := queue.Front().Value.(string)
			queue.Remove(queue.Front())
			if cur == target {
				return depth
			}
			for i := 0; i < len(cur); i++ {
				up := turn(cur, i, -1)
				if _, ok := visited[up]; !ok {
					queue.PushBack(up)
					visited[up] = struct{}{}
				}
				down := turn(cur, i, 1)
				if _, ok := visited[down]; !ok {
					queue.PushBack(down)
					visited[down] = struct{}{}
				}
			}
		}
		depth++
	}
	return -1
}

func turn(str string, index int, direct int) string {
	bytes := []byte(str)
	if bytes[index] == '0' && direct == -1 {
		bytes[index] = '9'
	} else if bytes[index] == '9' && direct == 1 {
		bytes[index] = '0'
	} else {
		bytes[index] += byte(direct)
	}
	return string(bytes)
}

/**
思路：
使用BFS进行搜索。
对4个数字每次向上/向下转动其中1位，相当于图中每个节点都有8个邻接节点。
在图中从起始节点"0000"开始，寻找到target的最短路径。路径中不能包含deadend列表中的节点。

1. 使用visited辅助数组避免重复访问；因为其可以控制路径生长的方向，所以可以将deadend中的节点初始化到其中，以避免访问这些节点；
   visited使用map结构，value为struct{}类型，减小存储空间，通过读取时的ok标志位来判断是否曾经写入；
2. Go中的slice比container/list更高效；
3. 对于向上/向下的拨动，仅需额外处理0/9这两个边界情况，其余直接与表示方向的direct -1/1 相加即可。
   Go中对string进行修改，可以通过转换为Byte数组解决；
4. 在基于当前节点扩展下一步节点时，判断下一步节点是否已被访问，若未访问过则正常加入队列并标记已访问；而不是一股脑加入队列，在出列后才判断是否重复访问。
   这样减小了队列中存储元素的数量，提高算法效率；
   注意有几个操作是需要成对出现的，如：读取后立刻标记已访问、从队列中Peak后立刻Pop；
5. depth在本层处理逻辑最后才自增，是考虑到初始化为0，且需要扩展到下一步节点之后才算是走了一步。
*/
