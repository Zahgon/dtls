// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package util

import (
	"net"
	"sync"
)

// Hub is a helper to handle one to many chat.
type Hub struct {
	conns map[string]net.Conn
	lock  sync.RWMutex
}

// NewHub builds a new hub.
func NewHub() *Hub { _ = "STUB: not implemented"; return nil }

// Register adds a new conn to the Hub.
func (h *Hub) Register(conn net.Conn) { _ = "STUB: not implemented"; return }

func (h *Hub) readLoop(conn net.Conn) { _ = "STUB: not implemented"; return }

func (h *Hub) unregister(conn net.Conn) { _ = "STUB: not implemented"; return }

func (h *Hub) broadcast(msg []byte) { _ = "STUB: not implemented"; return }

// Chat starts the stdin readloop to dispatch messages to the hub.
func (h *Hub) Chat() { _ = "STUB: not implemented"; return }
