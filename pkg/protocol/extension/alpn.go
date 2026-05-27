// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// ALPN is a TLS extension for application-layer protocol negotiation within
// the TLS handshake.
//
// https://tools.ietf.org/html/rfc7301
type ALPN struct {
	ProtocolNameList []string
}

// TypeValue returns the extension TypeValue.
func (a ALPN) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the extension.
	new(TypeValue)
}

func (a *ALPN) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Satisfy range scope lint

// Unmarshal populates the extension from encoded data.
func (a *ALPN) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// ALPNProtocolSelection negotiates a shared protocol according to #3.2 of rfc7301.
func ALPNProtocolSelection(supportedProtocols, peerSupportedProtocols []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
