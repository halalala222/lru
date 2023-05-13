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

func (h *HashLinkLru) initNode(key string, value any) *Node {
	return &Node{
		before: nil,
		after:  nil,
		key:    key,
		value:  value,
	}
}

func (h *HashLinkLru) linkNodeLast(node *Node) {
	var (
		last = lru.tail
	)
	lru.tail = node
	if linkIsEmpty := last == nil; linkIsEmpty {
		lru.head = node
	} else {
		node.before = last
		last.after = node
	}
}

func (h *HashLinkLru) get(key string) any {
	var (
		node *Node
		ok   bool
	)

	if node, ok = lru.kv[key]; !ok {
		return nil
	}

	return node.value
}

func (h *HashLinkLru) getOrDefault(key string, defaultValue any) any {
	var (
		node *Node
		ok   bool
	)

	if node, ok = lru.kv[key]; !ok {
		return defaultValue
	}

	return node.value
}
