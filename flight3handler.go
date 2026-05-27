// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
)

//nolint:gocognit,gocyclo,maintidx,cyclop
func flight3Parse(
	ctx context.Context,
	conn flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	// Clients may receive multiple HelloVerifyRequest messages with different cookies.
	// Clients SHOULD handle this by sending a new ClientHello with a cookie in response
	// to the new HelloVerifyRequest. RFC 6347 Section 4.2.1
	return *new(flightVal), nil, nil
}

// DTLS 1.2 clients must not assume that the server will use the protocol version
// specified in HelloVerifyRequest message. RFC 6347 Section 4.2.1

// Don't have enough messages. Keep reading

//nolint:nestif

// This should be exactly 1, the zero case is handle when unmarshalling

// Meh, internal error?

// Only set connection ID to be sent if client supports connection
// IDs.

// If the server doesn't support connection IDs, the client should not
// expect one to be sent.

// Don't have enough messages. Keep reading

func handleResumption(
	ctx context.Context,
	c flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return *new(flightVal), nil, nil
}

// Now, encrypted packets can be handled

// No valid message received. Keep reading

//nolint:cyclop
func handleServerKeyExchange(
	_ flightConn,
	state *State,
	cfg *handshakeConfig,
	keyExchangeMessage *handshake.MessageServerKeyExchange,
) (*alert.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif

//nolint:nilnil

func flight3Generate(
	_ flightConn,
	state *State,
	_ *handshakeCache,
	cfg *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If we sent a connection ID on the first ClientHello, send it on the
// second.
