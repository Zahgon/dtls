// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension //nolint:dupl

const (
	postHandshakeAuthHeaderSize = 4
)

// PostHandshakeAuth defines a DTLS 1.3 extension that is used to indicate
// that a client is willing to perform post-handshake authentication.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.6
type PostHandshakeAuth struct {
	Enabled bool
}

// TypeValue returns the extension TypeValue.
func (p PostHandshakeAuth) TypeValue() TypeValue { _ = "STUB: not implemented"; return *new(TypeValue) }

// Marshal encodes the extension.
func (p *PostHandshakeAuth) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the extension from encoded data.
func (p *PostHandshakeAuth) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
