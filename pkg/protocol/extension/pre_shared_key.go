// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// PreSharedKey represents the "pre_shared_key" extension for DTLS 1.3.
// This extension is used in both ClientHello and ServerHello messages,
// but only the relevant fields should be populated for each context.
// See RFC 8446 section 4.2.11.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.11
type PreSharedKey struct {
	// ClientHello only - offered PSK identities
	Identities []PskIdentity
	// ClientHello only - binder values associated with a PSK identity
	Binders []PskBinderEntry
	// ServerHello only - index of selected identity
	SelectedIdentity uint16
}

// PskIdentity represents the PSK identitiy in the "pre_shared_key" extension
// for DTLS 1.3.
type PskIdentity struct {
	Identity            []byte
	ObfuscatedTicketAge uint32
}

// PskBinderEntry represents the binder related to a PSK identity in the
// "pre_shared_key" extension for DTLS 1.3.
type PskBinderEntry []byte

const minPSKBinderSize = 32

// TypeValue returns the extension TypeValue.
func (p PreSharedKey) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the extension.
	new(TypeValue)
}

func (p *PreSharedKey) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ServerHello

// ClientHello

// Unmarshal populates the extension from encoded data.
func (p *PreSharedKey) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// ServerHello

// ClientHello

// Ensure there is one binder value per identity in list
