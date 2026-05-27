// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

const (
	gcmTagLength   = 16
	gcmNonceLength = 12
)

// GCM Provides an API to Encrypt/Decrypt DTLS 1.2 Packets.
type GCM struct {
	aead *aead
}

// NewGCM creates a DTLS GCM Cipher.
func NewGCM(localKey, localWriteIV, remoteKey, remoteWriteIV []byte) (*GCM, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encrypt encrypts a DTLS RecordLayer message.
func (g *GCM) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt decrypts a DTLS RecordLayer message.
func (g *GCM) Decrypt(header recordlayer.Header, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
