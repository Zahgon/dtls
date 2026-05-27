// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package elliptic provides elliptic curve cryptography for DTLS
package elliptic

import (
	"crypto/ecdh"
	"errors"
)

var errInvalidNamedCurve = errors.New("invalid named curve")

// CurvePointFormat is used to represent the IANA registered curve points
//
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-9
type CurvePointFormat byte

// CurvePointFormat enums.
const (
	CurvePointFormatUncompressed CurvePointFormat = 0
)

// Keypair is a Curve with a Private/Public Keypair.
type Keypair struct {
	Curve      Curve
	PublicKey  []byte
	PrivateKey []byte //nolint:gosec // no real risk of exporting the private key.
}

// CurveType is used to represent the IANA registered curve types for TLS
//
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-10
type CurveType byte

// CurveType enums.
const (
	CurveTypeNamedCurve CurveType = 0x03
)

// CurveTypes returns all known curves.
func CurveTypes() map[CurveType]struct{} { _ = "STUB: not implemented"; return nil }

// Curve is used to represent the IANA registered curves for TLS
//
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8
type Curve uint16

// Curve enums.
const (
	P256   Curve = 0x0017
	P384   Curve = 0x0018
	X25519 Curve = 0x001d
	// X25519MLKEM768
	// https://pkg.go.dev/crypto/internal/fips140/mlkem
	// https://datatracker.ietf.org/doc/draft-ietf-tls-hybrid-design/
	// https://datatracker.ietf.org/doc/draft-ietf-tls-ecdhe-mlkem/
)

func (c Curve) String() string { _ = "STUB: not implemented"; return "" }

// Curves returns all curves we implement.
func Curves() map[Curve]bool { _ = "STUB: not implemented"; return nil }

// GenerateKeypair generates a keypair for the given Curve.
func GenerateKeypair(curve Curve) (*Keypair, error) { _ = "STUB: not implemented"; return nil, nil }

// NIST: SEC1 uncompressed (04||X||Y); X25519: 32 bytes
// Scalar suitable for ecdh.NewPrivateKey

// toECDH returns the crypto/ecdh curve for our enum.
func (c Curve) toECDH() (ecdh.Curve, error) {
	_ = "STUB: not implemented"
	return *new(ecdh.Curve), nil
}
