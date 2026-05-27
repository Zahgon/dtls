// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package recordlayer

import (
	"github.com/pion/dtls/v3/pkg/protocol"
)

// DTLS fixed size record layer header when Connection IDs are not in-use.

// ---------------------------------
// | Type   |   Version   |  Epoch |
// ---------------------------------
// | Epoch  |    Sequence Number   |
// ---------------------------------
// |   Sequence Number   |  Length |
// ---------------------------------
// | Length |      Fragment...     |
// ---------------------------------

// fixedHeaderLenIdx is the index at which the record layer content length is
// specified in a fixed length header (i.e. one that does not include a
// Connection ID).
const fixedHeaderLenIdx = 11

// RecordLayer which handles all data transport.
// The record layer is assumed to sit directly on top of some
// reliable transport such as TCP. The record layer can carry four types of content:
//
// 1. Handshake messages—used for algorithm negotiation and key establishment.
// 2. ChangeCipherSpec messages—really part of the handshake but technically a separate kind of message.
// 3. Alert messages—used to signal that errors have occurred
// 4. Application layer data
//
// The DTLS record layer is extremely similar to that of TLS 1.1.  The
// only change is the inclusion of an explicit sequence number in the
// record.  This sequence number allows the recipient to correctly
// verify the TLS MAC.
//
// https://tools.ietf.org/html/rfc4347#section-4.1
type RecordLayer struct {
	Header  Header
	Content protocol.Content
}

// Marshal encodes the RecordLayer to binary.
func (r *RecordLayer) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115

// Unmarshal populates the RecordLayer from binary.
func (r *RecordLayer) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnpackDatagram extracts all RecordLayer messages from a single datagram.
// Note that as with TLS, multiple handshake messages may be placed in
// the same DTLS record, provided that there is room and that they are
// part of the same flight.  Thus, there are two acceptable ways to pack
// two DTLS messages into the same datagram: in the same record or in
// separate records.
// https://tools.ietf.org/html/rfc6347#section-4.2.3
func UnpackDatagram(buf []byte) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ContentAwareUnpackDatagram is the same as UnpackDatagram but considers the
// presence of a connection identifier if the record is of content type
// tls12_cid.
func ContentAwareUnpackDatagram(buf []byte, cidLength int) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
