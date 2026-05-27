// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build aix || darwin || dragonfly || freebsd || linux || nacl || nacljs || netbsd || openbsd || solaris || windows

// For systems having syscall.Errno.
// Update build targets by following command:
// $ grep -R ECONN $(go env GOROOT)/src/syscall/zerrors_*.go \
//     | tr "." "_" | cut -d"_" -f"2" | sort | uniq

package dtls

import (
	"os"
)

func isOpErrorTemporary(err *os.SyscallError) bool { _ = "STUB: not implemented"; return false }
