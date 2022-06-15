package priorityqueue

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	p := New[int64](func(i, j int64) bool { return i < j })
	p.Push(3)
	p.Push(2)
	p.Push(4)

	for !p.Empty() {
		println(p.Pop())
	}
}

func Test1(t *testing.T) {
	p := New[time.Time](func(i, j time.Time) bool { return i.Before(j) })
	p.Push(time.Date(2020, 02, 27, 20, 20,
		20, 20, time.Local))
	p.Push(time.Date(2018, 02, 27, 20, 20,
		20, 20, time.Local))
	p.Push(time.Date(2022, 02, 27, 20, 20,
		20, 20, time.Local))
	for !p.Empty() {
		fmt.Println(p.Pop())
	}
}
