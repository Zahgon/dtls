// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package net implements DTLS specific networking primitives.
// NOTE: this package is an adaption of pion/transport/packetio that allows for
// storing a remote address alongside each packet in the buffer and implements
// relevant methods of net.PacketConn. If possible, the updates made in this
// repository will be reflected back upstream. If not, it is likely that this
// will be moved to a public package in this repository.
//
// This package was migrated from pion/transport/packetio at
// https://github.com/pion/transport/commit/6890c795c807a617c054149eee40a69d7fdfbfdb
package net

import (
	"bytes"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/pion/transport/v4/deadline"
)

// ErrTimeout indicates that deadline was reached before operation could be
// completed.
var ErrTimeout = errors.New("buffer: i/o timeout")

// AddrPacket is a packet payload and the associated remote address from which
// it was received.
type AddrPacket struct {
	addr net.Addr
	data bytes.Buffer
}

// PacketBuffer is a circular buffer for network packets. Each slot in the
// buffer contains the remote address from which the packet was received, as
// well as the packet data.
type PacketBuffer struct {
	mutex sync.Mutex

	packets     []AddrPacket
	write, read int

	// full indicates whether the buffer is full, which is needed to distinguish
	// when the write pointer and read pointer are at the same index.
	full bool

	notify chan struct{}
	closed bool

	readDeadline *deadline.Deadline
}

// NewPacketBuffer creates a new PacketBuffer.
func NewPacketBuffer() *PacketBuffer { _ = "STUB: not implemented"; return nil }

// In the narrow context in which this package is currently used, there
// will always be at least one packet written to the buffer. Therefore,
// we opt to allocate with size of 1 during construction, rather than
// waiting until that first packet is written.

// WriteTo writes a single packet to the buffer. The supplied address will
// remain associated with the packet.
func (b *PacketBuffer) WriteTo(pkt []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check to see if we are full.

// If so, grow AddrPacket buffer.

// Double the number of packets.

// Increase the number of packets by 25%.

// Update read/write pointers and mark buffer as not full.

// Store the packet at the write pointer.

// Increment write pointer.

// If the write pointer is equal to the length of the buffer, wrap around.

// If a write resulted in making write and read pointers equivalent, then we
// are full.

// ReadFrom reads a single packet from the buffer, or blocks until one is
// available.
func (b *PacketBuffer) ReadFrom(packet []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return 0, *new(net.Addr), nil
}

// Copy packet data from buffer.

// Advance read pointer.

// If we were full before reading and have successfully read, we are
// no longer full.

// Close closes the buffer, allowing unread packets to be read, but erroring on
// any new writes.
func (b *PacketBuffer) Close() (err error) { _ = "STUB: not implemented"; return nil }

// SetReadDeadline sets the read deadline for the buffer.
func (b *PacketBuffer) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
