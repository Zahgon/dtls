// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
)

func flight5Parse(
	_ context.Context,
	conn flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return *new(flightVal), nil, nil
}

// No valid message received. Keep reading

//nolint:gocognit,cyclop,maintidx
func flight5Generate(
	conn flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//nolint:nestif

// handshakeMessageServerKeyExchange is optional for PSK

// Append not-yet-sent packets

//nolint:gosec // G115

// If the client has sent a certificate with signing ability, a digitally-signed
// CertificateVerify message is sent to explicitly verify possession of the
// private key in the certificate.

// Find compatible signature scheme

// seqPred++ // this is the last use of seqPred

//nolint:gocognit,cyclop
func initializeCipherSuite(
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
	handshakeKeyExchange *handshake.MessageServerKeyExchange,
	sendingPlainText []byte,
) (*alert.Alert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint

//nolint:nestif
// Verify that the pair of hash algorithm and signiture is listed.

//nolint
