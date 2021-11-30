// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package branch

import (
	"github.com/ChainSafe/gossamer/lib/trie/encode"
	"github.com/ChainSafe/gossamer/lib/trie/node"
)

// Header creates the encoded header for the branch.
func (b *Branch) Header() (encoding []byte, err error) {
	var header byte
	if b.Value == nil {
		header = byte(node.BranchType) << 6
	} else {
		header = byte(node.BranchWithValueType) << 6
	}

	var encodedPublicKeyLength []byte
	if len(b.Key) >= 63 {
		header = header | 0x3f
		encodedPublicKeyLength, err = encode.ExtraPartialKeyLength(len(b.Key))
		if err != nil {
			return nil, err
		}
	} else {
		header = header | byte(len(b.Key))
	}

	encoding = make([]byte, 0, len(encodedPublicKeyLength)+1)
	encoding = append(encoding, header)
	encoding = append(encoding, encodedPublicKeyLength...)
	return encoding, nil
}