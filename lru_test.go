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

func TestPutValue(t *testing.T) {
	lruCache := InitLRU()
	for _, testcase := range testCaseData {
		lruCache.putValue(testcase.key, testcase.value)
	}

	t.Log(lruCache.kv)

	for head := lruCache.head; head != nil; head = head.after {
		t.Log(head)
		t.Log(head.key, head.value)
	}
}

func TestGet(t *testing.T) {
	lruCache := InitLRU()

	for _, testcase := range testCaseData {
		lruCache.putValue(testcase.key, testcase.value)
	}

	t.Log(lruCache)

	expectValue := testCaseData[2].expectValue
	getValue := lruCache.get(testCaseData[2].key)

	t.Log(lruCache)

	headNode := lruCache.head

	if lruCache.tail.after != nil {
		t.Error("tail node after node is not nil")
	}

	for {
		t.Log(headNode)
		if headNode != nil {
			headNode = headNode.after
		}
		if headNode.after == nil {
			t.Log(headNode)
			break
		}
	}

	if expectValue != getValue {
		t.Errorf("expect : %v,but get %v", expectValue, getValue)
	}
}
