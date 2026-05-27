// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package recordlayer

// UnifiedHeader implements the DTLS 1.3 Unified Header.
// See RFC 9147 section 4. The DTLS Record Layer
//
// https://datatracker.ietf.org/doc/html/rfc9147#name-the-dtls-record-layer
//
//	 0 1 2 3 4 5 6 7
//	+-+-+-+-+-+-+-+-+
//	|0|0|1|C|S|L|E E|
//	+-+-+-+-+-+-+-+-+
//	| Connection ID |   Legend:
//	| (if any,      |
//	/  length as    /   C   - Connection ID (CID) present
//	|  negotiated)  |   S   - Sequence number length
//	+-+-+-+-+-+-+-+-+   L   - Length present
//	|  8 or 16 bit  |   E   - Epoch
//	|Sequence Number|
//	+-+-+-+-+-+-+-+-+
//	| 16 bit Length |
//	| (if present)  |
//	+-+-+-+-+-+-+-+-+
type UnifiedHeader struct {
	ConnectionID   []byte // size of array should be expected CID length
	SequenceNumber uint16
	SeqBit         bool
	Length         uint16
	LengthBit      bool
	EpochLow       uint8
}

const (
	UnifiedHeaderFixedBits = 0b00100000
	UnifiedHeaderCIDBit    = 0b00010000
	UnifiedHeaderSeqBit    = 0b00001000
	UnifiedHeaderLengthBit = 0b00000100
	TwoLowBitsMask         = 0b11
)

// Marshal encodes a DTLS 1.3 Unified Header to binary.
func (u *UnifiedHeader) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec

// Unmarshal populates a DTLS 1.3 Unified Header from binary.
func (u *UnifiedHeader) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (u *UnifiedHeader) Size() int { _ = "STUB: not implemented"; return 0 }
