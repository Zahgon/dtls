// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
)

const (
	supportedGroupsHeaderSize = 6
)

// SupportedEllipticCurves allows a Client/Server to communicate
// what curves they both support
//
// https://tools.ietf.org/html/rfc8422#section-5.1.1
//
// In DTLS 1.3, this extension in renamed to "supported_groups".
// https://datatracker.ietf.org/doc/html/rfc8446#section-4.2.7
type SupportedEllipticCurves struct {
	EllipticCurves []elliptic.Curve
}

// TypeValue returns the extension TypeValue.
func (s SupportedEllipticCurves) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *new(TypeValue)
}

// Marshal encodes the extension.
func (s *SupportedEllipticCurves) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115
//nolint:gosec // G115

//nolint:makezero // todo: fix

// Unmarshal populates the extension from encoded data.
func (s *SupportedEllipticCurves) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// type + declared length = 4
