// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

// CompressionMethodID is the ID for a CompressionMethod.
type CompressionMethodID byte

const (
	compressionMethodNull CompressionMethodID = 0
)

// CompressionMethod represents a TLS Compression Method.
type CompressionMethod struct {
	ID CompressionMethodID
}

// CompressionMethods returns all supported CompressionMethods.
func CompressionMethods() map[CompressionMethodID]*CompressionMethod {
	_ = "STUB: not implemented"
	return nil
}

// DecodeCompressionMethods the given compression methods.
func DecodeCompressionMethods(buf []byte) ([]*CompressionMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodeCompressionMethods the given compression methods.
func EncodeCompressionMethods(c []*CompressionMethod) []byte {
	_ = "STUB: not implemented"
	//nolint:gosec // G115: TLS encodes compression_methods vector length as a single byte.
	return nil
}
