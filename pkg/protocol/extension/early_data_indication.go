// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// EarlyDataIndication implements the early data indication extension in DTLS 1.3.
// See RFC 8446 section 4.2.10. Early Data Indication.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.10
type EarlyDataIndication struct {
	MaxEarlyData *uint32 // nil indicates CH or EE
}

// TypeValue returns the extension TypeValue.
func (e EarlyDataIndication) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *new(TypeValue)
}

// Marshal encodes the extension.
func (e *EarlyDataIndication) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// zero length

// new_session_ticket

// Unmarshal populates the extension from encoded data.
func (e *EarlyDataIndication) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// new_session_ticket
