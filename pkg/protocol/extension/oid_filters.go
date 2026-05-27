// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

// OIDFilters defines a DTLS 1.3 extension that is used to allow server to
// provide a set of OID/value pairs which it would like the client's
// certificate to match.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.5
type OIDFilters struct {
	Filters []OIDFilter
}

type OIDFilter struct {
	OID    []byte
	Values []byte
}

// TypeValue returns the extension TypeValue.
func (o OIDFilters) TypeValue() TypeValue { _ = "STUB: not implemented"; return *new(TypeValue) }

// Marshal encodes the extension.
func (o *OIDFilters) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the extension from encoded data.
func (o *OIDFilters) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
