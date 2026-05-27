// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// CertificateAuthorities implements the certificate_authorities extension in DTLS 1.3.
//
// See RFC 8446 section 4.2.4. Certificate Authorities.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.4
type CertificateAuthorities struct {
	Authorities [][]byte
}

// TypeValue returns the extension TypeValue.
func (c CertificateAuthorities) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *new(TypeValue)
}

// Marshal encodes the extension.
func (c *CertificateAuthorities) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal populates the extension from encoded data.
func (c *CertificateAuthorities) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
