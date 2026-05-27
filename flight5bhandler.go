// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
)

func flight5bParse(
	_ context.Context,
	_ flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return *new(flightVal), nil, nil
}

// No valid message received. Keep reading

// Other party may re-transmit the last flight. Keep state to be flight5b.

func flight5bGenerate(
	_ flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	_ = "STUB: not implemented" //nolint:gocognit
	return nil, nil, nil
}
