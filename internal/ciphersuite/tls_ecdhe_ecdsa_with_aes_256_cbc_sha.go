// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	//nolint: gosec,gci

	"hash"
	"sync/atomic"

	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// TLSEcdheEcdsaWithAes256CbcSha represents a TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA CipherSuite.
type TLSEcdheEcdsaWithAes256CbcSha struct {
	cbc atomic.Value // *cryptoCBC
}

// CertificateType returns what type of certficate this CipherSuite exchanges.
func (c *TLSEcdheEcdsaWithAes256CbcSha) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// KeyExchangeAlgorithm controls what key exchange algorithm is using during the handshake.
func (c *TLSEcdheEcdsaWithAes256CbcSha) KeyExchangeAlgorithm() KeyExchangeAlgorithm {
	_ = "STUB: not implemented"
	return *new(KeyExchangeAlgorithm)
}

// ECC uses Elliptic Curve Cryptography.
func (c *TLSEcdheEcdsaWithAes256CbcSha) ECC() bool {
	_ = "STUB: not implemented"

	// ID returns the ID of the CipherSuite.
	return false
}

func (c *TLSEcdheEcdsaWithAes256CbcSha) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSEcdheEcdsaWithAes256CbcSha) String() string { _ = "STUB: not implemented"; return "" }

// HashFunc returns the hashing func for this CipherSuite.
func (c *TLSEcdheEcdsaWithAes256CbcSha) HashFunc() func() hash.Hash {
	_ = "STUB: not implemented"

	// AuthenticationType controls what authentication method is using during the handshake.
	return nil
}

func (c *TLSEcdheEcdsaWithAes256CbcSha) AuthenticationType() AuthenticationType {
	_ = "STUB: not implemented"
	return *new(AuthenticationType)
}

// IsInitialized returns if the CipherSuite has keying material and can
// encrypt/decrypt packets.
func (c *TLSEcdheEcdsaWithAes256CbcSha) IsInitialized() bool {
	_ = "STUB: not implemented"
	return false

	// Init initializes the internal Cipher with keying material.
}

func (c *TLSEcdheEcdsaWithAes256CbcSha) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Encrypt encrypts a single TLS RecordLayer.
func (c *TLSEcdheEcdsaWithAes256CbcSha) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt decrypts a single TLS RecordLayer.
func (c *TLSEcdheEcdsaWithAes256CbcSha) Decrypt(h recordlayer.Header, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
