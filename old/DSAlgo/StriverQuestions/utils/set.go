package utils

var exists = struct{}{}

type runeSet struct {
	m map[rune]struct{}
}

func NewRuneSet() *runeSet {
	s := &runeSet{}
	s.m = make(map[rune]struct{})
	return s
}

func (s *runeSet) Add(value rune) {
	s.m[value] = exists
}

func (s *runeSet) Remove(value rune) {
	delete(s.m, value)
}

func (s *runeSet) Contains(value rune) bool {
	_, c := s.m[value]
	return c
}

