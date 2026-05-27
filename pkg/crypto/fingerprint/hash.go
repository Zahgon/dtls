// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package fingerprint

import (
	"crypto"
	"errors"
)

var errInvalidHashAlgorithm = errors.New("fingerprint: invalid hash algorithm")

func nameToHash() map[string]crypto.Hash { _ = "STUB: not implemented"; return nil }

// [RFC3279]
// [RFC3279]
// [RFC4055]
// [RFC4055]
// [RFC4055]
// [RFC4055]

// HashFromString allows looking up a hash algorithm by it's string representation.
func HashFromString(s string) (crypto.Hash, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), nil
}

// StringFromHash allows looking up a string representation of the crypto.Hash.
func StringFromHash(hash crypto.Hash) (string, error) { _ = "STUB: not implemented"; return "", nil }
