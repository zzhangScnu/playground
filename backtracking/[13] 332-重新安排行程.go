package backtracking

import (
	"sort"
)

// 给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
//
// 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
//
// 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
//
// 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
//
// 示例 1：
//
// 输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
// 输出：["JFK","MUC","LHR","SFO","SJC"]
//
// 示例 2：
//
// 输入：tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL",
// "SFO"]]
// 输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
// 解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
//
// 提示：
//
// 1 <= tickets.length <= 300
// tickets[i].length == 2
// fromi.length == 3
// toi.length == 3
// fromi 和 toi 由大写英文字母组成
// fromi != toi

type Pair struct {
	To      string
	Visited bool
}

type Pairs []*Pair

func (p Pairs) Len() int {
	return len(p)
}
func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Pairs) Less(i, j int) bool {
	return p[i].To < p[j].To
}

func findItinerary(tickets [][]string) []string {
	routes := make(map[string]Pairs)
	for _, ticket := range tickets {
		routes[ticket[0]] = append(routes[ticket[0]], &Pair{To: ticket[1], Visited: false})
	}
	for _, pairs := range routes {
		sort.Sort(pairs)
	}
	path := []string{"JFK"}
	var doFindItinerary func() bool
	doFindItinerary = func() bool {
		if len(path) == len(tickets)+1 {
			return true
		}
		from := path[len(path)-1]
		for i, pair := range routes[from] {
			if pair.Visited {
				continue
			}
			if i > 0 && routes[from][i-1].To == pair.To && !routes[from][i-1].Visited {
				continue
			}
			pair.Visited = true
			path = append(path, pair.To)
			if doFindItinerary() {
				return true
			}
			path = path[:len(path)-1]
			pair.Visited = false
		}
		return false
	}
	doFindItinerary()
	return path
}

/**
函数签名的作用：
- 无：遍历整棵树；
- 有：一般是bool，当找到某个满足条件的结果时，提前返回，无需再向下递归&回溯：
	if len(path) == len(tickets)+1 {
		return true
	}
	// ...
	if doFindItinerary() {
		return true // 当方法返回true时，不断向上返回true，不进入下面的递归和回溯逻辑
	}
	path = path[:len(path)-1]
	pair.Visited = false
  // ...
*/

/**
计数的作用：防止重复选取导致死循环；
排序的实现：自定义元素结构体，并实现sort.Interface接口，可使用sort.Sort()方法；
Pair的引用类型：【type Pairs []*Pair】-> 可直接更改Pair结构体的Visited字段，而不是副本值；
剪枝：本题最后一个测试用例会导致超时，所以需要引入剪枝逻辑：
	if i > 0 && routes[from][i-1].To == pair.To && !routes[from][i-1].Visited {
		continue
	}
	如果上一张票未被选取(!routes[from][i-1].Visited)，表示本次处理的分支是由兄弟节点回溯到父节点，再遍历下来的；
	由于提前排序，机票之间非递减。如果此次的票(from, to)跟上一张票重复(routes[from][i-1].To == pair.To)，则兄弟节点的子树一定涵盖了本子树；
	综上，本子树的遍历可以跳过。
*/

/**
可以用嵌套map+计数值实现，但Pair结构体会更直观一些：
func findItinerary(tickets [][]string) []string {
	routes := make(map[string]map[string]int)
	for _, ticket := range tickets {
		if routes[ticket[0]] == nil {
			routes[ticket[0]] = make(map[string]int)
		}
		routes[ticket[0]][ticket[1]]++
	}
	// 但这里排序不好做，slices.Sort()不支持，且map是随机访问，无法保证顺序
	path := []string{"JFK"}
	var doFindItinerary func() bool
	doFindItinerary = func() bool {
		if len(path) == len(tickets)+1 {
			return true
		}
		// 即使直接从path取最末作为本次处理的元素，而不是作为入参，也要提前将JFK加入结果集。因为每层处理逻辑是将to加入结果集的
		for to, times := range routes[path[len(path)-1]] {
			if times <= 0 {
				continue
			}
			routes[path[len(path)-1]][to]-- // 这里跟下一句要注意先后顺序
			path = append(path, to)
			if doFindItinerary() {
				return true
			}
			path = path[:len(path)-1]
			routes[path[len(path)-1]][to]++
		}
		return false
	}
	doFindItinerary()
	return path
}
*/
