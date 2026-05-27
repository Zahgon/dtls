// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"net"
)

// Resume imports an already established dtls connection using a specific dtls state.
//
// Deprecated: Use ResumeWithOptions instead.
func Resume(state *State, conn net.PacketConn, rAddr net.Addr, config *Config) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResumeWithOptions imports an already established dtls connection using a specific dtls state.
func ResumeWithOptions(state *State, conn net.PacketConn, rAddr net.Addr, opts ...Option) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
