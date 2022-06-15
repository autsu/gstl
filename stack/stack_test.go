package stack

import "testing"

func Test(t *testing.T) {
	s := New[int64]()
	e := []int64{1, 2, 3}
	for _, v := range e {
		s.Push(v)
	}
	top := s.Top()
	if top != e[len(e)-1] {
		t.Errorf("top error: want 3, get %v\n", top)
	}

	idx := len(e) - 1
	for !s.Empty() {
		p := s.Pop()
		if e[idx] != p {
			t.Errorf("pop error: want %v, get %v\n", e[idx], p)
		}
		idx--
	}
	// try pop empty stack
	s.Pop()
}
