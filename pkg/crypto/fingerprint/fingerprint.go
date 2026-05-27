// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package fingerprint provides a helper to create fingerprint string from certificate
package fingerprint

import (
	"crypto"
	"crypto/x509"
	"errors"
)

var (
	errHashUnavailable          = errors.New("fingerprint: hash algorithm is not linked into the binary")
	errInvalidFingerprintLength = errors.New("fingerprint: invalid fingerprint length")
)

// Fingerprint creates a fingerprint for a certificate using the specified hash algorithm.
func Fingerprint(cert *x509.Certificate, algo crypto.Hash) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Hash.Writer is specified to be never returning an error.
// https://golang.org/pkg/hash/#Hash
