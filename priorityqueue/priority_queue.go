package priorityqueue

type (
	priorityQueue[T any] struct {
		element []T
		cmp     func(a, b T) bool
	}
)

func New[T any](cmp func(a, b T) bool) *priorityQueue[T] {
	if cmp == nil {
		panic("fn cannot be nil")
	}
	p := &priorityQueue[T]{}
	p.cmp = cmp
	return p
}

func (p *priorityQueue[T]) Init() {
	// heapify
	n := len(p.element)
	for i := n/2 - 1; i >= 0; i-- {
		p.down(i, n)
	}
}

func (p *priorityQueue[T]) Push(val T) {
	p.element = append(p.element, val)
	p.up(len(p.element) - 1)
}

func (p *priorityQueue[T]) Empty() bool {
	return len(p.element) == 0
}

func (p *priorityQueue[T]) Pop() T {
	n := len(p.element) - 1
	p.swap(0, n)
	p.down(0, n)
	pp := p.element[len(p.element)-1]
	p.element = p.element[:len(p.element)-1]
	return pp
}

func (p *priorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !p.cmp(p.element[j], p.element[i]) {
			break
		}
		p.swap(i, j)
		j = i
	}
}

func (p *priorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && p.cmp(p.element[j2], p.element[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !p.cmp(p.element[j], p.element[i]) {
			break
		}
		p.swap(i, j)
		i = j
	}
	return i > i0
}

func (p *priorityQueue[T]) swap(i, j int) {
	p.element[i], p.element[j] = p.element[j], p.element[i]
}
