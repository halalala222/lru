package lru

import "testing"

func TestLinkNodeLast(t *testing.T) {
	lruCache := InitLRU()
	lruCache.linkNodeLast(lru.initNode("test", 1))
	t.Log(lruCache)
}
