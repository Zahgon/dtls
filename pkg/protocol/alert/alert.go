// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package alert implements TLS alert protocol https://tools.ietf.org/html/rfc5246#section-7.2
package alert

import (
	"errors"

	"github.com/pion/dtls/v3/pkg/protocol"
)

var errBufferTooSmall = &protocol.TemporaryError{Err: errors.New("buffer is too small")} //nolint:err113

// Level is the level of the TLS Alert.
type Level byte

// Level enums.
const (
	Warning Level = 1
	Fatal   Level = 2
)

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

// Description is the extended info of the TLS Alert.
type Description byte

// Description enums.
const (
	CloseNotify            Description = 0
	UnexpectedMessage      Description = 10
	BadRecordMac           Description = 20
	DecryptionFailed       Description = 21
	RecordOverflow         Description = 22
	DecompressionFailure   Description = 30
	HandshakeFailure       Description = 40
	NoCertificate          Description = 41
	BadCertificate         Description = 42
	UnsupportedCertificate Description = 43
	CertificateRevoked     Description = 44
	CertificateExpired     Description = 45
	CertificateUnknown     Description = 46
	IllegalParameter       Description = 47
	UnknownCA              Description = 48
	AccessDenied           Description = 49
	DecodeError            Description = 50
	DecryptError           Description = 51
	ExportRestriction      Description = 60
	ProtocolVersion        Description = 70
	InsufficientSecurity   Description = 71
	InternalError          Description = 80
	UserCanceled           Description = 90
	NoRenegotiation        Description = 100
	UnsupportedExtension   Description = 110
	NoApplicationProtocol  Description = 120
)

func (d Description) String() string {
	_ = "STUB: not implemented" //nolint:cyclop
	return ""
}

// Alert is one of the content types supported by the TLS record layer.
// Alert messages convey the severity of the message
// (warning or fatal) and a description of the alert.  Alert messages
// with a level of fatal result in the immediate termination of the
// connection.  In this case, other connections corresponding to the
// session may continue, but the session identifier MUST be invalidated,
// preventing the failed session from being used to establish new
// connections.  Like other messages, alert messages are encrypted and
// compressed, as specified by the current connection state.
// https://tools.ietf.org/html/rfc5246#section-7.2
type Alert struct {
	Level       Level
	Description Description
}

// ContentType returns the ContentType of this Content.
func (a Alert) ContentType() protocol.ContentType {
	_ = "STUB: not implemented"
	return *new(protocol.ContentType)
}

// Marshal returns the encoded alert.
func (a *Alert) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal populates the alert from binary data.
func (a *Alert) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (a *Alert) String() string { _ = "STUB: not implemented"; return "" }
