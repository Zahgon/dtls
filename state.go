// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"errors"
	"sync/atomic"

	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/crypto/signaturehash"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
	"github.com/pion/transport/v4/replaydetector"
)

// State holds the dtls connection state and implements both encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler.
type State struct {
	localEpoch, remoteEpoch   atomic.Value
	localSequenceNumber       []uint64 // uint48
	localRandom, remoteRandom handshake.Random
	masterSecret              []byte
	cipherSuite               CipherSuite // nil if a cipherSuite hasn't been chosen
	CipherSuiteID             CipherSuiteID

	remoteSupportsRenegotiation bool // True when Client Hello contained renegotiation extension

	srtpProtectionProfile         atomic.Value // Negotiated SRTPProtectionProfile
	remoteSRTPMasterKeyIdentifier []byte

	PeerCertificates [][]byte
	IdentityHint     []byte
	SessionID        []byte

	// Connection Identifiers must be negotiated afresh on session resumption.
	// https://datatracker.ietf.org/doc/html/rfc9146#name-the-connection_id-extension

	// localConnectionID is the locally generated connection ID that is expected
	// to be received from the remote endpoint.
	// For a server, this is the connection ID sent in ServerHello.
	// For a client, this is the connection ID sent in the ClientHello.
	localConnectionID atomic.Value
	// remoteConnectionID is the connection ID that the remote endpoint
	// specifies should be sent.
	// For a server, this is the connection ID received in the ClientHello.
	// For a client, this is the connection ID received in the ServerHello.
	remoteConnectionID []byte

	isClient bool

	preMasterSecret      []byte
	extendedMasterSecret bool

	namedCurve                 elliptic.Curve
	localKeypair               *elliptic.Keypair
	cookie                     []byte
	handshakeSendSequence      int
	handshakeRecvSequence      int
	serverName                 string
	remoteCertRequestAlgs      []signaturehash.Algorithm
	remoteCertSignatureSchemes []signaturehash.Algorithm // signature_algorithms_cert from peer
	remoteRequestedCertificate bool                      // Did we get a CertificateRequest
	localCertificatesVerify    []byte                    // cache CertificateVerify
	localVerifyData            []byte                    // cached VerifyData
	localKeySignature          []byte                    // cached keySignature
	peerCertificatesVerified   bool

	replayDetector []replaydetector.ReplayDetector

	peerSupportedProtocols []string
	NegotiatedProtocol     string
}

type serializedState struct {
	LocalEpoch            uint16
	RemoteEpoch           uint16
	LocalRandom           [handshake.RandomLength]byte
	RemoteRandom          [handshake.RandomLength]byte
	CipherSuiteID         uint16
	MasterSecret          []byte
	SequenceNumber        uint64
	SRTPProtectionProfile uint16
	PeerCertificates      [][]byte
	IdentityHint          []byte
	SessionID             []byte
	LocalConnectionID     []byte
	RemoteConnectionID    []byte
	IsClient              bool
	NegotiatedProtocol    string
}

var errCipherSuiteNotSet = &InternalError{Err: errors.New("cipher suite not set")} //nolint:err113

func (s *State) clone() (*State, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *State) serialize() (*serializedState, error) { _ = "STUB: not implemented"; return nil, nil }

// Marshal random values

func (s *State) deserialize(serialized serializedState) {
	_ = "STUB: not implemented"
	// Set epoch values
	return
}

// Set random values

// Set master secret

// Set cipher suite

// Set remote certificate

// Set local and remote connection IDs

func (s *State) initCipherSuite() error { _ = "STUB: not implemented"; return nil }

// MarshalBinary is a binary.BinaryMarshaler.MarshalBinary implementation.
func (s *State) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary is a binary.BinaryUnmarshaler.UnmarshalBinary implementation.
func (s *State) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// ExportKeyingMaterial returns length bytes of exported key material in a new
// slice as defined in RFC 5705.
// This allows protocols to use DTLS for key establishment, but
// then use some of the keying material for their own purposes.
func (s *State) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) getRemoteEpoch() uint16 { _ = "STUB: not implemented"; return 0 }

func (s *State) getLocalEpoch() uint16 { _ = "STUB: not implemented"; return 0 }

func (s *State) setSRTPProtectionProfile(profile SRTPProtectionProfile) {
	_ = "STUB: not implemented"
	return
}

func (s *State) getSRTPProtectionProfile() SRTPProtectionProfile {
	_ = "STUB: not implemented"
	return *new(SRTPProtectionProfile)
}

func (s *State) getLocalConnectionID() []byte { _ = "STUB: not implemented"; return nil }

func (s *State) setLocalConnectionID(v []byte) { _ = "STUB: not implemented"; return }

// RemoteRandomBytes returns the remote client hello random bytes.
func (s *State) RemoteRandomBytes() [handshake.RandomBytesLength]byte {
	_ = "STUB: not implemented"
	return nil
}
