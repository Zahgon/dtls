// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

/*
  DTLS messages are grouped into a series of message flights, according
  to the diagrams below.  Although each flight of messages may consist
  of a number of messages, they should be viewed as monolithic for the
  purpose of timeout and retransmission.
  https://tools.ietf.org/html/rfc4347#section-4.2.4

  Message flights for full handshake:

  Client                                          Server
  ------                                          ------
                                      Waiting                 Flight 0

  ClientHello             -------->                           Flight 1

                          <-------    HelloVerifyRequest      Flight 2

  ClientHello              -------->                           Flight 3

                                             ServerHello    \
                                            Certificate*     \
                                      ServerKeyExchange*      Flight 4
                                     CertificateRequest*     /
                          <--------      ServerHelloDone    /

  Certificate*                                              \
  ClientKeyExchange                                          \
  CertificateVerify*                                          Flight 5
  [ChangeCipherSpec]                                         /
  Finished                -------->                         /

                                      [ChangeCipherSpec]    \ Flight 6
                          <--------             Finished    /

  Message flights for session-resuming handshake (no cookie exchange):

  Client                                          Server
  ------                                          ------
                                      Waiting                 Flight 0

  ClientHello             -------->                           Flight 1

                                             ServerHello    \
                                      [ChangeCipherSpec]      Flight 4b
                          <--------             Finished    /

  [ChangeCipherSpec]                                        \ Flight 5b
  Finished                -------->                         /

                                      [ChangeCipherSpec]    \ Flight 6
                          <--------             Finished    /
*/

type flightVal uint8

const (
	flight0 flightVal = iota + 1
	flight1
	flight2
	flight3
	flight4
	flight4b
	flight5
	flight5b
	flight6
)

func (f flightVal) String() string {
	_ = "STUB: not implemented" //nolint:cyclop
	return ""
}

func (f flightVal) isLastSendFlight() bool { _ = "STUB: not implemented"; return false }

func (f flightVal) isLastRecvFlight() bool { _ = "STUB: not implemented"; return false }
