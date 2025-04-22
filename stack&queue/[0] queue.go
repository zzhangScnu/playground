package stack_queue

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{[]int{}}
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() int {
	if len(q.data) == 0 {
		return -1
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue) Peek() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
