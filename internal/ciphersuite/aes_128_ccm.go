// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"github.com/pion/dtls/v3/pkg/crypto/ciphersuite"
	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
)

// Aes128Ccm is a base class used by multiple AES-CCM Ciphers.
type Aes128Ccm struct {
	AesCcm
}

func newAes128Ccm(
	clientCertificateType clientcertificate.Type,
	id ID,
	psk bool,
	cryptoCCMTagLen ciphersuite.CCMTagLen,
	keyExchangeAlgorithm KeyExchangeAlgorithm,
	ecc bool,
) *Aes128Ccm {
	_ = "STUB: not implemented"
	return nil
}

// Init initializes the internal Cipher with keying material.
func (c *Aes128Ccm) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}
