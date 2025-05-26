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

const (
	SEP  = ","
	NULL = "#"
)

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return NULL
	}
	left, right := this.serialize(root.Left), this.serialize(root.Right)
	return strings.Join([]string{strconv.Itoa(root.Val), left, right}, SEP)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == NULL {
		return nil
	}
	nodes := strings.Split(data, SEP)
	return this.buildTree(&nodes)
}

func (this *Codec) buildTree(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	node := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if node == NULL {
		return nil
	}
	val, _ := strconv.Atoi(node)
	return &TreeNode{
		Val:   val,
		Left:  this.buildTree(nodes),
		Right: this.buildTree(nodes),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// todo：在递归过程中，节点的消费顺序决定了后续处理的nodes切片的状态。每一步的消费都会改变全局（或通过指针传递的切片），后续的递归调用依赖这个状态。如果消费不及时，后续处理会使用未更新的状态，导致数据错乱。
//
//总结起来，调整消费顺序是为了确保每次递归调用处理的是正确的剩余节点，从而正确构建树的左右子树结构。否则，节点指针的位置不正确，导致子树的数据来源错误，进而生成错误的树结构。