// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

// HeaderLength msg_len for Handshake messages assumes an extra
// 12 bytes for sequence, fragment and version information vs TLS.
const HeaderLength = 12

// Header is the static first 12 bytes of each RecordLayer
// of type Handshake. These fields allow us to support message loss, reordering, and
// message fragmentation,
//
// https://tools.ietf.org/html/rfc6347#section-4.2.2
type Header struct {
	Type            Type
	Length          uint32 // uint24 in spec
	MessageSequence uint16
	FragmentOffset  uint32 // uint24 in spec
	FragmentLength  uint32 // uint24 in spec
}

// Marshal encodes the Header.
func (h *Header) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the header from encoded data.
func (h *Header) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }
