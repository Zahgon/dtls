// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

// MessageServerHelloDone is final non-encrypted message from server
// this communicates server has sent all its handshake messages and next
// should be MessageFinished.
type MessageServerHelloDone struct{}

// Type returns the Handshake Type.
func (m MessageServerHelloDone) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Marshal encodes the Handshake.
func (m *MessageServerHelloDone) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unmarshal populates the message from encoded data.
		nil
}

func (m *MessageServerHelloDone) Unmarshal([]byte) error { _ = "STUB: not implemented"; return nil }
