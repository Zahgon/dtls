// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

const serverNameTypeDNSHostName = 0

// ServerName allows the client to inform the server the specific
// name it wishes to contact. Useful if multiple DNS names resolve
// to one IP
//
// https://tools.ietf.org/html/rfc6066#section-3
type ServerName struct {
	ServerName string
}

// TypeValue returns the extension TypeValue.
func (s ServerName) TypeValue() TypeValue { _ = "STUB: not implemented"; return *new(TypeValue) }

// Marshal encodes the extension.
func (s *ServerName) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the extension from encoded data.
func (s *ServerName) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Multiple names of the same name_type are prohibited.

// An SNI value may not include a trailing dot.
