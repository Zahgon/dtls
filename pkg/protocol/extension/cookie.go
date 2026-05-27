// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

const maxCookieSize = 0xffff - 2

// CookieExt implements the cookie extension in DTLS 1.3.
// See RFC 8446 section 4.2.2. Cookie.
type CookieExt struct {
	Cookie []byte
}

// TypeValue returns the extension TypeValue.
func (c CookieExt) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *

	// Marshal encodes the extension.
	new(TypeValue)
}

func (c *CookieExt) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the extension from encoded data.
func (c *CookieExt) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
