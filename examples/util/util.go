// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package util provides auxiliary utilities used in examples
package util //nolint: revive

import (
	"crypto/tls"
	"errors"
	"io"
)

const bufSize = 8192

var (
	errBlockIsNotCertificate = errors.New("block is not a certificate, unable to load certificates")
	errNoCertificateFound    = errors.New("no certificate found, unable to load certificates")
)

// Chat simulates a simple text chat session over the connection.
func Chat(conn io.ReadWriter) { _ = "STUB: not implemented"; return }

// Check is a helper to throw errors in the examples.
func Check(err error) { _ = "STUB: not implemented"; return }

//nolint:staticcheck

// LoadKeyAndCertificate reads certificates or key from file.
func LoadKeyAndCertificate(keyPath string, certificatePath string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// LoadCertificate Load/read certificate(s) from file.
func LoadCertificate(path string) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
