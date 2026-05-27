// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package hash provides TLS HashAlgorithm as defined in TLS 1.2
package hash

import ( //nolint:gci
	"crypto"
	//nolint:gosec
)

// Algorithm is used to indicate the hash algorithm used
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-18
type Algorithm uint16

// Supported hash algorithms.
const (
	None    Algorithm = 0 // Blacklisted
	MD5     Algorithm = 1 // Blacklisted
	SHA1    Algorithm = 2 // Blacklisted
	SHA224  Algorithm = 3
	SHA256  Algorithm = 4
	SHA384  Algorithm = 5
	SHA512  Algorithm = 6
	Ed25519 Algorithm = 8
)

// String makes hashAlgorithm printable.
func (a Algorithm) String() string { _ = "STUB: not implemented"; return "" }

// [RFC3279]

// [RFC3279]

// [RFC4055]

// [RFC4055]

// [RFC4055]

// [RFC4055]

// Digest performs a digest on the passed value.
func (a Algorithm) Digest(b []byte) []byte { _ = "STUB: not implemented"; return nil }

// #nosec

// #nosec

// Insecure returns if the given HashAlgorithm is considered secure in DTLS 1.2
// .
func (a Algorithm) Insecure() bool { _ = "STUB: not implemented"; return false }

// CryptoHash returns the crypto.Hash implementation for the given HashAlgorithm.
func (a Algorithm) CryptoHash() crypto.Hash { _ = "STUB: not implemented"; return *new(crypto.Hash) }

// Algorithms returns all the supported Hash Algorithms.
func Algorithms() map[Algorithm]struct{} { _ = "STUB: not implemented"; return nil }

// ExtractHashFromPSS extracts the hash algorithm from an RSA-PSS SignatureScheme value.
// This handles TLS 1.3 PSS schemes.
// Returns None if the scheme is not a recognized PSS scheme.
func ExtractHashFromPSS(pssScheme uint16) Algorithm {
	_ = "STUB: not implemented"
	// Note: We can't import signature package here due to circular dependency,
	// so we use the raw values. These correspond to:
	// 0x0804 = RSA_PSS_RSAE_SHA256, 0x0809 = RSA_PSS_PSS_SHA256
	// 0x0805 = RSA_PSS_RSAE_SHA384, 0x080a = RSA_PSS_PSS_SHA384
	// 0x0806 = RSA_PSS_RSAE_SHA512, 0x080b = RSA_PSS_PSS_SHA512
	return *new(Algorithm)
}
