// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
)

func flight1Parse(
	ctx context.Context,
	conn flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	// HelloVerifyRequest can be skipped by the server,
	// so allow ServerHello during flight1 also
	return *new(flightVal), nil, nil
}

// No valid message received. Keep reading

// Flight1 and flight2 were skipped.
// Parse as flight3.

// DTLS 1.2 clients must not assume that the server will use the protocol version
// specified in HelloVerifyRequest message. RFC 6347 Section 4.2.1

//nolint:cyclop
func flight1Generate(
	conn flightConn,
	state *State,
	_ *handshakeCache,
	cfg *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If we have a connection ID generator, use it. The CID may be zero length,
// in which case we are just requesting that the server send us a CID to
// use.

// The presence of a generator indicates support for connection IDs. We
// use the presence of a non-nil local CID in flight 3 to determine
// whether we send a CID in the second ClientHello, so we convert any
// nil CID returned by a generator to []byte{}.
