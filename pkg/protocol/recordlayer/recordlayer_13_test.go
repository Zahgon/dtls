// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package recordlayer

import (
	"testing"

	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/alert"
	"github.com/stretchr/testify/require"
)

func TestPlaintextRecord13RoundTrip(t *testing.T) {
	record := &PlaintextRecord13{
		Header: Header{
			Version: protocol.Version1_2,
		},
		Content: &alert.Alert{Level: alert.Warning, Description: alert.CloseNotify},
	}

	raw, err := record.Marshal()
	require.NoError(t, err)
	require.Equal(t, []byte{
		0x15, 0xfe, 0xfd,
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x02,
		0x01, 0x00,
	}, raw)

	var roundTrip PlaintextRecord13
	require.NoError(t, roundTrip.Unmarshal(raw))
	require.Equal(t, protocol.ContentTypeAlert, roundTrip.Header.ContentType)
	require.Equal(t, protocol.Version1_2, roundTrip.Header.Version)
	require.Equal(t, uint16(0), roundTrip.Header.Epoch)
	require.Equal(t, uint16(2), roundTrip.Header.ContentLen)

	got, ok := roundTrip.Content.(*alert.Alert)
	require.True(t, ok)
	require.Equal(t, alert.Warning, got.Level)
	require.Equal(t, alert.CloseNotify, got.Description)
}

func TestPlaintextRecord13RejectsProtectedEpoch(t *testing.T) {
	record := &PlaintextRecord13{
		Header: Header{
			Version: protocol.Version1_2,
			Epoch:   1,
		},
		Content: &alert.Alert{Level: alert.Warning, Description: alert.CloseNotify},
	}

	_, err := record.Marshal()
	require.ErrorIs(t, err, errInvalidEpoch)
}

func TestPlaintextRecord13RejectsDTLS10Version(t *testing.T) {
	record := &PlaintextRecord13{
		Header:  Header{Version: protocol.Version1_0},
		Content: &alert.Alert{Level: alert.Warning, Description: alert.CloseNotify},
	}

	_, err := record.Marshal()
	require.ErrorIs(t, err, errUnsupportedProtocolVersion)

	var roundTrip PlaintextRecord13
	err = roundTrip.Unmarshal([]byte{
		0x15, 0xfe, 0xff,
		0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x02,
		0x01, 0x00,
	})
	require.ErrorIs(t, err, errUnsupportedProtocolVersion)
}

func TestPlaintextRecord13RejectsOversizedContent(t *testing.T) {
	record := &PlaintextRecord13{
		Header:  Header{Version: protocol.Version1_2},
		Content: &protocol.ApplicationData{Data: make([]byte, maxDTLSPlaintextRecordLen+1)},
	}

	_, err := record.Marshal()
	require.ErrorIs(t, err, ErrInvalidPacketLength)
}

func TestPlaintextRecord13RejectsOversizedUnmarshal(t *testing.T) {
	header := Header{
		ContentType: protocol.ContentTypeApplicationData,
		Version:     protocol.Version1_2,
		ContentLen:  maxDTLSPlaintextRecordLen + 1,
	}
	raw, err := header.Marshal()
	require.NoError(t, err)
	raw = append(raw, make([]byte, maxDTLSPlaintextRecordLen+1)...)

	var record PlaintextRecord13
	err = record.Unmarshal(raw)
	require.ErrorIs(t, err, ErrInvalidPacketLength)
}

func TestCiphertextRecord13RoundTrip(t *testing.T) {
	record := &CiphertextRecord13{
		Header: UnifiedHeader{
			EpochLow:       3,
			SequenceNumber: 0xaabb,
		},
		EncryptedRecord: []byte{0xde, 0xad, 0xbe, 0xef},
	}

	raw, err := record.Marshal()
	require.NoError(t, err)
	require.Equal(t, []byte{
		0x2f,
		0xaa, 0xbb,
		0x00, 0x04,
		0xde, 0xad, 0xbe, 0xef,
	}, raw)

	var roundTrip CiphertextRecord13
	require.NoError(t, roundTrip.Unmarshal(raw))
	require.Equal(t, uint8(3), roundTrip.Header.EpochLow)
	require.Equal(t, uint16(0xaabb), roundTrip.Header.SequenceNumber)
	require.True(t, roundTrip.Header.SeqBit)
	require.Equal(t, uint16(4), roundTrip.Header.Length)
	require.True(t, roundTrip.Header.LengthBit)
	require.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, roundTrip.EncryptedRecord)
}

