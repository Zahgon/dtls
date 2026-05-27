// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import "github.com/pion/dtls/v3/pkg/crypto/clientcertificate"

// TLSEcdheRsaWithAes256CbcSha implements the TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA CipherSuite.
type TLSEcdheRsaWithAes256CbcSha struct {
	TLSEcdheEcdsaWithAes256CbcSha
}

// CertificateType returns what type of certificate this CipherSuite exchanges.
func (c *TLSEcdheRsaWithAes256CbcSha) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// ID returns the ID of the CipherSuite.
func (c *TLSEcdheRsaWithAes256CbcSha) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSEcdheRsaWithAes256CbcSha) String() string { _ = "STUB: not implemented"; return "" }
