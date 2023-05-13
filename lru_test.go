package lru

import (
	"testing"
)

type kv struct {
	key         string
	value       any
	expectValue any
}

var testCaseData = []kv{
	{
		key:         "1",
		value:       1,
		expectValue: 1,
	},
	{
		key:         "2",
		value:       2,
		expectValue: 2,
	},
	{
		key:         "3",
		value:       3,
		expectValue: 3,
	},
	{
		key:         "4",
		value:       4,
		expectValue: 4,
	},
	{
		key:         "5",
		value:       5,
		expectValue: 5,
	},
}

func TestLinkNodeLast(t *testing.T) {
	lruCache := InitLRU()
	lruCache.linkNodeLast(lru.initNode("test", 1))
	t.Log(lruCache)
}
