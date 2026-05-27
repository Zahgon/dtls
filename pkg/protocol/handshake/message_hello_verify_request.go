// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"github.com/pion/dtls/v3/pkg/protocol"
)

// MessageHelloVerifyRequest is as follows:
//
//	struct {
//	  ProtocolVersion server_version;
//	  opaque cookie<0..2^8-1>;
//	} HelloVerifyRequest;
//
//	The HelloVerifyRequest message type is hello_verify_request(3).
//
//	When the client sends its ClientHello message to the server, the server
//	MAY respond with a HelloVerifyRequest message.  This message contains
//	a stateless cookie generated using the technique of [PHOTURIS].  The
//	client MUST retransmit the ClientHello with the cookie added.
//
//	https://tools.ietf.org/html/rfc6347#section-4.2.1
type MessageHelloVerifyRequest struct {
	Version protocol.Version
	Cookie  []byte
}

// Type returns the Handshake Type.
func (m MessageHelloVerifyRequest) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Marshal encodes the Handshake.
func (m *MessageHelloVerifyRequest) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G115: cookie length is validated to be <= 255 above.

// Unmarshal populates the message from encoded data.
func (m *MessageHelloVerifyRequest) Unmarshal(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