func TestCiphertextRecord13MarshalRefreshesEmptyLength(t *testing.T) {
	record := &CiphertextRecord13{
		Header: UnifiedHeader{
			SequenceNumber: 0x01,
			Length:         4,
		},
		EncryptedRecord: []byte{0xaa},
	}

	raw, err := record.Marshal()
	require.NoError(t, err)
	require.Equal(t, []byte{0x2c, 0x00, 0x01, 0x00, 0x01, 0xaa}, raw)
	require.Equal(t, uint16(1), record.Header.Length)
	require.True(t, record.Header.SeqBit)
	require.True(t, record.Header.LengthBit)

	record.EncryptedRecord = nil
	raw, err = record.Marshal()
	require.NoError(t, err)
	require.Equal(t, []byte{0x2c, 0x00, 0x01, 0x00, 0x00}, raw)
	require.Equal(t, uint16(0), record.Header.Length)
	require.True(t, record.Header.SeqBit)
	require.True(t, record.Header.LengthBit)
}

func TestCiphertextRecord13RejectsOversizedEncryptedRecord(t *testing.T) {
	record := &CiphertextRecord13{
		EncryptedRecord: make([]byte, maxDTLSCiphertextRecordLen+1),
	}

	_, err := record.Marshal()
	require.ErrorIs(t, err, ErrInvalidPacketLength)
}

func TestCiphertextRecord13WithoutLengthUsesRemainder(t *testing.T) {
	raw := []byte{0x21, 0x12, 0xaa, 0xbb, 0xcc}
	require.Equal(t, []byte{0x21, 0x12, 0xaa, 0xbb, 0xcc}, raw)

	var roundTrip CiphertextRecord13
	require.NoError(t, roundTrip.Unmarshal(raw))
	require.Equal(t, uint8(1), roundTrip.Header.EpochLow)
	require.Equal(t, uint16(0x12), roundTrip.Header.SequenceNumber)
	require.False(t, roundTrip.Header.SeqBit)
	require.Equal(t, uint16(0), roundTrip.Header.Length)
	require.False(t, roundTrip.Header.LengthBit)
	require.Equal(t, []byte{0xaa, 0xbb, 0xcc}, roundTrip.EncryptedRecord)
}

func TestCiphertextRecord13RejectsLengthMismatch(t *testing.T) {
	var record CiphertextRecord13
	err := record.Unmarshal([]byte{0x2c, 0x00, 0x01, 0x00, 0x04, 0xaa, 0xbb})
	require.ErrorIs(t, err, ErrInvalidPacketLength)
}

func TestUnpackDatagram13Plaintext(t *testing.T) {
	plaintext := &PlaintextRecord13{
		Header:  Header{Version: protocol.Version1_2},
		Content: &alert.Alert{Level: alert.Warning, Description: alert.CloseNotify},
	}
	plaintextRaw, err := plaintext.Marshal()
	require.NoError(t, err)

	records, err := UnpackDatagram13(plaintextRaw, 0, false)
	require.NoError(t, err)
	require.Equal(t, [][]byte{plaintextRaw}, records)
}

func TestUnpackDatagram13Ciphertext(t *testing.T) {
	ciphertextWithLength := &CiphertextRecord13{
		Header: UnifiedHeader{
			SequenceNumber: 0x01,
		},
		EncryptedRecord: []byte{0xaa, 0xbb},
	}
	ciphertextWithLengthRaw, err := ciphertextWithLength.Marshal()
	require.NoError(t, err)

	ciphertextWithoutLengthRaw := []byte{0x20, 0x02, 0xcc, 0xdd}

	datagram := append(append([]byte{}, ciphertextWithLengthRaw...), ciphertextWithoutLengthRaw...)
	records, err := UnpackDatagram13(datagram, 0, true)
	require.NoError(t, err)
	require.Equal(t, [][]byte{ciphertextWithLengthRaw, ciphertextWithoutLengthRaw}, records)
}

func TestUnpackDatagram13RejectsPlaintextWhenCiphertextHeadersEnabled(t *testing.T) {
	plaintext := &PlaintextRecord13{
		Header:  Header{Version: protocol.Version1_2},
		Content: &alert.Alert{Level: alert.Warning, Description: alert.CloseNotify},
	}
	plaintextRaw, err := plaintext.Marshal()
	require.NoError(t, err)

	_, err = UnpackDatagram13(plaintextRaw, 0, true)
	require.ErrorIs(t, err, errInvalidContentType)
}

func TestRecordLayer13Interface(t *testing.T) {
	var plaintext RecordLayer13 = &PlaintextRecord13{}
	require.IsType(t, &Header{}, plaintext.RecordHeader())

	var ciphertext RecordLayer13 = &CiphertextRecord13{}
	require.IsType(t, &UnifiedHeader{}, ciphertext.RecordHeader())
}
