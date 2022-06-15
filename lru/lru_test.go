package lru

import "testing"

func TestLRU(t *testing.T) {
	//["LRUCache","get","put","get","put","put","get","get"]
	//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	// [null,-1,null,-1,null,null,2,6]
	l := New[string](2)
	r := l.Get("2")
	if r != "" {
		t.Errorf("get(2) want empty str, get %v\n", r)
	}
	l.Put("2", "6")
	r = l.Get("1")
	if r != "" {
		t.Errorf("get(1) want empty str, get %v\n", r)
	}
	l.Put("1", "5")
	l.Put("1", "2")
	r = l.Get("1")
	if r != "2" {
		t.Errorf("get(1) want 2, get %v\n", r)
	}
	r = l.Get("2")
	if r != "6" {
		t.Errorf("get(2) want 6, get %v\n", r)
	}
}
