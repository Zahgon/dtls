// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension // nolint:dupl

const (
	useExtendedMasterSecretHeaderSize = 4
)

// UseExtendedMasterSecret defines a TLS extension that contextually binds the
// master secret to a log of the full handshake that computes it, thus
// preventing MITM attacks.
type UseExtendedMasterSecret struct {
	Supported bool
}

// TypeValue returns the extension TypeValue.
func (u UseExtendedMasterSecret) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *new(TypeValue)
}

// Marshal encodes the extension.
func (u *UseExtendedMasterSecret) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// length

// Unmarshal populates the extension from encoded data.
func (u *UseExtendedMasterSecret) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
