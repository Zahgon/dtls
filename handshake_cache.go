// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"sync"

	"github.com/pion/dtls/v3/pkg/crypto/prf"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
)

type handshakeCacheItem struct {
	typ             handshake.Type
	isClient        bool
	epoch           uint16
	messageSequence uint16
	data            []byte
}

type handshakeCachePullRule struct {
	typ      handshake.Type
	epoch    uint16
	isClient bool
	optional bool
}

type handshakeCache struct {
	cache []*handshakeCacheItem
	mu    sync.Mutex
}

func newHandshakeCache() *handshakeCache { _ = "STUB: not implemented"; return nil }

func (h *handshakeCache) push(data []byte, epoch, messageSequence uint16, typ handshake.Type, isClient bool) {
	_ = "STUB: not implemented"
	return
}

// returns a list handshakes that match the requested rules
// the list will contain null entries for rules that can't be satisfied
// multiple entries may match a rule, but only the last match is returned (ie ClientHello with cookies).
func (h *handshakeCache) pull(rules ...handshakeCachePullRule) []*handshakeCacheItem {
	_ = "STUB: not implemented"
	return nil
}

// fullPullMap pulls all handshakes between rules[0] to rules[len(rules)-1] as map.
//
//nolint:cyclop
func (h *handshakeCache) fullPullMap(
	startSeq int,
	cipherSuite CipherSuite,
	rules ...handshakeCachePullRule,
) (int, map[handshake.Type]handshake.Message, bool) {
	_ = "STUB: not implemented"
	return 0, nil, false
}

// Missing mandatory message.

//nolint:gosec // G115
// There is a gap. Some messages are not arrived.

// pullAndMerge calls pull and then merges the results, ignoring any null entries.
func (h *handshakeCache) pullAndMerge(rules ...handshakeCachePullRule) []byte {
	_ = "STUB: not implemented"
	return nil
}

// sessionHash returns the session hash for Extended Master Secret support
// https://tools.ietf.org/html/draft-ietf-tls-session-hash-06#section-4
func (h *handshakeCache) sessionHash(hf prf.HashFunc, epoch uint16, additional ...[]byte) ([]byte, error) {
	_ = "STUB: not implemented"

	// Order defined by https://tools.ietf.org/html/rfc5246#section-7.3
	return nil, nil
}
