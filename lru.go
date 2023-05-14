package lru

type HashLinkLru struct {
	head     *Node
	tail     *Node
	kv       map[string]*Node
	capacity int
	size     int
}

var lru = &HashLinkLru{
	head:     nil,
	tail:     nil,
	kv:       make(map[string]*Node),
	size:     0,
	capacity: 0,
}

func InitLRU(capacity int) *HashLinkLru {
	lru.capacity = capacity
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

// afterNodeAccess hook method used at access the node data
func (h *HashLinkLru) afterNodeAccess(node *Node) {
	var (
		beforeNode = node.before
		afterNode  = node.after
		lastNode   = lru.tail
	)

	node.after = nil

	if isLastNode := node == lru.tail; isLastNode {
		return
	}

	if beforeNode == nil {
		lru.head = afterNode
	} else {
		beforeNode.after = afterNode
	}

	lastNode.after = node
	node.before = lastNode

	lru.tail = node
}

func (h *HashLinkLru) get(key string) any {
	var (
		node *Node
		ok   bool
	)

	if node, ok = lru.kv[key]; !ok {
		return nil
	}

	h.afterNodeAccess(node)

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

	h.afterNodeAccess(node)

	return node.value
}

func (h *HashLinkLru) putValue(key string, value any) {
	if key == "" || value == nil {
		return
	}

	node := h.initNode(key, value)

	lru.kv[key] = node
	lru.size++

	h.linkNodeLast(node)
}

func (h *HashLinkLru) removeNode(node *Node) {
	if _, ok := lru.kv[node.key]; !ok {
		return
	}

	delete(lru.kv, node.key)

	var (
		before = node.before
		after  = node.after
	)

	if before != nil {
		before.after = after
	} else {
		lru.head = after
	}

	if after != nil {
		after.before = before
	} else {
		lru.tail = before
	}
}
