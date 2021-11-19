# Gossamer `network` Package

This package emulates the [peer-to-peer networking capabilities](https://crates.parity.io/sc_network/index.html)
provided by the [Substrate](https://docs.substrate.io/) framework for blockchain development, which implies that it is
built on the extensible [`libp2p` networking stack](https://docs.libp2p.io/introduction/what-is-libp2p/). `libp2p`
provides implementations of a number of battle-tested peer-to-peer (P2P) networking protocols (e.g. [Noise](#noise) for
[key exchange](#identities--key-management), and [Yamux](#yamux) for [multiplexing](#multiplexing)), and also makes it
possible to implement the blockchain-specific protocols defined by Substrate (e.g. authoring and finalising blocks, and
maintaining the [transaction pool](https://docs.substrate.io/v3/concepts/tx-pool/)). The purpose of this document is to
provide the information that is needed to understand the P2P networking capabilities that are implemented by Gossamer -
this includes an introduction to P2P networks and `libp2p`, as well as detailed descriptions of the Gossamer P2P
networking protocols.

## Peer-to-Peer Networking & `libp2p`

[Peer-to-peer](https://en.wikipedia.org/wiki/Peer-to-peer) networking has been a dynamic field of research for over two
decades, and P2P protocols are at the heart of many blockchain networks. P2P networks can be contrasted with traditional
[client-server](https://en.wikipedia.org/wiki/Client%E2%80%93server_model) networks where there is a clear separation of
authority and privilege between the maintainers of the network and its users - in a P2P network, each participant
possesses equal authority and equal privilege. `libp2p` is a framework for implementing P2P networks that was
modularized out of [IPFS](https://ipfs.io/); there are implementations in many languages including Go (used by this
project), Rust, Javascript, C++, and more. In addition to the standard library of protocols in a `libp2p`
implementation, there is a rich ecosystem of P2P networking packages that work with the pluggable architecture of
`libp2p`. In some cases, Gossamer uses the `libp2p` networking primitives to implement custom protocols for
blockchain-specific use cases. What follows is an exploration into three concepts that underpin P2P networks: identity &
key management, peer discovery & management, and multiplexing.

### Identity & Key Management

Many peer-to-peer networks, including those built with Gossamer, use
[public-key cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography) (also known as asymmetric cryptography)
to allow network participants to securely identify themselves and interact with one another. The term "asymmetric"
refers to the fact that in a public-key cryptography system, each participant's identity is associated with a set of two
keys, each of which serve a distinct ("asymmetric") purpose. One of the keys in an asymmetric key pair is public, this
is the key that the participant uses to identify themselves; the other pair is private and is used by the network
participant to "sign" messages in order to cryptographically prove that the message originated from the private key's
owner. It may be constructive to think about a public key as a username and private key as a password, such as for a
banking or social media website. Participants in P2P networks that use asymmetric cryptography must protect their
private keys, as well as maintain indices of the public keys that belong to the other participants in the network.
Gossamer provides a [keystore](../../lib/keystore) for securely storing private keys. There are a number of Gossamer
processes that manage the public keys of network peers - some of these, such as
[peer discovery and management](#peer-discovery--management), are described in this document, but there are other
packages (most notably [`peerset`](../peerset)) that also interact with the public keys of network peers. One of the
most critical details in a network that uses asymmetric cryptography is the
[key distribution](https://en.wikipedia.org/wiki/Key_distribution) mechanism, which is the process that the hosts in the
network use to securely exchange public keys - `libp2p` supports [Noise](#noise), a key distribution framework that is
based on [Diffie-Hellman key exchange](https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange).

### Peer Discovery & Management

In a peer-to-peer network, "[discovery](https://docs.libp2p.io/concepts/publish-subscribe/#discovery)" is the term that
is used to describe the mechanism that peers use to find one another - this is an important topic since there is not a
privileged authority that can maintain an index of known/trusted network participants. Gossamer uses
[Kademlia](#kademlia) for peer discovery.

### Multiplexing

[Multiplexing](<(https://docs.libp2p.io/concepts/stream-multiplexing/)>) allows multiple independent logical streams to
all share a common underlying transport medium, which amortizes the overhead of establishing new connections with
network peers. Gossamer uses [Yamux](#yamux) for multiplexing.

## Gossamer Network Protocols

### `libp2p`-Integrated Protocols

#### `ping`

This is a simple liveness check [protocol](https://docs.libp2p.io/concepts/protocols/#ping) that peers can use to
quickly see if another peer is online - it is
[included](https://github.com/libp2p/go-libp2p/tree/master/p2p/protocol/ping) with the official Go implementation of
`libp2p`.

#### `identify`

The [`identify` protocol](https://docs.libp2p.io/concepts/protocols/#identify) allows peers to exchange information
about each other, most notably their public keys and known network addresses; like [`ping`](#ping), it is
[included with `go-libp2p`](https://github.com/libp2p/go-libp2p/tree/master/p2p/protocol/identify).

#### Noise

As described above, [Noise](http://noiseprotocol.org/) provides `libp2p` with its
[key distribution](#identity--key-management) capabilities. The Noise protocol is
[well documented](http://cryptowiki.net/index.php?title=Noise_Protocol_Framework) and the Go implementation is
maintained [under the official](https://github.com/libp2p/go-libp2p-noise) `libp2p` GitHub organization. Noise defines a
number of [variables](http://cryptowiki.net/index.php?title=Noise_Protocol_Framework#Noise_Variables) and
[handshake patterns](http://cryptowiki.net/index.php?title=Noise_Protocol_Framework#Handshake_patterns) that
participants in a peer-to-peer network can use to establish message-passing channels with one another.

#### Yamux

Learn more about [Yamux](https://docs.libp2p.io/concepts/stream-multiplexing/#yamux).

#### Kademlia

Learn more about [Kademlia](https://en.wikipedia.org/wiki/Kademlia).

### Blockchain-Specific Protocols

##### Notification Protocols

Learn more about [notification protocols](https://crates.parity.io/sc_network/index.html#notifications-protocols).

###### Transactions

###### Block Announces

##### Request/Response

Learn more about
[request/response protocols](https://crates.parity.io/sc_network/index.html#request-response-protocols).

###### Sync

###### GRANDPA

###### Light
