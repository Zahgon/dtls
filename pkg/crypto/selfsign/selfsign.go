// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package selfsign is a test helper that generates self signed certificate.
package selfsign

import (
	"crypto"
	"crypto/tls"
	"errors"
)

var errInvalidPrivateKey = errors.New("selfsign: invalid private key type")

// GenerateSelfSigned creates a self-signed certificate.
func GenerateSelfSigned() (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// GenerateSelfSignedWithDNS creates a self-signed certificate.
func GenerateSelfSignedWithDNS(cn string, sans ...string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// SelfSign creates a self-signed certificate from a elliptic curve key.
func SelfSign(key crypto.PrivateKey) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// WithDNS creates a self-signed certificate from a elliptic curve key.
func WithDNS(key crypto.PrivateKey, cn string, sans ...string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// Max random value, a 130-bits integer, i.e 2^130 - 1

/* #nosec */

/* #nosec */
