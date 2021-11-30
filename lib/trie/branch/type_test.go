// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package branch

import (
	"testing"

	"github.com/ChainSafe/gossamer/lib/trie/node"
	"github.com/stretchr/testify/assert"
)

func Test_Branch_Type(t *testing.T) {
	testCases := map[string]struct {
		branch *Branch
		Type   node.Type
	}{
		"nil value": {
			branch: &Branch{},
			Type:   node.BranchType,
		},
		"empty value": {
			branch: &Branch{
				Value: []byte{},
			},
			Type: node.BranchWithValueType,
		},
		"non empty value": {
			branch: &Branch{
				Value: []byte{1},
			},
			Type: node.BranchWithValueType,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			Type := testCase.branch.Type()

			assert.Equal(t, testCase.Type, Type)
		})
	}
}
