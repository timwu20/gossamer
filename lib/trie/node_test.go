// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"

	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/pkg/scale"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// byteArray makes byte array with length specified; used to test byte array encoding
func byteArray(length int) []byte {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = 0xf
	}
	return b
}

func generateRand(size int) [][]byte {
	rt := make([][]byte, size)
	for i := range rt {
		buf := make([]byte, rand.Intn(379)+1)
		rand.Read(buf)
		rt[i] = buf
	}
	return rt
}

func TestChildrenBitmap(t *testing.T) {
	b := &Branch{children: [16]Node{}}
	res := b.childrenBitmap()
	if res != 0 {
		t.Errorf("Fail to get children bitmap: got %x expected %x", res, 1)
	}

	b.children[0] = &Leaf{key: []byte{0x00}, value: []byte{0x00}}
	res = b.childrenBitmap()
	if res != 1 {
		t.Errorf("Fail to get children bitmap: got %x expected %x", res, 1)
	}

	b.children[4] = &Leaf{key: []byte{0x00}, value: []byte{0x00}}
	res = b.childrenBitmap()
	if res != 1<<4+1 {
		t.Errorf("Fail to get children bitmap: got %x expected %x", res, 17)
	}

	b.children[15] = &Leaf{key: []byte{0x00}, value: []byte{0x00}}
	res = b.childrenBitmap()
	if res != 1<<15+1<<4+1 {
		t.Errorf("Fail to get children bitmap: got %x expected %x", res, 257)
	}
}

func TestBranchHeader(t *testing.T) {
	tests := []struct {
		br     *Branch
		header []byte
	}{
		{&Branch{key: nil, children: [16]Node{}, value: nil}, []byte{0x80}},
		{&Branch{key: []byte{0x00}, children: [16]Node{}, value: nil}, []byte{0x81}},
		{&Branch{key: []byte{0x00, 0x00, 0xf, 0x3}, children: [16]Node{}, value: nil}, []byte{0x84}},

		{&Branch{key: nil, children: [16]Node{}, value: []byte{0x01}}, []byte{0xc0}},
		{&Branch{key: []byte{0x00}, children: [16]Node{}, value: []byte{0x01}}, []byte{0xc1}},
		{&Branch{key: []byte{0x00, 0x00}, children: [16]Node{}, value: []byte{0x01}}, []byte{0xc2}},
		{&Branch{key: []byte{0x00, 0x00, 0xf}, children: [16]Node{}, value: []byte{0x01}}, []byte{0xc3}},

		{&Branch{key: byteArray(62), children: [16]Node{}, value: nil}, []byte{0xbe}},
		{&Branch{key: byteArray(62), children: [16]Node{}, value: []byte{0x00}}, []byte{0xfe}},
		{&Branch{key: byteArray(63), children: [16]Node{}, value: nil}, []byte{0xbf, 0}},
		{&Branch{key: byteArray(64), children: [16]Node{}, value: nil}, []byte{0xbf, 1}},
		{&Branch{key: byteArray(64), children: [16]Node{}, value: []byte{0x01}}, []byte{0xff, 1}},

		{&Branch{key: byteArray(317), children: [16]Node{}, value: []byte{0x01}}, []byte{255, 254}},
		{&Branch{key: byteArray(318), children: [16]Node{}, value: []byte{0x01}}, []byte{255, 255, 0}},
		{&Branch{key: byteArray(573), children: [16]Node{}, value: []byte{0x01}}, []byte{255, 255, 255, 0}},
	}

	for _, test := range tests {
		test := test
		res, err := test.br.header()
		if err != nil {
			t.Fatalf("Error when encoding header: %s", err)
		} else if !bytes.Equal(res, test.header) {
			t.Errorf("Branch header fail case %v: got %x expected %x", test.br, res, test.header)
		}
	}
}

func TestFailingPk(t *testing.T) {
	tests := []struct {
		br     *Branch
		header []byte
	}{
		{&Branch{key: byteArray(2 << 16), children: [16]Node{}, value: []byte{0x01}}, []byte{255, 254}},
	}

	for _, test := range tests {
		_, err := test.br.header()
		if err == nil {
			t.Fatalf("should error when encoding node w pk length > 2^16")
		}
	}
}

