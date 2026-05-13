// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package recordlayer

import (
	"encoding/binary"

	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/alert"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
)

const (
	maxDTLSPlaintextRecordLen  = 1 << 14
	maxDTLSCiphertextRecordLen = maxDTLSPlaintextRecordLen + 256
)

// HeaderLike is implemented by DTLS record header encodings.
type HeaderLike interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	Size() int
}

// RecordLayer13 is implemented by DTLS 1.3 plaintext and ciphertext records.
type RecordLayer13 interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	RecordHeader() HeaderLike
}

// PlaintextRecord13 implements DTLSPlaintext for epoch 0 records.
type PlaintextRecord13 struct {
	Header  Header
	Content protocol.Content
}

// Marshal encodes a DTLS 1.3 DTLSPlaintext record.
func (r *PlaintextRecord13) Marshal() ([]byte, error) {
	if r.Header.Epoch != 0 {
		return nil, errInvalidEpoch
	}
	if r.Header.Version == (protocol.Version{}) {
		r.Header.Version = protocol.Version1_2
	}
	if !r.Header.Version.Equal(protocol.Version1_2) {
		return nil, errUnsupportedProtocolVersion
	}

	contentRaw, err := r.Content.Marshal()
	if err != nil {
		return nil, err
	}
	if len(contentRaw) > maxDTLSPlaintextRecordLen {
		return nil, ErrInvalidPacketLength
	}

	r.Header.ContentLen = uint16(len(contentRaw)) //nolint:gosec // G115: checked above.
	r.Header.ContentType = r.Content.ContentType()

	headerRaw, err := r.Header.Marshal()
	if err != nil {
		return nil, err
	}

	return append(headerRaw, contentRaw...), nil
}

// Unmarshal populates a DTLS 1.3 DTLSPlaintext record from binary.
func (r *PlaintextRecord13) Unmarshal(data []byte) error {
	if err := r.Header.Unmarshal(data); err != nil {
		return err
	}
	if r.Header.Epoch != 0 {
		return errInvalidEpoch
	}
	if !r.Header.Version.Equal(protocol.Version1_2) {
		return errUnsupportedProtocolVersion
	}
	if r.Header.ContentLen > maxDTLSPlaintextRecordLen {
		return ErrInvalidPacketLength
	}

	switch r.Header.ContentType {
	case protocol.ContentTypeChangeCipherSpec:
		r.Content = &protocol.ChangeCipherSpec{}
	case protocol.ContentTypeAlert:
		r.Content = &alert.Alert{}
	case protocol.ContentTypeHandshake:
		r.Content = &handshake.Handshake{}
	case protocol.ContentTypeApplicationData:
		r.Content = &protocol.ApplicationData{}
	default:
		return errInvalidContentType
	}

	contentStart := r.Header.Size()
	contentEnd := contentStart + int(r.Header.ContentLen)
	if len(data) != contentEnd {
		return ErrInvalidPacketLength
	}

	return r.Content.Unmarshal(data[contentStart:contentEnd])
}

// RecordHeader returns the record header.
func (r *PlaintextRecord13) RecordHeader() HeaderLike {
	return &r.Header
}

// CiphertextRecord13 implements DTLSCiphertext for protected records.
type CiphertextRecord13 struct {
	Header          UnifiedHeader
	EncryptedRecord []byte
}

// Marshal encodes a DTLS 1.3 DTLSCiphertext record.
func (r *CiphertextRecord13) Marshal() ([]byte, error) {
	if len(r.EncryptedRecord) > maxDTLSCiphertextRecordLen {
		return nil, ErrInvalidPacketLength
	}
	r.Header.SeqBit = true
	r.Header.Length = uint16(len(r.EncryptedRecord)) //nolint:gosec // G115: checked above.
	r.Header.LengthBit = true

	headerRaw, err := r.Header.Marshal()
	if err != nil {
		return nil, err
	}

	out := make([]byte, 0, len(headerRaw)+len(r.EncryptedRecord))
	out = append(out, headerRaw...)
	out = append(out, r.EncryptedRecord...)

	return out, nil
}

// Unmarshal populates a DTLS 1.3 DTLSCiphertext record from binary.
func (r *CiphertextRecord13) Unmarshal(data []byte) error {
	if err := r.Header.Unmarshal(data); err != nil {
		return err
	}

	headerSize := unifiedHeaderWireSize(data[0], len(r.Header.ConnectionID))
	if len(data) < headerSize {
		return errBufferTooSmall
	}

	recordLen := len(data) - headerSize
	if r.Header.LengthBit {
		recordLen = int(r.Header.Length)
		if len(data)-headerSize != recordLen {
			return ErrInvalidPacketLength
		}
	}
	if recordLen > maxDTLSCiphertextRecordLen {
		return ErrInvalidPacketLength
	}

	r.EncryptedRecord = append([]byte{}, data[headerSize:headerSize+recordLen]...)

	return nil
}

// RecordHeader returns the record header.
func (r *CiphertextRecord13) RecordHeader() HeaderLike {
	return &r.Header
}

// UnpackDatagram13 extracts DTLS 1.3 records from a single datagram.
func UnpackDatagram13(buf []byte, cidLength int, ciphertextHeadersEnabled bool) ([][]byte, error) {
	out := [][]byte{}

	for offset := 0; len(buf) != offset; {
		if ciphertextHeadersEnabled {
			if !protocol.IsDTLS13Ciphertext(protocol.ContentType(buf[offset])) {
				return nil, errInvalidContentType
			}

			header := UnifiedHeader{}
			if buf[offset]&UnifiedHeaderCIDBit != 0 {
				header.ConnectionID = make([]byte, cidLength)
			}
			if err := header.Unmarshal(buf[offset:]); err != nil {
				return nil, err
			}

			headerSize := unifiedHeaderWireSize(buf[offset], len(header.ConnectionID))
			if !header.LengthBit {
				out = append(out, buf[offset:])

				return out, nil
			}

			pktLen := headerSize + int(header.Length)
			if offset+pktLen > len(buf) {
				return nil, ErrInvalidPacketLength
			}

			out = append(out, buf[offset:offset+pktLen])
			offset += pktLen

			continue
		}

		if len(buf)-offset <= FixedHeaderSize {
			return nil, ErrInvalidPacketLength
		}

		pktLen := FixedHeaderSize + int(binary.BigEndian.Uint16(buf[offset+fixedHeaderLenIdx:]))
		if offset+pktLen > len(buf) {
			return nil, ErrInvalidPacketLength
		}

		out = append(out, buf[offset:offset+pktLen])
		offset += pktLen
	}

	return out, nil
}

func unifiedHeaderWireSize(firstByte byte, cidLength int) int {
	size := 1 + cidLength
	if firstByte&UnifiedHeaderSeqBit != 0 {
		size += 2
	} else {
		size++
	}
	if firstByte&UnifiedHeaderLengthBit != 0 {
		size += 2
	}

	return size
}
