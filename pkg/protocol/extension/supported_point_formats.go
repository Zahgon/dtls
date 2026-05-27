// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
)

const (
	supportedPointFormatsSize = 5
)

// SupportedPointFormats allows a Client/Server to negotiate
// the EllipticCurvePointFormats
//
// https://tools.ietf.org/html/rfc4492#section-5.1.2
type SupportedPointFormats struct {
	PointFormats []elliptic.CurvePointFormat
}

// TypeValue returns the extension TypeValue.
func (s SupportedPointFormats) TypeValue() TypeValue {
	_ = "STUB: not implemented"
	return *new(TypeValue)
}

// Marshal encodes the extension.
func (s *SupportedPointFormats) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115
//nolint:gosec // G115: point format count is validated to be <= 255 above.

//nolint:makezero // todo: fix

// Unmarshal populates the extension from encoded data.
func (s *SupportedPointFormats) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// type + declared length = 4