func TestLeafHeader(t *testing.T) {
	tests := []struct {
		br     *Leaf
		header []byte
	}{
		{&Leaf{key: nil, value: nil}, []byte{0x40}},
		{&Leaf{key: []byte{0x00}, value: nil}, []byte{0x41}},
		{&Leaf{key: []byte{0x00, 0x00, 0xf, 0x3}, value: nil}, []byte{0x44}},
		{&Leaf{key: byteArray(62), value: nil}, []byte{0x7e}},
		{&Leaf{key: byteArray(63), value: nil}, []byte{0x7f, 0}},
		{&Leaf{key: byteArray(64), value: []byte{0x01}}, []byte{0x7f, 1}},

		{&Leaf{key: byteArray(318), value: []byte{0x01}}, []byte{0x7f, 0xff, 0}},
		{&Leaf{key: byteArray(573), value: []byte{0x01}}, []byte{0x7f, 0xff, 0xff, 0}},
	}

	for i, test := range tests {
		test := test
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, err := test.br.header()
			if err != nil {
				t.Fatalf("Error when encoding header: %s", err)
			} else if !bytes.Equal(res, test.header) {
				t.Errorf("Leaf header fail: got %x expected %x", res, test.header)
			}
		})
	}
}

func TestBranchEncode(t *testing.T) {
	randKeys := generateRand(101)
	randVals := generateRand(101)

	for i, testKey := range randKeys {
		b := &Branch{key: testKey, children: [16]Node{}, value: randVals[i]}
		expected := bytes.NewBuffer(nil)

		header, err := b.header()
		if err != nil {
			t.Fatalf("Error when encoding header: %s", err)
		}

		expected.Write(header)
		expected.Write(nibblesToKeyLE(b.key))
		expected.Write(common.Uint16ToBytes(b.childrenBitmap()))

		enc, err := scale.Marshal(b.value)
		if err != nil {
			t.Fatalf("Fail when encoding value with scale: %s", err)
		}

		expected.Write(enc)

		for _, child := range b.children {
			if child == nil {
				continue
			}

			err := hashNode(child, expected)
			require.NoError(t, err)
		}

		buffer := bytes.NewBuffer(nil)
		const parallel = false
		err = encodeBranch(b, buffer, parallel)
		require.NoError(t, err)
		assert.Equal(t, expected.Bytes(), buffer.Bytes())
	}
}

func TestLeafEncode(t *testing.T) {
	randKeys := generateRand(100)
	randVals := generateRand(100)

	for i, testKey := range randKeys {
		l := &Leaf{key: testKey, value: randVals[i]}
		expected := []byte{}

		header, err := l.header()
		if err != nil {
			t.Fatalf("Error when encoding header: %s", err)
		}
		expected = append(expected, header...)
		expected = append(expected, nibblesToKeyLE(l.key)...)

		enc, err := scale.Marshal(l.value)
		if err != nil {
			t.Fatalf("Fail when encoding value with scale: %s", err)
		}

		expected = append(expected, enc...)

		buffer := bytes.NewBuffer(nil)
		err = encodeLeaf(l, buffer)
		require.NoError(t, err)
		assert.Equal(t, expected, buffer.Bytes())
	}
}

func TestEncodeRoot(t *testing.T) {
	trie := NewEmptyTrie()

	for i := 0; i < 20; i++ {
		rt := GenerateRandomTests(t, 16)
		for _, test := range rt {
			trie.Put(test.key, test.value)

			val := trie.Get(test.key)
			if !bytes.Equal(val, test.value) {
				t.Errorf("Fail to get key %x with value %x: got %x", test.key, test.value, val)
			}

			buffer := bytes.NewBuffer(nil)
			const parallel = false
			err := encodeNode(trie.root, buffer, parallel)
			require.NoError(t, err)
		}
	}
}

