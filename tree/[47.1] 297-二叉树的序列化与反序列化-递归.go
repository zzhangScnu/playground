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

/**
思路：
序列化&反序列化均使用同样的顺序、同样的处理逻辑。
【中、左、右】

为什么构建节点时，入参需要使用引用类型*[]string？
在构建当前节点时：
	return &TreeNode{
		Val:   val,
		Left:  this.buildTree(nodes),
		Right: this.buildTree(nodes),
	}
对于左子树和右子树，使用同一个地址指向的数组。所以需要进行状态共享，
在左子树使用某些元素后及时出列，避免右子树重复使用。

这里有个问题，这两种相似的写法会导致截然不同的结果：
func (this *Codec) buildTree(nodes *[]string) *TreeNode {
	node := (*nodes)[0]
	if node == NULL {
		return nil
	}
	val, _ := strconv.Atoi(node)
	*nodes = (*nodes)[1:] // 在这里消费元素
	return &TreeNode{
		Val:   val,
		Left:  this.buildTree(nodes),
		Right: this.buildTree(nodes),
	}
}
这种写法，会导致元素被根节点使用后，没有及时消费、及时出列，导致被递归this.buildTree(nodes)的左子树构造使用。
func (this *Codec) buildTree(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	node := (*nodes)[0]
	*nodes = (*nodes)[1:] // 在这里消费元素
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
这种写法，保证每次递归调用处理的是正常的剩余节点，从而正确构建左右子树。
*/

/**
一般语境下，单纯前/中/后序遍历结果是无法还原二叉树结构的，因为缺少空指针信息，
必须结合其中两种遍历结果才能还原二叉树。
但本题的序列化方法中，已将空指针放入结果集，所以反序列化时可以得到原本的二叉树结构。
*/
