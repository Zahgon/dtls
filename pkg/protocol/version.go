// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package protocol provides the DTLS wire format
package protocol

// Version enums.
var (
	Version1_0 = Version{Major: 0xfe, Minor: 0xff} //nolint:gochecknoglobals
	Version1_2 = Version{Major: 0xfe, Minor: 0xfd} //nolint:gochecknoglobals
	Version1_3 = Version{Major: 0xfe, Minor: 0xfc} //nolint:gochecknoglobals
)

// Version is the minor/major value in the RecordLayer
// and ClientHello/ServerHello
//
// https://tools.ietf.org/html/rfc4346#section-6.2.1
type Version struct {
	Major, Minor uint8
}

// Equal determines if two protocol versions are equal.
func (v Version) Equal(x Version) bool { _ = "STUB: not implemented"; return false }

// IsSupportedBytes returns true if it's supported by Pion. Only DTLS 1.2 is currently supported.
// DTLS 1.3 is a work in progress and is currently being implemented.
func IsSupportedBytes(major uint8, minor uint8) bool { _ = "STUB: not implemented"; return false }

// IsSupportedVersion returns true if it's supported by Pion. Only DTLS 1.2 is currently supported.
// DTLS 1.3 is a work in progress and is currently being implemented.
func IsSupportedVersion(v Version) bool { _ = "STUB: not implemented"; return false }

// IsValidBytes returns true if the bytes represent a valid DTLS version as defined in RFC9147 below.
// Note that this is not the same as whether it's *supported* by Pion. Please see IsSupportedBytes() for more info.
//
// https://tools.ietf.org/html/rfc9147#section-5.3 (see legacy_version)
func IsValidBytes(major uint8, minor uint8) bool { _ = "STUB: not implemented"; return false }

// IsValidVersion returns true if the bytes represent a valid DTLS version as defined in RFC9147 below.
// Note that this is not the same as whether it's *supported* by Pion. Please see IsSupportedBytes() for more info.
// /
// https://tools.ietf.org/html/rfc9147#section-5.3 (see legacy_version)
func IsValidVersion(v Version) bool { _ = "STUB: not implemented"; return false }
