// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package net defines packet-oriented primitives that are compatible with net
// in the standard library.
package net

import (
	"net"
	"time"
)

// A PacketListener is the same as net.Listener but returns a net.PacketConn on
// Accept() rather than a net.Conn.
//
// Multiple goroutines may invoke methods on a PacketListener simultaneously.
type PacketListener interface {
	// Accept waits for and returns the next connection to the listener.
	Accept() (net.PacketConn, net.Addr, error)

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() net.Addr
}

// PacketListenerFromListener converts a net.Listener into a
// dtlsnet.PacketListener.
func PacketListenerFromListener(l net.Listener) PacketListener {
	_ = "STUB: not implemented"
	return *new(PacketListener)
}

// packetListenerWrapper wraps a net.Listener and implements
// dtlsnet.PacketListener.
type packetListenerWrapper struct {
	l net.Listener
}

// Accept calls Accept on the underlying net.Listener and converts the returned
// net.Conn into a net.PacketConn.
func (p *packetListenerWrapper) Accept() (net.PacketConn, net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), *new(net.Addr), nil
}

// Close closes the underlying net.Listener.
func (p *packetListenerWrapper) Close() error {
	_ = "STUB: not implemented"

	// Addr returns the address of the underlying net.Listener.
	return nil
}

func (p *packetListenerWrapper) Addr() net.Addr {
	_ = "STUB: not implemented"

	// PacketConnFromConn converts a net.Conn into a net.PacketConn.
	return *new(net.Addr)
}

func PacketConnFromConn(conn net.Conn) net.PacketConn {
	_ = "STUB: not implemented"
	return *new(net.PacketConn)
}

// packetConnWrapper wraps a net.Conn and implements net.PacketConn.
type packetConnWrapper struct {
	conn net.Conn
}

// ReadFrom reads from the underlying net.Conn and returns its remote address.
func (p *packetConnWrapper) ReadFrom(b []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

// WriteTo writes to the underlying net.Conn.
func (p *packetConnWrapper) WriteTo(b []byte, _ net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close closes the underlying net.Conn.
func (p *packetConnWrapper) Close() error { _ = "STUB: not implemented"; return nil }

// LocalAddr returns the local address of the underlying net.Conn.
func (p *packetConnWrapper) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// SetDeadline sets the deadline on the underlying net.Conn.
	new(net.Addr)
}

func (p *packetConnWrapper) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline sets the read deadline on the underlying net.Conn.
func (p *packetConnWrapper) SetReadDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// SetWriteDeadline sets the write deadline on the underlying net.Conn.
func (p *packetConnWrapper) SetWriteDeadline(t time.Time) error {
	_ = "STUB: not implemented"
	return nil
}
