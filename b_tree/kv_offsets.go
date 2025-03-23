package btree

import (
	"encoding/binary"
	"fmt"
)

//func that calculates offset for given index
func offsetPos(node BNode, idx uint16) uint16 {
	if idx < 1 || idx > node.nkeys() {
		fmt.Println("idx is out of range")
	}

	return HEADER + 8 * node.nkeys() + 2 * (idx-1)
}

//func that gets offset for given index
func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}

	return binary.LittleEndian.Uint16(node[offsetPos(node, idx):])
}

//func that sets offset for given index
func (node BNode) setOffset(idx uint16, offset uint16) {
	binary.LittleEndian.PutUint16(node[offsetPos(node, idx):], offset)
}