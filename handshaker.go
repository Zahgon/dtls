// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"sync"
	"time"

	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/crypto/signaturehash"
	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/alert"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
	"github.com/pion/logging"
)

// [RFC6347 Section-4.2.4]
//                      +-----------+
//                +---> | PREPARING | <--------------------+
//                |     +-----------+                      |
//                |           |                            |
//                |           | Buffer next flight         |
//                |           |                            |
//                |          \|/                           |
//                |     +-----------+                      |
//                |     |  SENDING  |<------------------+  | Send
//                |     +-----------+                   |  | HelloRequest
//        Receive |           |                         |  |
//           next |           | Send flight             |  | or
//         flight |  +--------+                         |  |
//                |  |        | Set retransmit timer    |  | Receive
//                |  |       \|/                        |  | HelloRequest
//                |  |  +-----------+                   |  | Send
//                +--)--|  WAITING  |-------------------+  | ClientHello
//                |  |  +-----------+   Timer expires   |  |
//                |  |         |                        |  |
//                |  |         +------------------------+  |
//        Receive |  | Send           Read retransmit      |
//           last |  | last                                |
//         flight |  | flight                              |
//                |  |                                     |
//               \|/\|/                                    |
//            +-----------+                                |
//            | FINISHED  | -------------------------------+
//            +-----------+
//                 |  /|\
//                 |   |
//                 +---+
//              Read retransmit
//           Retransmit last flight

type handshakeState uint8

const (
	handshakeErrored handshakeState = iota
	handshakePreparing
	handshakeSending
	handshakeWaiting
	handshakeFinished
)

func (s handshakeState) String() string { _ = "STUB: not implemented"; return "" }

type handshakeFSM struct {
	currentFlight      flightVal
	flights            []*packet
	retransmit         bool
	retransmitInterval time.Duration
	state              *State
	cache              *handshakeCache
	cfg                *handshakeConfig
	closed             chan struct{}
}

type handshakeConfig struct {
	localPSKCallback             PSKCallback
	localPSKIdentityHint         []byte
	localCipherSuites            []CipherSuite             // Available CipherSuites
	localSignatureSchemes        []signaturehash.Algorithm // Available signature schemes
	localCertSignatureSchemes    []signaturehash.Algorithm // Available signature schemes for certificates
	extendedMasterSecret         ExtendedMasterSecretType  // Policy for the Extended Master Support extension
	localSRTPProtectionProfiles  []SRTPProtectionProfile   // Available SRTPProtectionProfiles, if empty no SRTP support
	localSRTPMasterKeyIdentifier []byte
	serverName                   string
	supportedProtocols           []string
	clientAuth                   ClientAuthType // If we are a client should we request a client certificate
	localCertificates            []tls.Certificate
	nameToCertificate            map[string]*tls.Certificate
	insecureSkipVerify           bool
	verifyPeerCertificate        func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error
	verifyConnection             func(*State) error
	sessionStore                 SessionStore
	rootCAs                      *x509.CertPool
	clientCAs                    *x509.CertPool
	initialRetransmitInterval    time.Duration
	disableRetransmitBackoff     bool
	customCipherSuites           func() []CipherSuite
	ellipticCurves               []elliptic.Curve
	insecureSkipHelloVerify      bool
	connectionIDGenerator        func() []byte
	helloRandomBytesGenerator    func() [handshake.RandomBytesLength]byte

	onFlightState func(flightVal, handshakeState)
	log           logging.LeveledLogger
	keyLogWriter  io.Writer

	localGetCertificate       func(*ClientHelloInfo) (*tls.Certificate, error)
	localGetClientCertificate func(*CertificateRequestInfo) (*tls.Certificate, error)

	initialEpoch uint16

	mu sync.Mutex

	clientHelloMessageHook        func(handshake.MessageClientHello) handshake.Message
	serverHelloMessageHook        func(handshake.MessageServerHello) handshake.Message
	certificateRequestMessageHook func(handshake.MessageCertificateRequest) handshake.Message

	resumeState *State

	minVersion protocol.Version
	maxVersion protocol.Version
}

type flightConn interface {
	notify(ctx context.Context, level alert.Level, desc alert.Description) error
	writePackets(context.Context, []*packet) error
	recvHandshake() <-chan recvHandshakeState
	setLocalEpoch(epoch uint16)
	handleQueuedPackets(context.Context) error
	sessionKey() []byte
}

func (c *handshakeConfig) writeKeyLog(label string, clientRandom, secret []byte) {
	_ = "STUB: not implemented"
	return
}

func srvCliStr(isClient bool) string { _ = "STUB: not implemented"; return "" }

func newHandshakeFSM(
	s *State, cache *handshakeCache, cfg *handshakeConfig,
	initialFlight flightVal,
) *handshakeFSM {
	_ = "STUB: not implemented"
	return nil
}

func (s *handshakeFSM) Run(ctx context.Context, conn flightConn, initialState handshakeState) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *handshakeFSM) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (s *handshakeFSM) prepare(ctx context.Context, conn flightConn) (handshakeState, error) {
	_ = "STUB: not implemented"

	// Prepare flights
	return *new(handshakeState), nil
}

//nolint:gosec // G115

func (s *handshakeFSM) send(ctx context.Context, c flightConn) (handshakeState, error) {
	_ = "STUB: not implemented"
	// Send flights
	return *new(handshakeState), nil
}

func (s *handshakeFSM) wait(ctx context.Context, conn flightConn) (handshakeState, error) {
	_ = "STUB: not implemented" //nolint:gocognit,cyclop
	return *new(handshakeState), nil
}

// only reset retransmit interval on non-retransmit state
// https://github.com/pion/dtls/issues/758

// RFC 4347 4.2.4.1:
// Implementations SHOULD use an initial timer value of 1 second (the minimum defined in RFC 2988 [RFC2988])
// and double the value at each retransmission, up to no less than the RFC 2988 maximum of 60 seconds.

func (s *handshakeFSM) finish(ctx context.Context, c flightConn) (handshakeState, error) {
	_ = "STUB: not implemented"
	return *new(handshakeState), nil
}
