// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"hash"
	"sync/atomic"

	"github.com/pion/dtls/v3/pkg/crypto/ciphersuite"
	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
	"github.com/pion/dtls/v3/pkg/crypto/prf"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// TLSEcdheEcdsaWithAes128GcmSha256  represents a TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 CipherSuite.
type TLSEcdheEcdsaWithAes128GcmSha256 struct {
	gcm atomic.Value // *cryptoGCM
}

// CertificateType returns what type of certficate this CipherSuite exchanges.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// KeyExchangeAlgorithm controls what key exchange algorithm is using during the handshake.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) KeyExchangeAlgorithm() KeyExchangeAlgorithm {
	_ = "STUB: not implemented"
	return *new(KeyExchangeAlgorithm)
}

// ECC uses Elliptic Curve Cryptography.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) ECC() bool {
	_ = "STUB: not implemented"

	// ID returns the ID of the CipherSuite.
	return false
}

func (c *TLSEcdheEcdsaWithAes128GcmSha256) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSEcdheEcdsaWithAes128GcmSha256) String() string { _ = "STUB: not implemented"; return "" }

// HashFunc returns the hashing func for this CipherSuite.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) HashFunc() func() hash.Hash {
	_ = "STUB: not implemented"

	// AuthenticationType controls what authentication method is using during the handshake.
	return nil
}

func (c *TLSEcdheEcdsaWithAes128GcmSha256) AuthenticationType() AuthenticationType {
	_ = "STUB: not implemented"
	return *new(AuthenticationType)
}

// IsInitialized returns if the CipherSuite has keying material and can
// encrypt/decrypt packets.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) IsInitialized() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *TLSEcdheEcdsaWithAes128GcmSha256) init(
	masterSecret, clientRandom, serverRandom []byte,
	isClient bool,
	prfMacLen, prfKeyLen, prfIvLen int,
	hashFunc func() hash.Hash,
) error {
	keys, err := prf.GenerateEncryptionKeys(
		masterSecret, clientRandom, serverRandom, prfMacLen, prfKeyLen, prfIvLen, hashFunc,
	)
	if err != nil {
		return err
	}

	var gcm *ciphersuite.GCM
	if isClient {
		gcm, err = ciphersuite.NewGCM(keys.ClientWriteKey, keys.ClientWriteIV, keys.ServerWriteKey, keys.ServerWriteIV)
	} else {
		gcm, err = ciphersuite.NewGCM(keys.ServerWriteKey, keys.ServerWriteIV, keys.ClientWriteKey, keys.ClientWriteIV)
	}
	c.gcm.Store(gcm)

	return err
}

// Init initializes the internal Cipher with keying material.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Encrypt encrypts a single TLS RecordLayer.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt decrypts a single TLS RecordLayer.
func (c *TLSEcdheEcdsaWithAes128GcmSha256) Decrypt(h recordlayer.Header, raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
