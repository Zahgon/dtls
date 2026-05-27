// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"hash"
	"sync/atomic"

	"github.com/pion/dtls/v3/pkg/crypto/ciphersuite"
	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// AesCcm is a base class used by multiple AES-CCM Ciphers.
type AesCcm struct {
	ccm                   atomic.Value // *cryptoCCM
	clientCertificateType clientcertificate.Type
	id                    ID
	psk                   bool
	keyExchangeAlgorithm  KeyExchangeAlgorithm
	cryptoCCMTagLen       ciphersuite.CCMTagLen
	ecc                   bool
}

// CertificateType returns what type of certificate this CipherSuite exchanges.
func (c *AesCcm) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// ID returns the ID of the CipherSuite.
func (c *AesCcm) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *AesCcm) String() string { _ = "STUB: not implemented"; return "" }

// ECC uses Elliptic Curve Cryptography.
func (c *AesCcm) ECC() bool {
	_ = "STUB: not implemented"

	// KeyExchangeAlgorithm controls what key exchange algorithm is using during the handshake.
	return false
}

func (c *AesCcm) KeyExchangeAlgorithm() KeyExchangeAlgorithm {
	_ = "STUB: not implemented"
	return *new(KeyExchangeAlgorithm)
}

// HashFunc returns the hashing func for this CipherSuite.
func (c *AesCcm) HashFunc() func() hash.Hash {
	_ = "STUB: not implemented"

	// AuthenticationType controls what authentication method is using during the handshake.
	return nil
}

func (c *AesCcm) AuthenticationType() AuthenticationType {
	_ = "STUB: not implemented"
	return *new(AuthenticationType)
}

// IsInitialized returns if the CipherSuite has keying material and can
// encrypt/decrypt packets.
func (c *AesCcm) IsInitialized() bool { _ = "STUB: not implemented"; return false }

// Init initializes the internal Cipher with keying material.
func (c *AesCcm) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool, prfKeyLen int) error {
	_ = "STUB: not implemented"
	return nil
}

// Encrypt encrypts a single TLS RecordLayer.
func (c *AesCcm) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt decrypts a single TLS RecordLayer.
func (c *AesCcm) Decrypt(h recordlayer.Header, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
