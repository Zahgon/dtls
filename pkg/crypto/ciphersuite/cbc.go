// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"crypto/cipher"
	"hash"

	"github.com/pion/dtls/v3/pkg/crypto/prf"
	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// block ciphers using cipher block chaining.
type cbcMode interface {
	cipher.BlockMode
	SetIV([]byte)
}

// CBC Provides an API to Encrypt/Decrypt DTLS 1.2 Packets.
type CBC struct {
	writeCBC, readCBC cbcMode
	writeMac, readMac []byte
	h                 prf.HashFunc
}

// NewCBC creates a DTLS CBC Cipher.
func NewCBC(
	localKey, localWriteIV, localMac, remoteKey, remoteWriteIV, remoteMac []byte,
	hashFunc prf.HashFunc,
) (*CBC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encrypt encrypt a DTLS RecordLayer message.
func (c *CBC) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Generate + Append MAC

// Generate + Append padding

//nolint:gosec //G115

// Generate IV

// Set IV + Encrypt + Prepend IV

//nolint:makezero // todo: FIX

// Prepend unencrypted header with encrypted payload

// Update recordLayer size to include IV+MAC+Padding
//nolint:gosec //G115

// Decrypt decrypts a DTLS RecordLayer message.
func (c *CBC) Decrypt(header recordlayer.Header, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Nothing to encrypt with ChangeCipherSpec

// Set + remove per record IV

// Decrypt

// Padding+MAC needs to be checked in constant time
// Otherwise we reveal information about the level of correctness

// Compute Local MAC and compare

func (c *CBC) hmac(
	epoch uint16,
	sequenceNumber uint64,
	contentType protocol.ContentType,
	protocolVersion protocol.Version,
	payload []byte,
	key []byte,
	hf func() hash.Hash,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec //G115

// hmacCID calculates a MAC according to
// https://datatracker.ietf.org/doc/html/rfc9146#section-5.1
func (c *CBC) hmacCID(
	epoch uint16,
	sequenceNumber uint64,
	protocolVersion protocol.Version,
	payload []byte,
	key []byte,
	hf func() hash.Hash,
	cid []byte,
) ([]byte, error) {
	_ = "STUB: not implemented"
	// Must unmarshal inner plaintext in orde to perform MAC.
	return nil, nil
}

//nolint:gosec //G115

//nolint:gosec //G115
