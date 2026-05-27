// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package extension implements the extension values in the ClientHello/ServerHello
package extension

import (
	"github.com/pion/dtls/v3/pkg/crypto/signaturehash"
)

// marshalGenericSignatureHashAlgorithm encodes the extension.
// This supports hybrid encoding: TLS 1.3 PSS schemes are encoded as full uint16,
// while TLS 1.2 schemes use hash (high byte) + signature (low byte) encoding.
func marshalGenericSignatureHashAlgorithm(typeValue TypeValue, sigHashAlgs []signaturehash.Algorithm) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// unmarshalGenericSignatureAlgorithm populates the extension from encoded data.
// This supports hybrid encoding: detects TLS 1.3 PSS schemes
// and handles them as full uint16, while TLS 1.2 schemes use byte-split encoding.
func unmarshalGenericSignatureHashAlgorithm(typeValue TypeValue, data []byte, dst *[]signaturehash.Algorithm) error {
	_ = "STUB: not implemented"
	return nil
}
