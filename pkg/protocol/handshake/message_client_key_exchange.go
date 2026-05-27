// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/internal/ciphersuite/types"
)

// MessageClientKeyExchange is a DTLS Handshake Message
// With this message, the premaster secret is set, either by direct
// transmission of the RSA-encrypted secret or by the transmission of
// Diffie-Hellman parameters that will allow each side to agree upon
// the same premaster secret.
//
// https://tools.ietf.org/html/rfc5246#section-7.4.7
type MessageClientKeyExchange struct {
	IdentityHint []byte
	PublicKey    []byte

	// for unmarshaling
	KeyExchangeAlgorithm types.KeyExchangeAlgorithm
}

// Type returns the Handshake Type.
func (m MessageClientKeyExchange) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Marshal encodes the Handshake.
func (m *MessageClientKeyExchange) Marshal() (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115

//nolint:gosec // G115: public key length is validated to be <= 255 above.

// Unmarshal populates the message from encoded data.
func (m *MessageClientKeyExchange) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
