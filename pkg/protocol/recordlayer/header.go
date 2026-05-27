// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package recordlayer

import (
	"github.com/pion/dtls/v3/pkg/protocol"
)

// Header implements a TLS RecordLayer header.
type Header struct {
	ContentType    protocol.ContentType
	ContentLen     uint16
	Version        protocol.Version
	Epoch          uint16
	SequenceNumber uint64 // uint48 in spec

	// Optional Fields
	ConnectionID []byte
}

// RecordLayer enums.
const (
	// FixedHeaderSize is the size of a DTLS record header when connection IDs
	// are not in use.
	FixedHeaderSize   = 13
	MaxSequenceNumber = 0x0000FFFFFFFFFFFF
)

// Marshal encodes a TLS RecordLayer Header to binary.
func (h *Header) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates a TLS RecordLayer Header from binary.
func (h *Header) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// If a CID was expected the ConnectionID should have been initialized.

// SequenceNumber is stored as uint48, make into uint64

// Size returns the total size of the header.
func (h *Header) Size() int { _ = "STUB: not implemented"; return 0 }
