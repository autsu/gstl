package vector

type vector[T any] struct {
	element []T
}

func New[T any]() *vector[T] {
	return &vector[T]{}
}

func (v *vector[T]) PushBack(val T) {
	v.element = append(v.element, val)
}

func (v *vector[T]) PopBack() T {
	if v.Len() == 0 {
		panic("vector is empty")
	}
	p := v.element[v.Len()-1]
	v.element = v.element[:v.Len()-1]
	return p
}

func (v *vector[T]) Len() int {
	return len(v.element)
}

func (v *vector[T]) Erase(index int) T {
	if v.Len() <= index {
		panic("invalid index")
	}
	p := v.element[index]
	v.element = append(v.element[:index], v.element[index+1:])
	return p
}

func (v *vector[T]) Fill(val T) {
	for i := 0; i < v.Len(); i++ {
		v.element[i] = val
	}
}

func (v *vector[T]) Clear() {
	v.element = v.element[:0]
}
