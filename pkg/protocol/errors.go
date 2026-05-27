// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

import (
	"errors"
)

var (
	errBufferTooSmall    = &TemporaryError{Err: errors.New("buffer is too small")} //nolint:err113
	errInvalidCipherSpec = &FatalError{Err: errors.New("cipher spec invalid")}     //nolint:err113
)

// FatalError indicates that the DTLS connection is no longer available.
// It is mainly caused by wrong configuration of server or client.
type FatalError struct {
	Err error
}

// InternalError indicates and internal error caused by the implementation,
// and the DTLS connection is no longer available.
// It is mainly caused by bugs or tried to use unimplemented features.
type InternalError struct {
	Err error
}

// TemporaryError indicates that the DTLS connection is still available, but the request was failed temporary.
type TemporaryError struct {
	Err error
}

// TimeoutError indicates that the request was timed out.
type TimeoutError struct {
	Err error
}

// HandshakeError indicates that the handshake failed.
type HandshakeError struct {
	Err error
}

// Timeout implements net.Error.Timeout().
func (*FatalError) Timeout() bool {
	_ = "STUB: not implemented"

	// Temporary implements net.Error.Temporary().
	return false
}

func (*FatalError) Temporary() bool {
	_ = "STUB: not implemented"

	// Unwrap implements Go1.13 error unwrapper.
	return false
}

func (e *FatalError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *FatalError) Error() string { _ = "STUB: not implemented"; return "" }

// Timeout implements net.Error.Timeout().
func (*InternalError) Timeout() bool {
	_ = "STUB: not implemented"

	// Temporary implements net.Error.Temporary().
	return false
}

func (*InternalError) Temporary() bool {
	_ = "STUB: not implemented"

	// Unwrap implements Go1.13 error unwrapper.
	return false
}

func (e *InternalError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *InternalError) Error() string { _ = "STUB: not implemented"; return "" }

// Timeout implements net.Error.Timeout().
func (*TemporaryError) Timeout() bool {
	_ = "STUB: not implemented"

	// Temporary implements net.Error.Temporary().
	return false
}

func (*TemporaryError) Temporary() bool {
	_ = "STUB: not implemented"

	// Unwrap implements Go1.13 error unwrapper.
	return false
}

func (e *TemporaryError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *TemporaryError) Error() string { _ = "STUB: not implemented"; return "" }

// Timeout implements net.Error.Timeout().
func (*TimeoutError) Timeout() bool {
	_ = "STUB: not implemented"

	// Temporary implements net.Error.Temporary().
	return false
}

func (*TimeoutError) Temporary() bool {
	_ = "STUB: not implemented"

	// Unwrap implements Go1.13 error unwrapper.
	return false
}

func (e *TimeoutError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *TimeoutError) Error() string { _ = "STUB: not implemented"; return "" }

// Timeout implements net.Error.Timeout().
func (e *HandshakeError) Timeout() bool { _ = "STUB: not implemented"; return false }

// Temporary implements net.Error.Temporary().
func (e *HandshakeError) Temporary() bool { _ = "STUB: not implemented"; return false }

//nolint

// Unwrap implements Go1.13 error unwrapper.
func (e *HandshakeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *HandshakeError) Error() string { _ = "STUB: not implemented"; return "" }
