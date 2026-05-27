// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/extension"
)

/*
MessageClientHello is for when a client first connects to a server it is
required to send the client hello as its first message.  The client can also send a
client hello in response to a hello request or on its own
initiative in order to renegotiate the security parameters in an
existing connection.
*/
type MessageClientHello struct {
	Version protocol.Version
	Random  Random
	Cookie  []byte

	SessionID []byte

	CipherSuiteIDs     []uint16
	CompressionMethods []*protocol.CompressionMethod
	Extensions         []extension.Extension
}

const handshakeMessageClientHelloVariableWidthStart = 34

// Type returns the Handshake Type.
func (m MessageClientHello) Type() Type {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the Handshake.
	new(Type)
}

func (m *MessageClientHello) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115: session ID length is validated to be <= 255 above.

//nolint:gosec // G115: cookie length is validated to be <= 255 above.

// Unmarshal populates the message from encoded data.
func (m *MessageClientHello) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// rest of packet has variable width sections

// Cipher Suites

// Compression Methods

// Extensions
