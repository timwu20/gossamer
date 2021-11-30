// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package leaf

import (
	"testing"

	"github.com/ChainSafe/gossamer/lib/trie/node"
	"github.com/stretchr/testify/assert"
)

func Test_Leaf_Type(t *testing.T) {
	t.Parallel()

	leaf := new(Leaf)

	Type := leaf.Type()

	assert.Equal(t, node.LeafType, Type)
}
