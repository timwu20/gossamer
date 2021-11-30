package leaf

import "github.com/ChainSafe/gossamer/lib/trie/node"

// Type returns node.LeafType.
func (l *Leaf) Type() node.Type {
	return node.LeafType
}
