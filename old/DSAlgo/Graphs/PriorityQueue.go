package main

type Path struct {
	Name int
	Weight int
	Psf string
	AcquiredVertex int

	Index int
}

type PriorityQueue []*Path

func NewPriorityQueue() PriorityQueue {
	return make(PriorityQueue, 0)
}

func (pq PriorityQueue) Len() int { return  len(pq) }

func (pq PriorityQueue) Less(i, j int) bool  {
	return pq[i].Weight < pq[j].Weight
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0: n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Path)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
