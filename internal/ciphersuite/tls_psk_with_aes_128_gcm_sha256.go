// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package ciphersuite

import "github.com/pion/dtls/v3/pkg/crypto/clientcertificate"

// TLSPskWithAes128GcmSha256 implements the TLS_PSK_WITH_AES_128_GCM_SHA256 CipherSuite.
type TLSPskWithAes128GcmSha256 struct {
	TLSEcdheEcdsaWithAes128GcmSha256
}

// CertificateType returns what type of certificate this CipherSuite exchanges.
func (c *TLSPskWithAes128GcmSha256) CertificateType() clientcertificate.Type {
	_ = "STUB: not implemented"
	return *new(clientcertificate.Type)
}

// KeyExchangeAlgorithm controls what key exchange algorithm is using during the handshake.
func (c *TLSPskWithAes128GcmSha256) KeyExchangeAlgorithm() KeyExchangeAlgorithm {
	_ = "STUB: not implemented"
	return *new(KeyExchangeAlgorithm)
}

// ID returns the ID of the CipherSuite.
func (c *TLSPskWithAes128GcmSha256) ID() ID { _ = "STUB: not implemented"; return *new(ID) }

func (c *TLSPskWithAes128GcmSha256) String() string { _ = "STUB: not implemented"; return "" }

// AuthenticationType controls what authentication method is using during the handshake.
func (c *TLSPskWithAes128GcmSha256) AuthenticationType() AuthenticationType {
	_ = "STUB: not implemented"
	return *new(AuthenticationType)
}
