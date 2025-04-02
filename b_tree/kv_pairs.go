package btree

import (
	"encoding/binary"
	"fmt"
)

func (node BNode) kvPos(idx uint16) uint16 {
	if idx > node.nkeys() {
		fmt.Println("idx is out of range")
		return 0
	}
	return HEADER + 8*node.nkeys() + 2*node.nkeys() + node.getOffset(idx)
}

func (node BNode) getKey(idx uint16) []byte {
	if idx >= node.nkeys() {
		fmt.Println("idx is out of range")
		return nil
	}
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	return node[pos + 4:][:klen]
}

func (node BNode) getVal(idx uint16) []byte {
	if idx >= node.nkeys() {
		fmt.Println("idx is out of range")
		return nil
	}
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	vlen := binary.LittleEndian.Uint16(node[pos + 2:])
	return node[pos + 4 + klen:][:vlen]
}