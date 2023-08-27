package main

type BNode struct {
	data []byte // Can be dumpd to disk
}

const (
	BNODE_NODE = 1 // Internal nodes without values
	BNODE_LEAF = 2 // Leaf nodes with values
)

type BTree struct {
	// Pointer (a non zero page number)
	root uint64
	// Callbacks for managing on disk pages
	get func(uint64) BNode // dereference a pointer
	new func(BNode) uint64 // allocate a new page
	del func(uint64)       // deallocate a new page
}

const HEADER = 4
const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	assert(node1max <= BTREE_PAGE_SIZE)
}

func assert(b bool) {
	if !b {
		panic("Assertion Error")
	}
}
