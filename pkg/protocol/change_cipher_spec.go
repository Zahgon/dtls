// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

// ChangeCipherSpec protocol exists to signal transitions in
// ciphering strategies.  The protocol consists of a single message,
// which is encrypted and compressed under the current (not the pending)
// connection state.  The message consists of a single byte of value 1.
// https://tools.ietf.org/html/rfc5246#section-7.1
type ChangeCipherSpec struct{}

// ContentType returns the ContentType of this content.
func (c ChangeCipherSpec) ContentType() ContentType {
	_ = "STUB: not implemented"
	return *new(ContentType)
}

// Marshal encodes the ChangeCipherSpec to binary.
func (c *ChangeCipherSpec) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unmarshal populates the ChangeCipherSpec from binary.
		nil
}

func (c *ChangeCipherSpec) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
