// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

// MessageFinished is a DTLS Handshake Message
// this message is the first one protected with the just
// negotiated algorithms, keys, and secrets.  Recipients of Finished
// messages MUST verify that the contents are correct.
//
// https://tools.ietf.org/html/rfc5246#section-7.4.9
type MessageFinished struct {
	VerifyData []byte
}

// Type returns the Handshake Type.
func (m MessageFinished) Type() Type {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the Handshake.
	new(Type)
}

func (m *MessageFinished) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the message from encoded data.
func (m *MessageFinished) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
