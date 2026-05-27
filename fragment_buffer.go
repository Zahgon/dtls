// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

const (
	// 2 megabytes.
	fragmentBufferMaxSize  = 2000000
	fragmentBufferMaxCount = 1000
)

type fragment struct {
	recordLayerHeader recordlayer.Header
	handshakeHeader   handshake.Header
	data              []byte
}

type fragments struct {
	fragmentByOffset map[uint32]*fragment
	fragmentsLength  uint32
	handshakeLength  uint32
}

type fragmentBuffer struct {
	// map of MessageSequenceNumbers that hold slices of fragments
	cache map[uint16]*fragments

	currentMessageSequenceNumber uint16

	totalBufferSize    int
	totalFragmentCount int
}

func newFragmentBuffer() *fragmentBuffer { _ = "STUB: not implemented"; return nil }

// current total size of buffer.
func (f *fragmentBuffer) size() int { _ = "STUB: not implemented"; return 0 }

// Attempts to push a DTLS packet to the fragmentBuffer
// when it returns true it means the fragmentBuffer has inserted and the buffer shouldn't be handled
// when an error returns it is fatal, and the DTLS connection should be stopped.
func (f *fragmentBuffer) push(buf []byte) (isHandshake, isRetransmit bool, err error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return false, false, nil
}

// fragment isn't a handshake, we don't need to handle it

//nolint:gosec // G602

// Fragment is a retransmission. We have already assembled it before successfully

// Discard all headers, when rebuilding the packet we will re-build

func (f *fragmentBuffer) pop() (content []byte, epoch uint16) {
	_ = "STUB: not implemented"
	return nil, 0
}
