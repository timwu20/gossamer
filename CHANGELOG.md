# Semantic Versioning Changelog

# [0.5.0](https://github.com/timwu20/gossamer/compare/v0.4.1...v0.5.0) (2021-12-06)


### Bug Fixes

* **babe:** Fix extrinsic format in block. ([#1530](https://github.com/timwu20/gossamer/issues/1530)) ([1a03b2a](https://github.com/timwu20/gossamer/commit/1a03b2a08a191b37ac620ae8ff5c47dccf92b421))
* **dot/network:** Fix notification handshake and reuse stream. ([#1545](https://github.com/timwu20/gossamer/issues/1545)) ([a632dc4](https://github.com/timwu20/gossamer/commit/a632dc444ed62896371806b3b20ffe0e50ecc7d0))
* **dot/network:** split stored streams and handshakeData into inbound and outbound ([#1553](https://github.com/timwu20/gossamer/issues/1553)) ([637050b](https://github.com/timwu20/gossamer/commit/637050bb94e12eb518ef93122c5ef323c27a53fb))
* **lib/babe:** fix BABE state storing after building block ([#1536](https://github.com/timwu20/gossamer/issues/1536)) ([1a3dea2](https://github.com/timwu20/gossamer/commit/1a3dea29c419c67b5a1d8c66c96e56c4da0cdd31))
* persist node name ([#1543](https://github.com/timwu20/gossamer/issues/1543)) ([88b88f2](https://github.com/timwu20/gossamer/commit/88b88f2c889aa27506dc482eadda5a295b2626a4))
* **release:** Trigger release when pushed to main branch ([#1566](https://github.com/timwu20/gossamer/issues/1566)) ([4f2ba56](https://github.com/timwu20/gossamer/commit/4f2ba56b476962518d6b46856265c7b32074fd34))
* update go-schnorrkel version ([#1557](https://github.com/timwu20/gossamer/issues/1557)) ([b86c7ff](https://github.com/timwu20/gossamer/commit/b86c7ff8507c097879996f0d5eb9ec84628127f9))


### Features

* Add properties and chainId on build-spec command ([#1520](https://github.com/timwu20/gossamer/issues/1520)) ([b18290c](https://github.com/timwu20/gossamer/commit/b18290cb269d4f1a8e2915dac6add5c9363a2a25))
* **dot/network, lib/grandpa:** request justification on receiving NeighbourMessage, verify justification on receipt ([#1529](https://github.com/timwu20/gossamer/issues/1529)) ([e1f9f42](https://github.com/timwu20/gossamer/commit/e1f9f427c47255ca5cdcdd443ed8ebcf2451c759))
* **dot/network:** add propagate return bool to messageHandler func type to determine whether to propagate message or not ([#1555](https://github.com/timwu20/gossamer/issues/1555)) ([0d6f488](https://github.com/timwu20/gossamer/commit/0d6f48834327fa453977c7ccde2ad4ac99588e75))
* **lib/grandpa:** fully verify justifications using GrandpaState ([#1544](https://github.com/timwu20/gossamer/issues/1544)) ([028d25e](https://github.com/timwu20/gossamer/commit/028d25e72e5a3c32380f929d0db11ac4e874af37))
* **lib/grandpa:** send NeighbourMessage to peers ([#1558](https://github.com/timwu20/gossamer/issues/1558)) ([322ccf9](https://github.com/timwu20/gossamer/commit/322ccf9cf5fe18cc83ec05a5490ec250bd028f14))
