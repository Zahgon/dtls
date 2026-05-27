// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import "github.com/pion/dtls/v3/pkg/crypto/clientcertificate"

// TLSEcdheRsaWithAes128GcmSha256 implements the TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 CipherSuite.
type TLSEcdheRsaWithAes128GcmSha256 struct {
	TLSEcdheEcdsaWithAes128GcmSha256
}

// CertificateType returns what type of certificate this CipherSuite exchanges.
func (c *TLSEcdheRsaWithAes128GcmSha256) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// ID returns the ID of the CipherSuite.
func (c *TLSEcdheRsaWithAes128GcmSha256) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSEcdheRsaWithAes128GcmSha256) String() string { _ = "STUB: not implemented"; return "" }
