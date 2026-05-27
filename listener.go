// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"net"

	dtlsnet "github.com/pion/dtls/v3/pkg/net"
)

// Listen creates a DTLS listener.
//
// Deprecated: Use ListenWithOptions instead.
func Listen(network string, laddr *net.UDPAddr, config *Config) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// If connection ID support is enabled, then they must be supported in
// routing.

// ListenWithOptions creates a DTLS listener.
func ListenWithOptions(network string, laddr *net.UDPAddr, opts ...ServerOption) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// NewListener creates a DTLS listener which accepts connections from an inner Listener.
//
// Deprecated: Use NewListenerWithOptions instead.
func NewListener(inner dtlsnet.PacketListener, config *Config) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// NewListenerWithOptions creates a DTLS listener which accepts connections from an inner Listener.
func NewListenerWithOptions(inner dtlsnet.PacketListener, opts ...ServerOption) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// listener represents a DTLS listener.
type listener struct {
	config *Config
	parent dtlsnet.PacketListener
}

// Accept waits for and returns the next connection to the listener.
// You have to either close or read on all connection that are created.
func (l *listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
// Already Accepted connections are not closed.
func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

// Addr returns the listener's network address.
func (l *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
