// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package closer provides signaling channel for shutdown
package closer

import (
	"context"
)

// Closer allows for each signaling a channel for shutdown.
type Closer struct {
	ctx       context.Context //nolint:containedctx
	closeFunc func()
}

// NewCloser creates a new instance of Closer.
func NewCloser() *Closer { _ = "STUB: not implemented"; return nil }

// NewCloserWithParent creates a new instance of Closer with a parent context.
func NewCloserWithParent(ctx context.Context) *Closer { _ = "STUB: not implemented"; return nil }

// Done returns a channel signaling when it is done.
func (c *Closer) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Err returns an error of the context.
func (c *Closer) Err() error {
	_ = "STUB: not implemented"

	// Close sends a signal to trigger the ctx done channel.
	return nil
}

func (c *Closer) Close() { _ = "STUB: not implemented"; return }
