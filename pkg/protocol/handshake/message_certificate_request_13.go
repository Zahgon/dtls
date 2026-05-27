// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/protocol/extension"
)

// MessageCertificateRequest13 represents the CertificateRequest handshake message for DTLS 1.3.
// This message is used by the server to request a certificate from the client.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.3.2
type MessageCertificateRequest13 struct {
	// CertificateRequestContext is an opaque value that the server creates
	// to bind the client's certificate to the handshake context.
	CertificateRequestContext []byte

	// Extensions contains the list of extensions.
	// The signature_algorithms extension is REQUIRED per RFC 8446.
	Extensions []extension.Extension
}

// Type returns the handshake message type.
func (m MessageCertificateRequest13) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

const (
	maxUint16                 = 0xffff
	certReq13ContextMaxLength = 255
	certReq13MinLength        = 3
)

// Marshal encodes the MessageCertificateRequest13 into its wire format.
//
// Wire format:
//
//	[1 byte]  certificate_request_context length
//	[0-255]   certificate_request_context data
//	[2 bytes] extensions length (from extension.Marshal)
//	[variable] extensions data
func (m *MessageCertificateRequest13) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	// Validate certificate_request_context length
	return nil, nil
}

// Validate that signature_algorithms extension is present (required by RFC 8446)

// Add certificate_request_context (1-byte length prefix)

// Marshal extensions (includes 2-byte length prefix, like in TLS 1.2)

// Validate extensions length is in valid range <2..2^16-1>

// Unmarshal decodes the MessageCertificateRequest13 from its wire format.
func (m *MessageCertificateRequest13) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	// Validate minimum data length
	return nil
}

// Read certificate_request_context

// Read extensions length (2 bytes)

// Validate we have exactly extensionsLen bytes remaining after the length field

// Validate that signature_algorithms extension is present (required by RFC 8446)
