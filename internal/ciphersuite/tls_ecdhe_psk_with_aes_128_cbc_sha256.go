// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"hash"
	"sync/atomic"

	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// TLSEcdhePskWithAes128CbcSha256 implements the TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256 CipherSuite.
type TLSEcdhePskWithAes128CbcSha256 struct {
	cbc atomic.Value // *cryptoCBC
}

// NewTLSEcdhePskWithAes128CbcSha256 creates TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256 cipher.
func NewTLSEcdhePskWithAes128CbcSha256() *TLSEcdhePskWithAes128CbcSha256 {
	_ = "STUB: not implemented"
	return nil
}

// CertificateType returns what type of certificate this CipherSuite exchanges.
func (c *TLSEcdhePskWithAes128CbcSha256) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// KeyExchangeAlgorithm controls what key exchange algorithm is using during the handshake.
func (c *TLSEcdhePskWithAes128CbcSha256) KeyExchangeAlgorithm() KeyExchangeAlgorithm {
	_ = "STUB: not implemented"
	return *new(KeyExchangeAlgorithm)
}

// ECC uses Elliptic Curve Cryptography.
func (c *TLSEcdhePskWithAes128CbcSha256) ECC() bool {
	_ = "STUB: not implemented"

	// ID returns the ID of the CipherSuite.
	return false
}

func (c *TLSEcdhePskWithAes128CbcSha256) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSEcdhePskWithAes128CbcSha256) String() string { _ = "STUB: not implemented"; return "" }

// HashFunc returns the hashing func for this CipherSuite.
func (c *TLSEcdhePskWithAes128CbcSha256) HashFunc() func() hash.Hash {
	_ = "STUB: not implemented"

	// AuthenticationType controls what authentication method is using during the handshake.
	return nil
}

func (c *TLSEcdhePskWithAes128CbcSha256) AuthenticationType() AuthenticationType {
	_ = "STUB: not implemented"
	return *new(AuthenticationType)
}

// IsInitialized returns if the CipherSuite has keying material and can
// encrypt/decrypt packets.
func (c *TLSEcdhePskWithAes128CbcSha256) IsInitialized() bool {
	_ = "STUB: not implemented"
	return false

	// Init initializes the internal Cipher with keying material.
}

func (c *TLSEcdhePskWithAes128CbcSha256) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Encrypt encrypts a single TLS RecordLayer.
func (c *TLSEcdhePskWithAes128CbcSha256) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// !c.isInitialized()

// Decrypt decrypts a single TLS RecordLayer.
func (c *TLSEcdhePskWithAes128CbcSha256) Decrypt(h recordlayer.Header, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// !c.isInitialized()
