package lru

import "testing"

// table-driven test: Go ka standard tarika
func TestCache(t *testing.T) {
	c := New(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)    // 1 recently used ho gaya
	c.Put(3, 3) // 2 evict hoga

	tests := []struct {
		key    int
		wantOK bool
	}{{1, true}, {2, false}, {3, true}}

	for _, tt := range tests {
		if _, ok := c.Get(tt.key); ok != tt.wantOK {
			t.Errorf("Get(%d) ok = %v, want %v", tt.key, ok, tt.wantOK)
		}
	}
}

// benchmark: `go test -bench=. ./lru`
func BenchmarkPutGet(b *testing.B) {
	c := New(100)
	for i := 0; i < b.N; i++ {
		c.Put(i%200, i)
		c.Get(i % 200)
	}
}
