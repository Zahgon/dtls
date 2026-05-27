// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package extension implements the extension values in the ClientHello/ServerHello
package extension

// TypeValue is the 2 byte value for a TLS Extension as registered in the IANA
//
// https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml
type TypeValue uint16

// TypeValue constants.
const (
	ServerNameTypeValue TypeValue = 0
	// In DTLS 1.3, this extension in renamed to "supported_groups".
	SupportedEllipticCurvesTypeValue      TypeValue = 10
	SupportedPointFormatsTypeValue        TypeValue = 11
	SupportedSignatureAlgorithmsTypeValue TypeValue = 13
	UseSRTPTypeValue                      TypeValue = 14
	ALPNTypeValue                         TypeValue = 16
	UseExtendedMasterSecretTypeValue      TypeValue = 23
	PreSharedKeyValue                     TypeValue = 41
	EarlyDataIndicationTypeValue          TypeValue = 42
	SupportedVersionsTypeValue            TypeValue = 43
	CookieTypeValue                       TypeValue = 44
	PskKeyExchangeModesTypeValue          TypeValue = 45
	CertificateAuthoritiesTypeValue       TypeValue = 47
	OIDFiltersTypeValue                   TypeValue = 48
	PostHandshakeAuthTypeValue            TypeValue = 49
	SignatureAlgorithmsCertTypeValue      TypeValue = 50
	KeyShareTypeValue                     TypeValue = 51
	ConnectionIDTypeValue                 TypeValue = 54
	RenegotiationInfoTypeValue            TypeValue = 65281
)

// Extension represents a single TLS extension.
type Extension interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	TypeValue() TypeValue
}

// Unmarshal many extensions at once.
func Unmarshal(buf []byte) ([]Extension, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil, nil
}

//nolint:gosec // offset bounded by loop condition

// Marshal many extensions at once.
func Marshal(e []Extension) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115
