// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

const (
	renegotiationInfoHeaderSize = 5
)

// RenegotiationInfo allows a Client/Server to
// communicate their renegotation support
//
// https://tools.ietf.org/html/rfc5746
type RenegotiationInfo struct {
	RenegotiatedConnection uint8
}

// TypeValue returns the extension TypeValue.
func (r RenegotiationInfo) TypeValue() TypeValue { _ = "STUB: not implemented"; return *new(TypeValue) }

// Marshal encodes the extension.
func (r *RenegotiationInfo) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// length

// Unmarshal populates the extension from encoded data.
func (r *RenegotiationInfo) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
