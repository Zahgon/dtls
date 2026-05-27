// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"github.com/pion/dtls/v3/pkg/protocol"
)

// SupportedVersions is a TLS extension used by the client to indicate
// which versions of TLS it supports and by the server to indicate which
// version it is using.
//
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.1
type SupportedVersions struct {
	// ClientHello's preference-ordered list.
	Versions []protocol.Version
}

func (s SupportedVersions) TypeValue() TypeValue { _ = "STUB: not implemented"; return *new(TypeValue) }

// Marshal encodes the extension without carrying negotiation state.
func (s *SupportedVersions) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// The 2..254 bound is defined in the following:
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.1

// We're only checking for *valid* versions, not to be confused with supported versions.
// Error on invalid versions to protect against malformed messages/DOS attacks.

// in the case that there's only one version, the do not add the length (uint8).

// Unmarshal parses either the ClientHello list or the ServerHello/HelloRetryRequest single value.
// Any version not recognized is discarded.
func (s *SupportedVersions) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

// Try ClientHello list: versions<2..254> (1-byte length, then pairs)

// We're only checking for *valid* versions, not to be confused with supported versions.

// Otherwise, expect ServerHello/HelloRetryRequest selected_version, which should be exactly 2 bytes.

// We're only checking for *valid* versions, not to be confused with supported versions.
