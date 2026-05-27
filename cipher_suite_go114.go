// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build go1.14

package dtls

import (
	"crypto/tls"
)

// VersionDTLS12 is the DTLS version in the same style as
// VersionTLSXX from crypto/tls.
const VersionDTLS12 = 0xfefd

// Convert from our cipherSuite interface to a tls.CipherSuite struct.
func toTLSCipherSuite(c CipherSuite) *tls.CipherSuite { _ = "STUB: not implemented"; return nil }

// CipherSuites returns a list of cipher suites currently implemented by this
// package, excluding those with security issues, which are returned by
// InsecureCipherSuites.
func CipherSuites() []*tls.CipherSuite { _ = "STUB: not implemented"; return nil }

// InsecureCipherSuites returns a list of cipher suites currently implemented by
// this package and which have security issues.
func InsecureCipherSuites() []*tls.CipherSuite { _ = "STUB: not implemented"; return nil }
