// Package lru: O(1) LRU cache, map + doubly linked list se.
package lru

import "container/list"

type entry struct {
	key, val int
}

type Cache struct {
	cap   int
	ll    *list.List            // front = sabse recently used
	items map[int]*list.Element // key -> list node, O(1) lookup ke liye
}

func New(capacity int) *Cache {
	return &Cache{cap: capacity, ll: list.New(), items: map[int]*list.Element{}}
}

func (c *Cache) Get(key int) (int, bool) {
	el, ok := c.items[key]
	if !ok {
		return 0, false
	}
	c.ll.MoveToFront(el) // use hua to aage le aao
	return el.Value.(*entry).val, true
}

func (c *Cache) Put(key, val int) {
	if el, ok := c.items[key]; ok {
		el.Value.(*entry).val = val
		c.ll.MoveToFront(el)
		return
	}
	c.items[key] = c.ll.PushFront(&entry{key, val})
	if c.ll.Len() > c.cap {
		last := c.ll.Back() // sabse purana, isko nikaalo
		c.ll.Remove(last)
		delete(c.items, last.Value.(*entry).key)
	}
}

// NOTE: ye thread-safe nahi hai. Concurrent use ke liye sync.Mutex lagao.
