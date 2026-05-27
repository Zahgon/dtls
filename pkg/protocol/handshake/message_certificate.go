// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

// MessageCertificate is a DTLS Handshake Message
// it can contain either a Client or Server Certificate
//
// https://tools.ietf.org/html/rfc5246#section-7.4.2
type MessageCertificate struct {
	Certificate [][]byte
}

// Type returns the Handshake Type.
func (m MessageCertificate) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

const (
	handshakeMessageCertificateLengthFieldSize = 3
)

// Marshal encodes the Handshake.
func (m *MessageCertificate) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Total Payload Size
//nolint:gosec // G115

// Certificate Length
//nolint:gosec // G115

// Certificate body

// Unmarshal populates the message from encoded data.
func (m *MessageCertificate) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
