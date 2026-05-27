// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package ccm implements a CCM, Counter with CBC-MAC
// as per RFC 3610.
//
// See https://tools.ietf.org/html/rfc3610
//
// This code was lifted from https://github.com/bocajim/dtls/blob/a3300364a283fcb490d28a93d7fcfa7ba437fbbe/ccm/ccm.go
// and as such was not written by the Pions authors. Like Pions this
// code is licensed under MIT.
//
// A request for including CCM into the Go standard library
// can be found as issue #27484 on the https://github.com/golang/go/
// repository.
package ccm

import (
	"crypto/cipher"
	"errors"
)

// ccm represents a Counter with CBC-MAC with a specific key.
type ccm struct {
	b cipher.Block
	M uint8
	L uint8
}

const ccmBlockSize = 16

// CCM is a block cipher in Counter with CBC-MAC mode.
// Providing authenticated encryption with associated data via the cipher.AEAD interface.
type CCM interface {
	cipher.AEAD
	// MaxLength returns the maxium length of plaintext in calls to Seal.
	// The maximum length of ciphertext in calls to Open is MaxLength()+Overhead().
	// The maximum length is related to CCM's `L` parameter (15-noncesize) and
	// is 1<<(8*L) - 1 (but also limited by the maxium size of an int).
	MaxLength() int
}

var (
	errInvalidBlockSize = errors.New("ccm: NewCCM requires 128-bit block cipher")
	errInvalidTagSize   = errors.New("ccm: tagsize must be 4, 6, 8, 10, 12, 14, or 16")
	errInvalidNonceSize = errors.New("ccm: invalid nonce size")
)

// NewCCM returns the given 128-bit block cipher wrapped in CCM.
// The tagsize must be an even integer between 4 and 16 inclusive
// and is used as CCM's `M` parameter.
// The noncesize must be an integer between 7 and 13 inclusive,
// 15-noncesize is used as CCM's `L` parameter.
func NewCCM(b cipher.Block, tagsize, noncesize int) (CCM, error) {
	_ = "STUB: not implemented"
	return *new(CCM), nil
}

//nolint:gosec // G114

func (c *ccm) NonceSize() int { _ = "STUB: not implemented"; return 0 }
func (c *ccm) Overhead() int  { _ = "STUB: not implemented"; return 0 }
func (c *ccm) MaxLength() int { _ = "STUB: not implemented"; return 0 }

func maxlen(l uint8, tagsize int) int { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // G114
// The maximum lentgh on a 64bit arch

//nolint:gosec // G114
// We have only 32bit int's

//nolint:gosec // G114

// MaxNonceLength returns the maximum nonce length for a given plaintext length.
// A return value <= 0 indicates that plaintext length is too large for
// any nonce length.
func MaxNonceLength(pdatalen int) int { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // G115

func (c *ccm) cbcRound(mac, data []byte) { _ = "STUB: not implemented"; return }

func (c *ccm) cbcData(mac, data []byte) { _ = "STUB: not implemented"; return }

var errPlaintextTooLong = errors.New("ccm: plaintext too large")

func (c *ccm) tag(nonce, plaintext, adata []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif
// First adata block includes adata length

//nolint:gosec // G115

// sliceForAppend takes a slice and a requested number of bytes. It returns a
// slice with the contents of the given slice followed by that many bytes and a
// second slice that aliases into it and contains only the extra bytes. If the
// original slice has sufficient capacity then no allocation is performed.
// From crypto/cipher/gcm.go
// .
func sliceForAppend(in []byte, n int) (head, tail []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Seal encrypts and authenticates plaintext, authenticates the
// additional data and appends the result to dst, returning the updated
// slice. The nonce must be NonceSize() bytes long and unique for all
// time, for a given key.
// The plaintext must be no longer than MaxLength() bytes long.
//
// The plaintext and dst may alias exactly or not at all.
func (c *ccm) Seal(dst, nonce, plaintext, adata []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// The cipher.AEAD interface doesn't allow for an error return.
// nolint

var (
	errOpen               = errors.New("ccm: message authentication failed")
	errCiphertextTooShort = errors.New("ccm: ciphertext too short")
	errCiphertextTooLong  = errors.New("ccm: ciphertext too long")
)

func (c *ccm) Open(dst, nonce, ciphertext, adata []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cannot decrypt directly to dst since we're not supposed to
// reveal the plaintext to the caller if authentication fails.
