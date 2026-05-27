// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/protocol/extension"
	"golang.org/x/crypto/cryptobyte"
)

// CertificateEntry13 represents a single certificate entry in the DTLS 1.3 Certificate message.
// Each entry contains certificate data and optional per-certificate extensions.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.4.2
type CertificateEntry13 struct {
	// CertificateData contains the DER-encoded X.509 certificate.
	// Can be empty for certain contexts (e.g., RawPublicKey mode).
	CertificateData []byte

	// Extensions contains per-certificate extensions.
	// Examples: OCSP status, SignedCertificateTimestamp, etc.
	Extensions []extension.Extension
}

// MessageCertificate13 represents the Certificate handshake message for DTLS 1.3.
// This message is used to transport the certificate chain and associated extensions.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.4.2
type MessageCertificate13 struct {
	// CertificateRequestContext is an opaque value that binds this certificate
	// to a specific CertificateRequest (for client certificates) or is empty
	// for server certificates.
	CertificateRequestContext []byte

	// CertificateList contains the certificate chain with each entry having
	// optional per-certificate extensions.
	CertificateList []CertificateEntry13
}

// Type returns the handshake message type.
func (m MessageCertificate13) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

const (
	maxUint24                    = 0xffffff
	cert13ContextLengthFieldSize = 1
	cert13ContextMaxLength       = 255
	cert13CertLengthFieldSize    = 3
	cert13ExtLengthFieldSize     = 2
)

// Marshal encodes the MessageCertificate13 into its wire format.
//
// Wire format:
//
//	[1 byte]  certificate_request_context length
//	[0-255]   certificate_request_context data
//	[3 bytes] certificate_list length
//	For each certificate:
//	  [3 bytes]  cert_data length
//	  [variable] cert_data (DER certificate)
//	  [2 bytes]  extensions length (from extension.Marshal)
//	  [variable] extensions data
func (m *MessageCertificate13) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	// Validate certificate_request_context length
	return nil, nil
}

// Start with certificate_request_context (1-byte length prefix)
//nolint:gosec // G115: certificate_request_context length is validated to be <= 255 above.

// Build certificate_list

// Add cert_data as a 3-byte length prefix

//nolint:gosec // G115

// Marshal extensions (includes a 2-byte length prefix)

// Check size of certificate_list is still within bounds

// Add certificate_list with 3-byte length prefix

//nolint:gosec // G115

// parseCertificate13Entry parses a single certificate entry from the cryptobyte string.
func parseCertificate13Entry(str *cryptobyte.String) (*CertificateEntry13, error) {
	_ = "STUB: not implemented"
	// Read cert_data with 3-byte length prefix
	return nil, nil
}

// Validate cert_data length is in valid range <1..2^24-1>

// Copy cert_data to avoid aliasing issues

// Validate extensions length (2-byte length prefix + up to 2^16-1 bytes of data)

// Read extensions length to validate we have enough data

// Unmarshal extensions data

// Advance the cryptobyte.String's position

// Unmarshal decodes the MessageCertificate13 from its wire format.
func (m *MessageCertificate13) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	// Validate minimum data length
	return nil
}

// Read certificate_request_context with 1-byte length prefix

// Read certificate_list with 3-byte length prefix

// Ensure no trailing data

// Parse certificate_list
