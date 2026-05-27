// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package prf implements TLS 1.2 Pseudorandom functions
package prf

import (
	"errors"
	"hash"

	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/protocol"
)

const (
	masterSecretLabel         = "master secret"
	extendedMasterSecretLabel = "extended master secret"
	keyExpansionLabel         = "key expansion"
	verifyDataClientLabel     = "client finished"
	verifyDataServerLabel     = "server finished"
)

// HashFunc allows callers to decide what hash is used in PRF.
type HashFunc func() hash.Hash

// EncryptionKeys is all the state needed for a TLS CipherSuite.
type EncryptionKeys struct {
	MasterSecret   []byte
	ClientMACKey   []byte
	ServerMACKey   []byte
	ClientWriteKey []byte
	ServerWriteKey []byte
	ClientWriteIV  []byte
	ServerWriteIV  []byte
}

var errInvalidNamedCurve = &protocol.FatalError{Err: errors.New("invalid named curve")} //nolint:err113

func (e *EncryptionKeys) String() string { _ = "STUB: not implemented"; return "" }

// PSKPreMasterSecret generates the PSK Premaster Secret
// The premaster secret is formed as follows: if the PSK is N octets
// long, concatenate a uint16 with the value N, N zero octets, a second
// uint16 with the value N, and the PSK itself.
//
// https://tools.ietf.org/html/rfc4279#section-2
func PSKPreMasterSecret(psk []byte) []byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G115

// EcdhePSKPreMasterSecret implements TLS 1.2 Premaster Secret generation given a psk, a keypair and a curve
//
// https://datatracker.ietf.org/doc/html/rfc5489#section-2
func EcdhePSKPreMasterSecret(psk, publicKey, privateKey []byte, curve elliptic.Curve) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// write preMasterSecret length

//nolint:gosec // G115

// write preMasterSecret

// write psk length
//nolint:gosec // G115

// write psk

// PreMasterSecret implements TLS 1.2 Premaster Secret generation given a keypair and a curve.
func PreMasterSecret(publicKey, privateKey []byte, curve elliptic.Curve) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NIST: SEC1 uncompressed; X25519: 32-byte u

// PHash is PRF is the SHA-256 hash function is used for all cipher suites
// defined in this TLS 1.2 document and in TLS documents published prior to this
// document when TLS 1.2 is negotiated.  New cipher suites MUST explicitly
// specify a PRF and, in general, SHOULD use the TLS PRF with SHA-256 or a
// stronger standard hash function.
//
//	P_hash(secret, seed) = HMAC_hash(secret, A(1) + seed) +
//	                       HMAC_hash(secret, A(2) + seed) +
//	                       HMAC_hash(secret, A(3) + seed) + ...
//
// A() is defined as:
//
//	A(0) = seed
//	A(i) = HMAC_hash(secret, A(i-1))
//
// P_hash can be iterated as many times as necessary to produce the
// required quantity of data.  For example, if P_SHA256 is being used to
// create 80 bytes of data, it will have to be iterated three times
// (through A(3)), creating 96 bytes of output data; the last 16 bytes
// of the final iteration will then be discarded, leaving 80 bytes of
// output data.
//
// https://tools.ietf.org/html/rfc4346w
func PHash(secret, seed []byte, requestedLength int, hashFunc HashFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtendedMasterSecret generates a Extended MasterSecret as defined in
// https://tools.ietf.org/html/rfc7627
func ExtendedMasterSecret(preMasterSecret, sessionHash []byte, h HashFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MasterSecret generates a TLS 1.2 MasterSecret.
func MasterSecret(preMasterSecret, clientRandom, serverRandom []byte, h HashFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateEncryptionKeys is the final step TLS 1.2 PRF. Given all state generated so far generates
// the final keys need for encryption.
func GenerateEncryptionKeys(
	masterSecret, clientRandom, serverRandom []byte,
	macLen, keyLen, ivLen int,
	h HashFunc,
) (*EncryptionKeys, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prfVerifyData(masterSecret, handshakeBodies []byte, label string, hashFunc HashFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyDataClient is caled on the Client Side to either verify or generate the VerifyData message.
func VerifyDataClient(masterSecret, handshakeBodies []byte, h HashFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyDataServer is caled on the Server Side to either verify or generate the VerifyData message.
func VerifyDataServer(masterSecret, handshakeBodies []byte, h HashFunc) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
