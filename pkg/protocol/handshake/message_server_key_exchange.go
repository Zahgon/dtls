// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/internal/ciphersuite/types"
	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/crypto/hash"
	"github.com/pion/dtls/v3/pkg/crypto/signature"
)

// MessageServerKeyExchange supports ECDH and PSK.
type MessageServerKeyExchange struct {
	IdentityHint []byte

	EllipticCurveType  elliptic.CurveType
	NamedCurve         elliptic.Curve
	PublicKey          []byte
	HashAlgorithm      hash.Algorithm
	SignatureAlgorithm signature.Algorithm
	Signature          []byte

	// for unmarshaling
	KeyExchangeAlgorithm types.KeyExchangeAlgorithm
}

// Type returns the Handshake Type.
func (m MessageServerKeyExchange) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Marshal encodes the Handshake.
func (m *MessageServerKeyExchange) Marshal() ([]byte, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil
}

//nolint:gosec //G115

//nolint:gosec // G115, no risk of overflow, the biggest supported curve is 97 bytes.

//nolint:gosec // G115

// Unmarshal populates the message from encoded data.
func (m *MessageServerKeyExchange) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Anon connection doesn't contains hashAlgorithm, signatureAlgorithm, signature
