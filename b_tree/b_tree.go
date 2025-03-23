package btree

import (
	"fmt"
	"os"
)

const HEADER = 4

const BTREE_PAGE_SIZE = 4096
const BTREE_KEY_MAX_SIZE = 1000
const BTREE_VALUE_MAX_SIZE = 3000

func init() {
	maxNodeSize := HEADER + 8 + 2 + 4 + BTREE_KEY_MAX_SIZE + BTREE_VALUE_MAX_SIZE
	if maxNodeSize > BTREE_PAGE_SIZE {
		fmt.Println("Node size is greater than BTREE_PAGE_SIZE")
		os.Exit(1)
	}
}

