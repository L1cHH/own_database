package btree

import (
	"encoding/binary"
	"fmt"
)

func (node BNode) getPtr(idx uint16) uint64 {
	if idx > node.nkeys() {
		fmt.Println("index out of range")
	}

	//shift to the right by 4 bytes + index * 8 bytes
	pos := HEADER + idx * 8

	return binary.LittleEndian.Uint64(node[pos:])
}

func (node BNode) setPtr(idx uint16, value uint64) {
	if idx > node.nkeys() {
		fmt.Println("index out of range")
	}

	pos := HEADER + idx * 8

	binary.LittleEndian.PutUint64(node[pos:], value)
}