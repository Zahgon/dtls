// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package recordlayer

import (
	"github.com/pion/dtls/v3/pkg/protocol"
)

// InnerPlaintext implements DTLSInnerPlaintext
//
// https://datatracker.ietf.org/doc/html/rfc9146#name-record-layer-extensions
type InnerPlaintext struct {
	Content  []byte
	RealType protocol.ContentType
	Zeros    uint
}

// Marshal encodes a DTLS InnerPlaintext to binary.
func (p *InnerPlaintext) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates a DTLS InnerPlaintext from binary.
func (p *InnerPlaintext) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	// Process in reverse
	return nil
}

//nolint:gosec // G115
