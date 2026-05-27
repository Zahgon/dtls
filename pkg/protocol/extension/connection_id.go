// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// ConnectionID is a DTLS extension that provides an alternative to IP address
// and port for session association.
//
// https://tools.ietf.org/html/rfc9146
type ConnectionID struct {
	// A zero-length connection ID indicates for a client or server that
	// negotiated connection IDs from the peer will be sent but there is no need
	// to respond with one
	CID []byte // variable length
}

// TypeValue returns the extension TypeValue.
func (c ConnectionID) TypeValue() TypeValue { _ = "STUB: not implemented"; return *new(TypeValue) }

// Marshal encodes the extension.
func (c *ConnectionID) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the extension from encoded data.
func (c *ConnectionID) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
