// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/crypto/hash"
	"github.com/pion/dtls/v3/pkg/crypto/signature"
)

// MessageCertificateVerify provide explicit verification of a
// client certificate.
//
// https://tools.ietf.org/html/rfc5246#section-7.4.8
type MessageCertificateVerify struct {
	HashAlgorithm      hash.Algorithm
	SignatureAlgorithm signature.Algorithm
	Signature          []byte
}

const handshakeMessageCertificateVerifyMinLength = 4

// Type returns the Handshake Type.
func (m MessageCertificateVerify) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Marshal encodes the Handshake.
func (m *MessageCertificateVerify) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CertificateVerify in DTLS 1.2 encodes hash/signature as 1 byte each.

//nolint:gosec // G115

// Unmarshal populates the message from encoded data.
func (m *MessageCertificateVerify) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
