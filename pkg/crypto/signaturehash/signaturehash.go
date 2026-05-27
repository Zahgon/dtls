// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package signaturehash provides the SignatureHashAlgorithm as defined in TLS 1.2
package signaturehash

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"

	"github.com/pion/dtls/v3/pkg/crypto/hash"
	"github.com/pion/dtls/v3/pkg/crypto/signature"
)

// Algorithm is a signature/hash algorithm pairs which may be used in
// digital signatures.
//
// https://tools.ietf.org/html/rfc5246#section-7.4.1.4.1
type Algorithm struct {
	Hash      hash.Algorithm
	Signature signature.Algorithm
}

// Algorithms returns signature algorithms compatible with DTLS 1.2 / TLS 1.2.
//
// IMPORTANT: order in this slice determines priority used by SelectSignatureScheme.
//
// Order follows industry standard preference (ECDSA-first) as used by OpenSSL,
// BoringSSL, Firefox, Chrome, and other major TLS implementations.
func Algorithms() []Algorithm {
	_ = "STUB: not implemented"

	// ECDSA schemes (modern, efficient - industry standard preference)
	return nil
}

// Ed25519

// RSA-PSS RSAE schemes (DTLS 1.3 compatible with standard RSA certs)
// Note: We only offer RSA_PSS_RSAE variants (0x0804-0x0806), not RSA_PSS_PSS
// (0x0809-0x080b). RSA-PSS certificates with OID id-RSASSA-PSS are virtually
// unused in the real world and are not allowed by the CA/Browser Forum Baseline
// Requirements for WebPKI. We avoid unnecessary complexity for certificates that
// don't exist in practice, following the pragmatic approach of Go's crypto/tls
// and BoringSSL: target real-world WebPKI use cases rather than RFC completeness.
// RSA_PSS_PSS schemes are parsed for wire-format compatibility but never negotiated.

// {hash.SHA256, signature.RSA_PSS_PSS_SHA256},
// {hash.SHA384, signature.RSA_PSS_PSS_SHA384},
// {hash.SHA512, signature.RSA_PSS_PSS_SHA512},

// RSA PKCS#1 v1.5 schemes (legacy, DTLS 1.2)

// SelectSignatureScheme returns most preferred and compatible scheme for DTLS <= 1.2.
func SelectSignatureScheme(sigs []Algorithm, privateKey crypto.PrivateKey) (Algorithm, error) {
	_ = "STUB: not implemented"
	return *new(Algorithm), nil
}

// SelectSignatureScheme13 returns most preferred and compatible scheme for DTLS 1.3.
func SelectSignatureScheme13(sigs []Algorithm, privateKey crypto.PrivateKey) (Algorithm, error) {
	_ = "STUB: not implemented"
	return *new(Algorithm), nil
}

// isCompatible checks that given private key is compatible with the signature scheme.
func (a *Algorithm) isCompatible(signer crypto.Signer) bool {
	_ = "STUB: not implemented"
	return false
}

// RSA keys are compatible with both PKCS#1 v1.5 and PSS signatures

// ParseSignatureSchemes translates []tls.SignatureScheme to []signatureHashAlgorithm.
// It returns default signature scheme list if no SignatureScheme is passed.
// This function handles both TLS 1.2 byte-split encoding and TLS 1.3 PSS full uint16 schemes.
//
// For DTLS 1.2 / TLS 1.2, this returns Algorithms() which excludes TLS 1.3-specific
// schemes like RSA-PSS for compatibility with implementations like OpenSSL.
// When DTLS 1.3 is implemented, use Algorithms13() or create ParseSignatureSchemes13().
func ParseSignatureSchemes(sigs []tls.SignatureScheme, insecureHashes bool) ([]Algorithm, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FromCertificate maps x509.SignatureAlgorithm to the corresponding Algorithm type.
func FromCertificate(cert *x509.Certificate) (Algorithm, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return *new(Algorithm), nil
}

// Ed25519 doesn't use a separate hash

// Unmarshal translates a tls.SignatureScheme to a Algorithm.
// This function handles both TLS 1.2 byte-split encoding and
// TLS 1.3 PSS full uint16 schemes.
func (a *Algorithm) Unmarshal(sigScheme tls.SignatureScheme) error {
	_ = "STUB: not implemented"
	return nil
}

// TLS 1.3 PSS scheme - full uint16 is the signature algorithm

// TLS 1.2 style - split into hash (high byte) and signature (low byte)

// Validate signature algorithm

// Validate hash algorithm

// Marshal encodes the Algorithm to the correct TLS 1.2 byte-split or
// the TLS 1.3 PSS full uint16 schemes.
func (a *Algorithm) Marshal() []byte { _ = "STUB: not implemented"; return nil }

// For PSS schemes, write the full uint16 SignatureScheme value.
// For other schemes, write hash (high byte) + signature (low byte) in TLS 1.2 style.

// TLS 1.3 PSS: full uint16 is the signature scheme

// TLS 1.2 style: hash byte + signature byte
//nolint:gosec // G115: TLS 1.2 hash algorithm field is defined as 1 byte.
//nolint:gosec // G115: TLS 1.2 signature algorithm field is defined as 1 byte.
