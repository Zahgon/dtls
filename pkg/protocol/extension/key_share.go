// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"golang.org/x/crypto/cryptobyte"
)

type KeyShareEntry struct {
	Group       elliptic.Curve
	KeyExchange []byte
}

// KeyShare represents the "key_share" extension. Only one of the fields can be used at a time.
// See RFC 8446 section 4.2.8.
type KeyShare struct {
	ClientShares  []KeyShareEntry // ClientHello
	ServerShare   *KeyShareEntry  // ServerHello
	SelectedGroup *elliptic.Curve // HelloRetryRequest
}

func (k KeyShare) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the extension.
	new(TypeValue)
}

func (k *KeyShare) Marshal() ([]byte, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil
}

// vector MAY be empty

// there must be exactly one context.

// KeyShareHelloRetryRequest { NamedGroup selected_group; }

// KeyShareServerHello { KeyShareEntry server_share; }

// KeyShareClientHello { KeyShareEntry client_shares<0..2^16-1>; }

// Unmarshal decodes the extension.
func (k *KeyShare) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// ClientHello: client_shares is a uint16-length-prefixed vector.
//nolint:nestif

// consume vector (2 bytes length + vecLen)

// HelloRetryRequest: exactly 2 bytes = selected_group

// ServerHello: exactly one KeyShareEntry and no trailing bytes

func addKeyShareEntry(b *cryptobyte.Builder, e KeyShareEntry) { _ = "STUB: not implemented"; return }

// hasTooManyContexts is used in Marshal(). It returns whether the KeyShare struct has more than exactly one context.
func hasTooManyContexts(a bool, b bool, c bool) bool { _ = "STUB: not implemented"; return false }
