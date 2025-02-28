// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	network "github.com/ChainSafe/gossamer/dot/network"
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p-core/peer"

	protocol "github.com/libp2p/go-libp2p-core/protocol"
)

// Network is an autogenerated mock type for the Network type
type Network struct {
	mock.Mock
}

// GossipMessage provides a mock function with given fields: msg
func (_m *Network) GossipMessage(msg network.NotificationsMessage) {
	_m.Called(msg)
}

// RegisterNotificationsProtocol provides a mock function with given fields: sub, messageID, handshakeGetter, handshakeDecoder, handshakeValidator, messageDecoder, messageHandler, batchHandler
func (_m *Network) RegisterNotificationsProtocol(sub protocol.ID, messageID byte, handshakeGetter func() (network.Handshake, error), handshakeDecoder func([]byte) (network.Handshake, error), handshakeValidator func(peer.ID, network.Handshake) error, messageDecoder func([]byte) (network.NotificationsMessage, error), messageHandler func(peer.ID, network.NotificationsMessage) (bool, error), batchHandler func(peer.ID, network.NotificationsMessage)) error {
	ret := _m.Called(sub, messageID, handshakeGetter, handshakeDecoder, handshakeValidator, messageDecoder, messageHandler, batchHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(protocol.ID, byte, func() (network.Handshake, error), func([]byte) (network.Handshake, error), func(peer.ID, network.Handshake) error, func([]byte) (network.NotificationsMessage, error), func(peer.ID, network.NotificationsMessage) (bool, error), func(peer.ID, network.NotificationsMessage)) error); ok {
		r0 = rf(sub, messageID, handshakeGetter, handshakeDecoder, handshakeValidator, messageDecoder, messageHandler, batchHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessage provides a mock function with given fields: to, msg
func (_m *Network) SendMessage(to peer.ID, msg network.NotificationsMessage) error {
	ret := _m.Called(to, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(peer.ID, network.NotificationsMessage) error); ok {
		r0 = rf(to, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
