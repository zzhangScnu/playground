package tree

import (
	"strconv"
	"strings"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	existence := make(map[string]int)
	var traverse func(node *TreeNode) string
	traverse = func(node *TreeNode) string {
		if node == nil {
			return NULL
		}
		left, right := traverse(node.Left), traverse(node.Right)
		cur := strings.Join([]string{strconv.Itoa(node.Val), left, right}, SEP)
		if existence[cur] == 1 {
			res = append(res, node)
		}
		existence[cur]++
		return cur
	}
	traverse(root)
	return res
}
