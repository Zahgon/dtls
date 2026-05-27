// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"crypto/cipher"

	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

const (
	chachaTagLength   = 16
	chachaNonceLength = 12
)

// ChaCha20Poly1305 Provides an API to Encrypt/Decrypt DTLS 1.2 Packets.
//
// Per RFC 7905, ChaCha20-Poly1305 nonce is formed by XOR-ing the write_IV with
// the padded 64-bit sequence number (epoch || sequence_number).
type ChaCha20Poly1305 struct {
	localCipher   cipher.AEAD
	remoteCipher  cipher.AEAD
	localWriteIV  []byte
	remoteWriteIV []byte
}

// NewChaCha20Poly1305 creates a DTLS ChaCha20-Poly1305 Cipher.
func NewChaCha20Poly1305(localKey, localWriteIV, remoteKey, remoteWriteIV []byte) (*ChaCha20Poly1305, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encrypt encrypts a DTLS RecordLayer message.
func (c *ChaCha20Poly1305) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://www.rfc-editor.org/rfc/rfc9325#name-nonce-reuse-in-tls-12

// XOR the last 8 bytes of the nonce with the sequence number

//nolint:gosec

// NOTE: ChaCha20-Poly1305 does NOT include an explicit nonce
// in the record (unlike GCM which includes 8 bytes)

//nolint:gosec

// Decrypt decrypts a DTLS RecordLayer message.
func (c *ChaCha20Poly1305) Decrypt(header recordlayer.Header, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Nothing to decrypt with ChangeCipherSpec

// https://www.rfc-editor.org/rfc/rfc9325#name-nonce-reuse-in-tls-12

// XOR the last 8 bytes of the nonce with the sequence number

//nolint:gosec

// NOTE: ChaCha20-Poly1305 has NO explicit nonce in the record

//nolint:errorlint
