package tree

//给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
//
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
// 初始状态下，所有 next 指针都被设置为 NULL。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,5,6,7]
//输出：[1,#,2,3,#,4,5,6,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由
//next 指针连接，'#' 标志着每一层的结束。
//
//
//
//
//
// 示例 2:
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点的数量在
// [0, 2¹² - 1] 范围内
// -1000 <= node.val <= 1000
//
//
//
//
// 进阶：
//
//
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
//

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i < levelSize-1 {
				next := queue[0]
				cur.Next = next
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return root
}

func connectII(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		var pre *Node
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			if pre != nil {
				pre.Next = cur
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			pre = cur
		}
	}
	return root
}

/**
思路：
将二叉树的节点入列，在每层中处理Next节点的指向。
本题对层序有要求，所以需要在外层for循环中，嵌套一个层级循环：
levelSize := len(queue)
for i := 0; i < levelSize; i++ {
	// ...
}

两种做法：
第一种：cur.Next = next
需控制if i < levelSize-1，才从队列中取next。
即对于本层的最后一个节点，应保留默认赋值，即指向nil。
否则会错误指向下一层的第一个节点。

第二种：pre.Next = cur
一开始pre为nil，随后滚动更新pre，且不断从队列中取cur。
需控制if pre != nil，才将pre指向cur。
好处是，cur始终指向本层级的节点，不会越级取到下一层节点。
*/

/**
需要注意的细节：
一开始写成了
for len(queue) > 0 {
	for i := 0; i < len(queue); i++ {
		// ...
	}
}
内层for循环的次数，会随着for循环中向队列添加元素的行为而不断变动，无法得到确切的值。
*/

/**
slice := []int{1, 2, 3}

// for range 循环
// 循环开始前slice的长度已被捕获，即使修改slice也不会影响循环次数3
for i, v := range slice {
    slice = append(slice, 4) // 对循环次数无影响
    fmt.Println(v)          // 始终输出 3 次
}

// 标准 for 循环
// len(slice)每次循环都会重新计算
for i := 0; i < len(slice); i++ {
    slice = append(slice, 4) // 会导致无限循环
    fmt.Println(slice[i])    // 索引可能越界
}

应用场景建议：
- 【安全遍历动态集合】：for range
- 【实时处理动态长度】：标准for + 长度判断
- 【反向遍历】：标准for + 索引递减
*/
