// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"github.com/pion/dtls/v3/pkg/crypto/ciphersuite"
	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
)

// Aes256Ccm is a base class used by multiple AES-CCM Ciphers.
type Aes256Ccm struct {
	AesCcm
}

func newAes256Ccm(
	clientCertificateType clientcertificate.Type,
	id ID,
	psk bool,
	cryptoCCMTagLen ciphersuite.CCMTagLen,
	keyExchangeAlgorithm KeyExchangeAlgorithm,
	ecc bool,
) *Aes256Ccm {
	_ = "STUB: not implemented"
	return nil
}

// Init initializes the internal Cipher with keying material.
func (c *Aes256Ccm) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}
