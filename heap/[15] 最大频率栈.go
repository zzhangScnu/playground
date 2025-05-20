package heap

import "container/list"

type MaxFrequencyStack struct {
	KF           map[int]int
	FK           map[int]*list.List
	MaxFrequency int
}

func MaxFrequencyStackConstructor() MaxFrequencyStack {
	return MaxFrequencyStack{
		KF:           make(map[int]int),
		FK:           make(map[int]*list.List),
		MaxFrequency: 0,
	}
}

func (m MaxFrequencyStack) Push(x int) {
	m.KF[x]++
	if m.FK[m.KF[x]] == nil {
		m.FK[m.KF[x]] = list.New()
	}
	m.FK[m.KF[x]].PushBack(x)
	if m.MaxFrequency < m.KF[x] {
		m.MaxFrequency = m.KF[x]
	}
}

func (m MaxFrequencyStack) Pop() int {
	x := m.FK[m.MaxFrequency].Front().Value.(int)
	m.FK[m.MaxFrequency].Remove(m.FK[m.MaxFrequency].Front())
	m.KF[x]--
	if m.FK[m.MaxFrequency].Len() == 0 {
		m.MaxFrequency--
	}
	return x
}
