package branch

import "github.com/ChainSafe/gossamer/lib/trie/node"

// Type returns node.BranchType if the branch value
// is nil, and node.BranchWithValueType otherwise.
func (b *Branch) Type() node.Type {
	if b.Value == nil {
		return node.BranchType
	}
	return node.BranchWithValueType
}
