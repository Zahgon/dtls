// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package handshake

import (
	"time"
)

// Consts for Random in Handshake.
const (
	RandomBytesLength = 28
	RandomLength      = RandomBytesLength + 4
)

// Random value that is used in ClientHello and ServerHello
//
// https://tools.ietf.org/html/rfc4346#section-7.4.1.2
type Random struct {
	GMTUnixTime time.Time
	RandomBytes [RandomBytesLength]byte
}

// MarshalFixed encodes the Handshake.
func (r *Random) MarshalFixed() [RandomLength]byte { _ = "STUB: not implemented"; return nil }

//nolint:gosec // G115

// UnmarshalFixed populates the message from encoded data.
func (r *Random) UnmarshalFixed(data [RandomLength]byte) { _ = "STUB: not implemented"; return }

// Populate fills the handshakeRandom with random values
// may be called multiple times.
func (r *Random) Populate() error { _ = "STUB: not implemented"; return nil }
