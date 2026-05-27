// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package handshake provides the DTLS wire protocol for handshakes
package handshake

import (
	"github.com/pion/dtls/v3/internal/ciphersuite/types"
	"github.com/pion/dtls/v3/pkg/protocol"
)

// Type is the unique identifier for each handshake message
// https://tools.ietf.org/html/rfc5246#section-7.4
type Type uint8

// Types of DTLS Handshake messages we know about.
const (
	TypeHelloRequest       Type = 0
	TypeClientHello        Type = 1
	TypeServerHello        Type = 2
	TypeHelloVerifyRequest Type = 3
	TypeCertificate        Type = 11
	TypeServerKeyExchange  Type = 12
	TypeCertificateRequest Type = 13
	TypeServerHelloDone    Type = 14
	TypeCertificateVerify  Type = 15
	TypeClientKeyExchange  Type = 16
	TypeFinished           Type = 20
)

// String returns the string representation of this type.
func (t Type) String() string {
	_ = "STUB: not implemented" //nolint:cyclop
	return ""
}

// Message is the body of a Handshake datagram.
type Message interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	Type() Type
}

// Handshake protocol is responsible for selecting a cipher spec and
// generating a master secret, which together comprise the primary
// cryptographic parameters associated with a secure session.  The
// handshake protocol can also optionally authenticate parties who have
// certificates signed by a trusted certificate authority.
// https://tools.ietf.org/html/rfc5246#section-7.3
type Handshake struct {
	Header  Header
	Message Message

	KeyExchangeAlgorithm types.KeyExchangeAlgorithm
}

// ContentType returns what kind of content this message is carying.
func (h Handshake) ContentType() protocol.ContentType {
	_ = "STUB: not implemented"
	return *new(protocol.ContentType)
}

// Marshal encodes a handshake into a binary message.
func (h *Handshake) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec // G115

// Unmarshal decodes a handshake from a binary message.
func (h *Handshake) Unmarshal(data []byte) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

//nolint:gosec // G115
