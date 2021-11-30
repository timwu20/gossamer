// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package leaf

import "github.com/ChainSafe/gossamer/lib/trie/node"

// Type returns node.LeafType.
func (l *Leaf) Type() node.Type {
	return node.LeafType
}
