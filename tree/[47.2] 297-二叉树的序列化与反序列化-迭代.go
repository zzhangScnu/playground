package tree

import (
	"strconv"
	"strings"
)

//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方
//式重构得到原数据。
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串
//反序列化为原始的树结构。
//
// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的
//方法解决这个问题。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,null,4,5]
//输出：[1,2,3,null,null,4,5]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 10⁴] 内
// -1000 <= Node.val <= 1000

type IterativeCodec struct {
}

func IterativeCodecConstructor() IterativeCodec {
	return IterativeCodec{}
}

// Serializes a tree to a single string.
func (this *IterativeCodec) serialize(root *TreeNode) string {
	if root == nil {
		return NULL
	}
	var res []string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			res = append(res, NULL)
		} else {
			res = append(res, strconv.Itoa(cur.Val))
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}
	return strings.Join(res, SEP)
}

// Deserializes your encoded data to tree.
func (this *IterativeCodec) deserialize(data string) *TreeNode {
	if data == NULL {
		return nil
	}
	index := 0
	nodes := strings.Split(data, SEP)
	root := buildNode(index, nodes)
	index++
	queue := []*TreeNode{root}
	for len(queue) > 0 && index < len(nodes) {
		cur := queue[0]
		queue = queue[1:]
		left := buildNode(index, nodes)
		cur.Left = left
		index++
		if left != nil {
			queue = append(queue, left)
		}
		right := buildNode(index, nodes)
		cur.Right = right
		index++
		if right != nil {
			queue = append(queue, right)
		}
	}
	return root
}

func buildNode(index int, nodes []string) *TreeNode {
	if index >= len(nodes) {
		return nil
	}
	val := nodes[index]
	if val == NULL {
		return nil
	}
	v, _ := strconv.Atoi(val)
	return &TreeNode{Val: v}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

/**
使用队列作为辅助，层序遍历序列化二叉树、再层序反序列化二叉树。
*/
