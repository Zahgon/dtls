// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
	"github.com/pion/dtls/v3/pkg/crypto/signaturehash"
)

/*
MessageCertificateRequest is so a non-anonymous server can optionally
request a certificate from the client, if appropriate for the selected cipher
suite.  This message, if sent, will immediately follow the ServerKeyExchange
message (if it is sent; otherwise, this message follows the
server's Certificate message).

https://tools.ietf.org/html/rfc5246#section-7.4.4
*/
type MessageCertificateRequest struct {
	CertificateTypes            []clientcertificate.Type
	SignatureHashAlgorithms     []signaturehash.Algorithm
	CertificateAuthoritiesNames [][]byte
}

const (
	messageCertificateRequestMinLength = 5
)

// Type returns the Handshake Type.
func (m MessageCertificateRequest) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Marshal encodes the Handshake.
func (m *MessageCertificateRequest) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115: certificate types count is validated to be <= 255 above.

//nolint:gosec //G115

// Distinguished Names

//nolint:gosec //G115

//nolint:gosec //G115

// Unmarshal populates the message from encoded data.
func (m *MessageCertificateRequest) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
