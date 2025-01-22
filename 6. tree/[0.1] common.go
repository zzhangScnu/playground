package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	nodes []*TreeNode
}

func NewStack() *Stack {
	return &Stack{nodes: []*TreeNode{}}
}

func (s *Stack) Push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.nodes) == 0 {
		return nil
	}
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return node
}

func (s *Stack) Peek() *TreeNode {
	if len(s.nodes) == 0 {
		return nil
	}
	return s.nodes[len(s.nodes)-1]
}

func (s *Stack) Size() int {
	return len(s.nodes)
}

func (s *Stack) IsEmpty() bool {
	return len(s.nodes) == 0
}

type Queue struct {
	nodes []*TreeNode
}

func NewQueue() *Queue {
	return &Queue{[]*TreeNode{}}
}

func (q *Queue) Push(v *TreeNode) {
	q.nodes = append(q.nodes, v)
}

func (q *Queue) Pop() *TreeNode {
	if len(q.nodes) == 0 {
		return nil
	}
	v := q.nodes[0]
	q.nodes = q.nodes[1:]
	return v
}

func (q *Queue) Peek() *TreeNode {
	if len(q.nodes) == 0 {
		return nil
	}
	return q.nodes[0]
}

func (q *Queue) Size() int {
	return len(q.nodes)
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}
