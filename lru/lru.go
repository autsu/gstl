package lru

import "github.com/youseebiggirl/gstl/list"

type kv[T any] struct {
	k string
	v T
}

type lru[T any] struct {
	m   map[string]*list.Element[*kv[T]]
	ls  *list.List[*kv[T]]
	cap int64
	len int64
}

func New[T any](capacity int64) *lru[T] {
	return &lru[T]{
		m:   make(map[string]*list.Element[*kv[T]]),
		ls:  list.New[*kv[T]](),
		cap: capacity,
	}
}

func (l *lru[T]) Get(key string) T {
	e, ok := l.m[key]
	if !ok {
		var emptyVal T
		return emptyVal
	}
	ret := e.Value.v
	l.ls.MoveToFront(e)
	return ret
}

func (l *lru[T]) Put(key string, value T) {
	if e, ok := l.m[key]; ok {
		e.Value.v = value
		l.ls.MoveToFront(e)
		return
	}
	if l.cap == l.len {
		kv := l.ls.Remove(l.ls.Back())
		delete(l.m, kv.k)
		l.len--
	}

	kv := &kv[T]{
		k: key,
		v: value,
	}
	e := l.ls.PushFront(kv)
	l.m[key] = e
	l.len++
}
