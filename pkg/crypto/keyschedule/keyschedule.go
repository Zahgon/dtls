// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package keyschedule implements DTLS 1.3's key derivation related functions
package keyschedule

import (
	"errors"
	"hash"
)

var (
	errMissingHashFunction = errors.New("HKDF-Extract expected a non-nil hash function")
	errLabelTooSmall       = errors.New("HKDF-Expand-Label expected a label with length >= 7")
	errLabelTooBig         = errors.New("HKDF-Expand-Label expected a label with length <= 255")
	errContextTooBig       = errors.New("HKDF-Expand-Label expected a context with length <= 255")
	errLengthTooBig        = errors.New("HKDF-Expand-Label expected a length <= 65535")
)

const (
	DTLS13prefix = "dtls13" // RFC 9147 section 5.9
)

// HkdfExtract implements RFC 5869 section 2.2.
func HkdfExtract(hash func() hash.Hash, salt, ikm []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: Go's hkdf.Extract signature is (hash, ikm, salt),
// while RFC 5869 specifies HKDF-Extract(salt, IKM)

// HkdfExpandLabel implements RFC 8446 section 7.1 with RFC 9147 section 5.9's defined DTLS prefix.
func HkdfExpandLabel(hash func() hash.Hash, secret []byte, label string, context []byte, length int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RFC 8446 section 7.1
// opaque label<7..255>

// RFC 5869 section 2.3
// L        length of output keying material in octets
//          (<= 255*HashLen)
// https://datatracker.ietf.org/doc/html/rfc5869#section-2.3

//nolint:gosec

// DeriveSecret implements RFC 8446 section 7.1.
//
// TranscriptHash is defined in RFC 8446 section 4.4.
func DeriveSecret(hash func() hash.Hash, secret []byte, label string, transcriptHash hash.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
