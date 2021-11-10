package main

type Queue struct {
	pairs []*Pair
}

func NewQueue() *Queue {
	return &Queue{pairs: make([]*Pair, 0)}
}

func (q *Queue) Add(pair *Pair) {
	q.pairs = append(q.pairs, pair)
}

func (q *Queue) RemoveFirst() *Pair {
	if q.IsEmpty() {
		return nil
	}

	top := 	q.pairs[0]
	q.pairs = q.pairs[1:]

	return top
}

func (q *Queue) IsEmpty() bool {
	return len(q.pairs) == 0
}