func TestBranchDecode(t *testing.T) {
	tests := []*Branch{
		{key: []byte{}, children: [16]Node{}, value: nil},
		{key: []byte{0x00}, children: [16]Node{}, value: nil},
		{key: []byte{0x00, 0x00, 0xf, 0x3}, children: [16]Node{}, value: nil},
		{key: []byte{}, children: [16]Node{}, value: []byte{0x01}},
		{key: []byte{}, children: [16]Node{&Leaf{}}, value: []byte{0x01}},
		{key: []byte{}, children: [16]Node{&Leaf{}, nil, &Leaf{}}, value: []byte{0x01}},
		{
			key: []byte{},
			children: [16]Node{
				&Leaf{}, nil, &Leaf{}, nil,
				nil, nil, nil, nil,
				nil, &Leaf{}, nil, &Leaf{},
			},
			value: []byte{0x01},
		},
		{key: byteArray(62), children: [16]Node{}, value: nil},
		{key: byteArray(63), children: [16]Node{}, value: nil},
		{key: byteArray(64), children: [16]Node{}, value: nil},
		{key: byteArray(317), children: [16]Node{}, value: []byte{0x01}},
		{key: byteArray(318), children: [16]Node{}, value: []byte{0x01}},
		{key: byteArray(573), children: [16]Node{}, value: []byte{0x01}},
	}

	buffer := bytes.NewBuffer(nil)
	const parallel = false

	for _, test := range tests {
		err := encodeBranch(test, buffer, parallel)
		require.NoError(t, err)

		res := new(Branch)
		err = res.Decode(buffer, 0)

		require.NoError(t, err)
		require.Equal(t, test.key, res.key)
		require.Equal(t, test.childrenBitmap(), res.childrenBitmap())
		require.Equal(t, test.value, res.value)
	}
}

func TestLeafDecode(t *testing.T) {
	tests := []*Leaf{
		{key: []byte{}, value: nil, dirty: true},
		{key: []byte{0x01}, value: nil, dirty: true},
		{key: []byte{0x00, 0x00, 0xf, 0x3}, value: nil, dirty: true},
		{key: byteArray(62), value: nil, dirty: true},
		{key: byteArray(63), value: nil, dirty: true},
		{key: byteArray(64), value: []byte{0x01}, dirty: true},
		{key: byteArray(318), value: []byte{0x01}, dirty: true},
		{key: byteArray(573), value: []byte{0x01}, dirty: true},
	}

	buffer := bytes.NewBuffer(nil)

	for _, test := range tests {
		err := encodeLeaf(test, buffer)
		require.NoError(t, err)

		res := new(Leaf)
		err = res.Decode(buffer, 0)
		require.NoError(t, err)

		res.hash = nil
		test.encoding = nil
		require.Equal(t, test, res)
	}
}

func TestDecode(t *testing.T) {
	tests := []Node{
		&Branch{key: []byte{}, children: [16]Node{}, value: nil},
		&Branch{key: []byte{0x00}, children: [16]Node{}, value: nil},
		&Branch{key: []byte{0x00, 0x00, 0xf, 0x3}, children: [16]Node{}, value: nil},
		&Branch{key: []byte{}, children: [16]Node{}, value: []byte{0x01}},
		&Branch{key: []byte{}, children: [16]Node{&Leaf{}}, value: []byte{0x01}},
		&Branch{key: []byte{}, children: [16]Node{&Leaf{}, nil, &Leaf{}}, value: []byte{0x01}},
		&Branch{
			key: []byte{},
			children: [16]Node{
				&Leaf{}, nil, &Leaf{}, nil,
				nil, nil, nil, nil,
				nil, &Leaf{}, nil, &Leaf{}},
			value: []byte{0x01},
		},
		&Leaf{key: []byte{}, value: nil},
		&Leaf{key: []byte{0x00}, value: nil},
		&Leaf{key: []byte{0x00, 0x00, 0xf, 0x3}, value: nil},
		&Leaf{key: byteArray(62), value: nil},
		&Leaf{key: byteArray(63), value: nil},
		&Leaf{key: byteArray(64), value: []byte{0x01}},
		&Leaf{key: byteArray(318), value: []byte{0x01}},
		&Leaf{key: byteArray(573), value: []byte{0x01}},
	}

	buffer := bytes.NewBuffer(nil)
	const parallel = false

	for _, test := range tests {
		err := encodeNode(test, buffer, parallel)
		require.NoError(t, err)

		res, err := decode(buffer)
		require.NoError(t, err)

		switch n := test.(type) {
		case *Branch:
			require.Equal(t, n.key, res.(*Branch).key)
			require.Equal(t, n.childrenBitmap(), res.(*Branch).childrenBitmap())
			require.Equal(t, n.value, res.(*Branch).value)
		case *Leaf:
			require.Equal(t, n.key, res.(*Leaf).key)
			require.Equal(t, n.value, res.(*Leaf).value)
		default:
			t.Fatal("unexpected node")
		}
	}
}
