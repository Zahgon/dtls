// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"context"

	"github.com/pion/dtls/v3/pkg/protocol/alert"
)

// Parse received handshakes and return next flightVal.
type flightParser func(
	context.Context,
	flightConn,
	*State,
	*handshakeCache,
	*handshakeConfig,
) (flightVal, *alert.Alert, error)

// Generate flights.
type flightGenerator func(flightConn, *State, *handshakeCache, *handshakeConfig) ([]*packet, *alert.Alert, error)

func (f flightVal) getFlightParser() (flightParser, error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return *new(flightParser), nil
}

func (f flightVal) getFlightGenerator() (gen flightGenerator, retransmit bool, err error) {
	_ = "STUB: not implemented" //nolint:cyclop
	return *new(flightGenerator), false, nil
}

// https://tools.ietf.org/html/rfc6347#section-3.2.1
// HelloVerifyRequests must not be retransmitted.
