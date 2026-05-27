// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
)

// renegotiationInfoSCSV is TLS_EMPTY_RENEGOTIATION_INFO_SCSV defined in RFC 5746.
// https://datatracker.ietf.org/doc/html/rfc5746#section-3.3.
const renegotiationInfoSCSV uint16 = 0x00ff

//nolint:cyclop,gocognit
func flight0Parse(
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

// Connection Identifiers must be negotiated afresh on session resumption.
// https://datatracker.ietf.org/doc/html/rfc9146#name-the-connection_id-extension

// Validate type

// remote server name

// Only set connection ID to be sent if server supports connection
// IDs.

// Store the client's certificate signature schemes for later validation

// If the client doesn't support connection IDs, the server should not
// expect one to be sent.

func handleHelloResume(
	sessionID []byte,
	state *State,
	cfg *handshakeConfig,
	next flightVal,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return *new(flightVal), nil, nil
}

func flight0Generate(
	_ flightConn,
	state *State,
	_ *handshakeCache,
	cfg *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	_ = "STUB: not implemented"
	// Initialize
	return nil, nil, nil
}
