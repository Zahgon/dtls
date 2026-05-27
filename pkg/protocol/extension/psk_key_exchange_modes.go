// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// PskKeyExchangeModes implements the PskKeyExchangeModes extension in DTLS 1.3.
// See RFC 8446 section 4.2.9. Pre-Shared Key Exchange Modes.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.9
type PskKeyExchangeModes struct {
	KeModes []PskKeyExchangeMode
}

type PskKeyExchangeMode uint8

// TypeValue constants.
const (
	PskKe    PskKeyExchangeMode = 0
	PskDheKe PskKeyExchangeMode = 1
)

// TypeValue returns the extension TypeValue.
func (p PskKeyExchangeModes) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *new(TypeValue)
}

// Marshal encodes the extension.
func (p *PskKeyExchangeModes) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the extension from encoded data.
func (p *PskKeyExchangeModes) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
