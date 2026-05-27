// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"crypto/tls"
	"hash"

	"github.com/pion/dtls/v3/internal/ciphersuite"
	"github.com/pion/dtls/v3/pkg/crypto/clientcertificate"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// CipherSuiteID is an ID for our supported CipherSuites.
type CipherSuiteID = ciphersuite.ID

// Supported Cipher Suites.
const (

	// nolint: godot
	// AES-128-CCM
	TLS_ECDHE_ECDSA_WITH_AES_128_CCM   CipherSuiteID = ciphersuite.TLS_ECDHE_ECDSA_WITH_AES_128_CCM   // nolint: revive,staticcheck,lll
	TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8 CipherSuiteID = ciphersuite.TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8 // nolint: revive,staticcheck,lll

	// nolint: godot
	// AES-128-GCM-SHA256
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 CipherSuiteID = ciphersuite.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 // nolint: revive,staticcheck,lll
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256   CipherSuiteID = ciphersuite.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256   // nolint: revive,staticcheck,lll

	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 CipherSuiteID = ciphersuite.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 // nolint: revive,staticcheck,lll
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384   CipherSuiteID = ciphersuite.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384   // nolint: revive,staticcheck,lll

	// nolint: godot
	// AES-256-CBC-SHA
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA CipherSuiteID = ciphersuite.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA // nolint: revive,staticcheck,lll
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA   CipherSuiteID = ciphersuite.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA   // nolint: revive,staticcheck,lll

	TLS_PSK_WITH_AES_128_CCM        CipherSuiteID = ciphersuite.TLS_PSK_WITH_AES_128_CCM        // nolint: revive,staticcheck,lll
	TLS_PSK_WITH_AES_128_CCM_8      CipherSuiteID = ciphersuite.TLS_PSK_WITH_AES_128_CCM_8      // nolint: revive,staticcheck,lll
	TLS_PSK_WITH_AES_256_CCM_8      CipherSuiteID = ciphersuite.TLS_PSK_WITH_AES_256_CCM_8      // nolint: revive,staticcheck,lll
	TLS_PSK_WITH_AES_128_GCM_SHA256 CipherSuiteID = ciphersuite.TLS_PSK_WITH_AES_128_GCM_SHA256 // nolint: revive,staticcheck,lll
	TLS_PSK_WITH_AES_128_CBC_SHA256 CipherSuiteID = ciphersuite.TLS_PSK_WITH_AES_128_CBC_SHA256 // nolint: revive,staticcheck,lll

	TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256 CipherSuiteID = ciphersuite.TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256 // nolint: revive,staticcheck,lll

	// nolint: godot
	// ChaCha20-Poly1305-SHA256
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 CipherSuiteID = ciphersuite.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 // nolint: revive,staticcheck,lll
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   CipherSuiteID = ciphersuite.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   // nolint: revive,staticcheck,lll
	TLS_PSK_WITH_CHACHA20_POLY1305_SHA256         CipherSuiteID = ciphersuite.TLS_PSK_WITH_CHACHA20_POLY1305_SHA256         // nolint: revive,staticcheck,lll
)

// CipherSuiteAuthenticationType controls what authentication method is using during the handshake for a CipherSuite.
type CipherSuiteAuthenticationType = ciphersuite.AuthenticationType

// AuthenticationType Enums.
const (
	CipherSuiteAuthenticationTypeCertificate  CipherSuiteAuthenticationType = ciphersuite.AuthenticationTypeCertificate
	CipherSuiteAuthenticationTypePreSharedKey CipherSuiteAuthenticationType = ciphersuite.AuthenticationTypePreSharedKey
	CipherSuiteAuthenticationTypeAnonymous    CipherSuiteAuthenticationType = ciphersuite.AuthenticationTypeAnonymous
)

// CipherSuiteKeyExchangeAlgorithm controls what exchange algorithm is using during the handshake for a CipherSuite.
type CipherSuiteKeyExchangeAlgorithm = ciphersuite.KeyExchangeAlgorithm

// CipherSuiteKeyExchangeAlgorithm Bitmask.
const (
	CipherSuiteKeyExchangeAlgorithmNone  CipherSuiteKeyExchangeAlgorithm = ciphersuite.KeyExchangeAlgorithmNone
	CipherSuiteKeyExchangeAlgorithmPsk   CipherSuiteKeyExchangeAlgorithm = ciphersuite.KeyExchangeAlgorithmPsk
	CipherSuiteKeyExchangeAlgorithmEcdhe CipherSuiteKeyExchangeAlgorithm = ciphersuite.KeyExchangeAlgorithmEcdhe
)

var _ = allCipherSuites() // Necessary until this function isn't only used by Go 1.14

// CipherSuite is an interface that all DTLS CipherSuites must satisfy.
type CipherSuite interface {
	// String of CipherSuite, only used for logging
	String() string

	// ID of CipherSuite.
	ID() CipherSuiteID

	// What type of Certificate does this CipherSuite use
	CertificateType() clientcertificate.Type

	// What Hash function is used during verification
	HashFunc() func() hash.Hash

	// AuthenticationType controls what authentication method is using during the handshake
	AuthenticationType() CipherSuiteAuthenticationType

	// KeyExchangeAlgorithm controls what exchange algorithm is using during the handshake
	KeyExchangeAlgorithm() CipherSuiteKeyExchangeAlgorithm

	// ECC (Elliptic Curve Cryptography) determines whether ECC extesions will be send during handshake.
	// https://datatracker.ietf.org/doc/html/rfc4492#page-10
	ECC() bool

	// Called when keying material has been generated, should initialize the internal cipher
	Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error
	IsInitialized() bool
	Encrypt(pkt *recordlayer.RecordLayer, raw []byte) ([]byte, error)
	Decrypt(h recordlayer.Header, in []byte) ([]byte, error)
}

// CipherSuiteName provides the same functionality as tls.CipherSuiteName
// that appeared first in Go 1.14.
//
// Our implementation differs slightly in that it takes in a CiperSuiteID,
// like the rest of our library, instead of a uint16 like crypto/tls.
func CipherSuiteName(id CipherSuiteID) string { _ = "STUB: not implemented"; return "" }

// Taken from https://www.iana.org/assignments/tls-parameters/tls-parameters.xml
// A cipherSuite is a specific combination of key agreement, cipher and MAC
// function.
func cipherSuiteForID(id CipherSuiteID, customCiphers func() []CipherSuite) CipherSuite {
	_ = "STUB: not implemented" //nolint:cyclop
	return *
	//nolint:exhaustive
	new(CipherSuite)
}

// CipherSuites we support in order of preference.
func defaultCipherSuites() []CipherSuite { _ = "STUB: not implemented"; return nil }

func allCipherSuites() []CipherSuite { _ = "STUB: not implemented"; return nil }

func cipherSuiteIDs(cipherSuites []CipherSuite) []uint16 { _ = "STUB: not implemented"; return nil }

//nolint:cyclop
func parseCipherSuites(
	userSelectedSuites []CipherSuiteID,
	customCipherSuites func() []CipherSuite,
	includeCertificateSuites, includePSKSuites bool,
) ([]CipherSuite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put CustomCipherSuites before ID selected suites

func filterCipherSuitesForCertificate(cert *tls.Certificate, cipherSuites []CipherSuite) []CipherSuite {
	_ = "STUB: not implemented"
	return nil
}
