package lru

type HashLinkLru struct {
	head *Node
	tail *Node
	kv   map[string]*Node
}

var lru = &HashLinkLru{
	head: nil,
	tail: nil,
	kv:   make(map[string]*Node),
}

func InitLRU() *HashLinkLru {
	return lru
}

type Node struct {
	size     int
	capacity int
	key      string
	before   *Node
	after    *Node
	value    any
}
