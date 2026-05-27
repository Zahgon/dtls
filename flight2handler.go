// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
)

func flight2Parse(
	ctx context.Context,
	c flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return *new(flightVal), nil, nil
}

// Client may retransmit the first ClientHello when HelloVerifyRequest is dropped.
// Parse as flight 0 in this case.

// Validate type

func flight2Generate(
	_ flightConn,
	state *State,
	_ *handshakeCache,
	_ *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
