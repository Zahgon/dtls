// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/extension"
)

// MessageServerHello is sent in response to a ClientHello
// message when it was able to find an acceptable set of algorithms.
// If it cannot find such a match, it will respond with a handshake
// failure alert.
//
// https://tools.ietf.org/html/rfc5246#section-7.4.1.3
type MessageServerHello struct {
	Version protocol.Version
	Random  Random

	SessionID []byte

	CipherSuiteID     *uint16
	CompressionMethod *protocol.CompressionMethod
	Extensions        []extension.Extension
}

const messageServerHelloVariableWidthStart = 2 + RandomLength

// Type returns the Handshake Type.
func (m MessageServerHello) Type() Type {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the Handshake.
	new(Type)
}

func (m *MessageServerHello) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115: session ID length is validated to be <= 255 above.

// Unmarshal populates the message from encoded data.
func (m *MessageServerHello) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
