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
