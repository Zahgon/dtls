// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

func findMatchingSRTPProfile(a, b []SRTPProtectionProfile) (SRTPProtectionProfile, bool) {
	_ = "STUB: not implemented"
	return *new(SRTPProtectionProfile), false
}

func findMatchingCipherSuite(a, b []CipherSuite) (CipherSuite, bool) {
	_ = "STUB: not implemented"
	return *new(CipherSuite), false
}

func splitBytes(bytes []byte, splitLen int) [][]byte { _ = "STUB: not implemented"; return nil }
