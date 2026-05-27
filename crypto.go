// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"crypto"
	"crypto/x509"
	"math/big"

	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/crypto/hash"
	"github.com/pion/dtls/v3/pkg/crypto/signature"
	"github.com/pion/dtls/v3/pkg/crypto/signaturehash"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func valueKeyMessage(clientRandom, serverRandom, publicKey []byte, namedCurve elliptic.Curve) []byte {
	_ = "STUB: not implemented"
	return nil
}

// named curve

//nolint:gosec // G115, no risk of overflow, the biggest supported curve is 97 bytes.

// validateSignatureAlgOID validates that the signature scheme matches the
// certificate's public key algorithm OID. This is required by RFC 8446 Section 4.2.3:
// - RSA_PSS_RSAE requires rsaEncryption OID
// - RSA_PSS_PSS requires id-RSASSA-PSS OID
//
// Note: returns nil if the given signature.Algorithm is not PSS based.
//
// https://www.rfc-editor.org/rfc/rfc8446#section-4.2.3
func validateSignatureAlgOID(cert *x509.Certificate, sigAlg signature.Algorithm) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the certificate's public key algorithm OID from the raw certificate
// We need to parse the SubjectPublicKeyInfo to get the algorithm OID

// Check RSAE variants (0x0804-0x0806) require rsaEncryption OID

// OID: rsaEncryption

// Check PSS variants (0x0809-0x080b) require id-RSASSA-PSS OID

// OID: id-RSASSA-PSS

// If the client provided a "signature_algorithms" extension, then all
// certificates provided by the server MUST be signed by a
// hash/signature algorithm pair that appears in that extension
//
// https://tools.ietf.org/html/rfc5246#section-7.4.2
func generateKeySignature(
	clientRandom, serverRandom, publicKey []byte,
	namedCurve elliptic.Curve,
	signer crypto.Signer,
	hashAlgorithm hash.Algorithm,
	signatureAlgorithm signature.Algorithm,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://crypto.stackexchange.com/a/55483

// Use RSA-PSS if the signature algorithm is PSS

// Otherwise use PKCS#1 v1.5

//nolint:dupl,cyclop
func verifyKeySignature(
	message, remoteKeySignature []byte,
	hashAlgorithm hash.Algorithm,
	signatureAlgorithm signature.Algorithm,
	rawCertificates [][]byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate that the signature algorithm matches the certificate's OID

// Use RSA-PSS verification if the signature algorithm is PSS

// Otherwise use PKCS#1 v1.5

// If the server has sent a CertificateRequest message, the client MUST send the Certificate
// message.  The ClientKeyExchange message is now sent, and the content
// of that message will depend on the public key algorithm selected
// between the ClientHello and the ServerHello.  If the client has sent
// a certificate with signing ability, a digitally-signed
// CertificateVerify message is sent to explicitly verify possession of
// the private key in the certificate.
// https://tools.ietf.org/html/rfc5246#section-7.3
func generateCertificateVerify(
	handshakeBodies []byte,
	signer crypto.Signer,
	hashAlgorithm hash.Algorithm,
	signatureAlgorithm signature.Algorithm,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://pkg.go.dev/crypto/ed25519#PrivateKey.Sign
// Sign signs the given message with priv. Ed25519 performs two passes over
// messages to be signed and therefore cannot handle pre-hashed messages.

// Use RSA-PSS if the signature algorithm is PSS

// Otherwise use PKCS#1 v1.5

//nolint:dupl,cyclop
func verifyCertificateVerify(
	handshakeBodies []byte,
	hashAlgorithm hash.Algorithm,
	signatureAlgorithm signature.Algorithm,
	remoteKeySignature []byte,
	rawCertificates [][]byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate that the signature algorithm matches the certificate's OID

// Use RSA-PSS verification if the signature algorithm is PSS

// Otherwise use PKCS#1 v1.5

func loadCerts(rawCertificates [][]byte) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func verifyClientCert(
	rawCertificates [][]byte,
	roots *x509.CertPool,
	certSignatureSchemes []signaturehash.Algorithm,
) (chains [][]*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate certificate signature algorithms if specified.
// At least one chain must use only allowed signature algorithms.

func verifyServerCert(
	rawCertificates [][]byte,
	roots *x509.CertPool,
	serverName string,
	certSignatureSchemes []signaturehash.Algorithm,
) (chains [][]*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate certificate signature algorithms if specified.
// At least one chain must use only allowed signature algorithms.

// validateCertificateSignatureAlgorithms validates that all certificates in the chain
// use signature algorithms that are in the allowed list. This implements the
// signature_algorithms_cert extension validation per RFC 8446 Section 4.2.3.
func validateCertificateSignatureAlgorithms(
	certs []*x509.Certificate,
	allowedAlgorithms []signaturehash.Algorithm,
) error {
	_ = "STUB: not implemented"
	return nil
}

// No restrictions specified

// Validate each certificate's signature algorithm (except the root, which we trust)

// Check if this algorithm is in the allowed list
