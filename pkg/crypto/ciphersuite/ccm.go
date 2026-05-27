// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// CCMTagLen is the length of Authentication Tag.
type CCMTagLen int

// CCM Enums.
const (
	CCMTagLength8  CCMTagLen = 8
	CCMTagLength   CCMTagLen = 16
	ccmNonceLength           = 12
)

// CCM Provides an API to Encrypt/Decrypt DTLS 1.2 Packets.
type CCM struct {
	aead *aead
}

// NewCCM creates a DTLS GCM Cipher.
func NewCCM(tagLen CCMTagLen, localKey, localWriteIV, remoteKey, remoteWriteIV []byte) (*CCM, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encrypt encrypt a DTLS RecordLayer message.
func (c *CCM) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt decrypts a DTLS RecordLayer message.
func (c *CCM) Decrypt(header recordlayer.Header, in []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
