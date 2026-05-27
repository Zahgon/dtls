// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

const (
	useSRTPHeaderSize = 6
	maxUint16         = (1 << 16) - 1
)

// UseSRTP allows a Client/Server to negotiate what SRTPProtectionProfiles
// they both support
//
// https://tools.ietf.org/html/rfc8422
type UseSRTP struct {
	ProtectionProfiles  []SRTPProtectionProfile
	MasterKeyIdentifier []byte
}

// TypeValue returns the extension TypeValue.
func (u UseSRTP) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the extension.
	new(TypeValue)
}

func (u *UseSRTP) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115

//nolint:gosec // G115

//nolint:gosec // G115: MKI length is validated to be <= 255 above.

// Unmarshal populates the extension from encoded data.
func (u *UseSRTP) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
