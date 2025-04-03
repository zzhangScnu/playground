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
在 Go 语言中， for range 循环的行为与标准 for 循环有本质区别。这是两种循环的关键差异：

```go
// for range 循环示例
slice := []int{1, 2, 3}
for i, v := range slice {
    // 循环开始前 slice 的长度 (3) 已被捕获
    // 即使修改 slice 也不会影响循环次数
    slice = append(slice, 4) // 对循环次数无影响
    fmt.Println(v)          // 始终输出 3 次
}

// 对比标准 for 循环
for i := 0; i < len(slice); i++ {
    // len(slice) 每次都会重新计算
    slice = append(slice, 4) // 会导致无限循环
    fmt.Println(slice[i])    // 索引可能越界
}
 ```

应用场景建议：

- 需要 安全遍历动态集合 时用 for range
- 需要 实时处理动态长度 时用标准 for + 长度判断
- 需要 反向遍历 时用标准 for i := len(slice)-1; i >= 0; i--
*/
