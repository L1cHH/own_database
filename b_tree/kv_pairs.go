package btree

import (
	"bytes"
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

func (node BNode) nbytes() uint16 {
	return node.kvPos(node.nkeys())
}

func nodeLookupLE(node BNode, key []byte) uint16 {
	nkeys := node.nkeys()
	found := uint16(0)
	for i := uint16(1); i < nkeys; i++ {
		cmp := bytes.Compare(node.getKey(i), key)
		if cmp <= 0 {
			found = i
		}
		if cmp >= 0 {
			break
		}
	}
	return found
}

//Append a KV into the position 
func nodeAppendKV(new BNode, idx uint16, ptr uint64, key []byte, val []byte) {
	new.setPtr(idx, ptr)

	pos := new.kvPos(idx)

	//Put length of key
	binary.LittleEndian.PutUint16(new[pos:], uint16(len(key)))
	//Put length of value
	binary.LittleEndian.PutUint16(new[pos+2:], uint16(len(val)))
	//copy a key into position
	copy(new[pos+4:], key)
	//copy a val into position
	copy(new[pos+4+uint16(len(key)):], val)
	// the offset of the next key
	new.setOffset(idx+1, new.getOffset(idx) + 4 + uint16(len(key) + len(val)))
}

func nodeAppendRange(new BNode, old BNode, dstNew uint16, srcOld uint16, n uint16) {
	for i := uint16(0); i < n; i++ {
		dstIdx, oldIdx := dstNew+i, srcOld+i
		nodeAppendKV(new, dstIdx, old.getPtr(oldIdx), old.getKey(oldIdx), old.getVal(oldIdx))
	}
}