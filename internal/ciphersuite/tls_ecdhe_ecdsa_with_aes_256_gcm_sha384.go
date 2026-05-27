// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import (
	"hash"
)

// TLSEcdheEcdsaWithAes256GcmSha384  represents a TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 CipherSuite.
type TLSEcdheEcdsaWithAes256GcmSha384 struct {
	TLSEcdheEcdsaWithAes128GcmSha256
}

// ID returns the ID of the CipherSuite.
func (c *TLSEcdheEcdsaWithAes256GcmSha384) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSEcdheEcdsaWithAes256GcmSha384) String() string { _ = "STUB: not implemented"; return "" }

// HashFunc returns the hashing func for this CipherSuite.
func (c *TLSEcdheEcdsaWithAes256GcmSha384) HashFunc() func() hash.Hash {
	_ = "STUB: not implemented"
	return nil

	// Init initializes the internal Cipher with keying material.
}

func (c *TLSEcdheEcdsaWithAes256GcmSha384) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}
