// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"bytes"
	"context"

	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/alert"
	"github.com/pion/dtls/v3/pkg/protocol/extension"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
	"github.com/pion/dtls/v3/pkg/protocol/recordlayer"
)

// we'll add the flight handlers for the DTLS 1.3 server here.
//
// Flight0
//
// +----------+
// | Flight 2 |
// | Flight 4 |
// | Flight 6 |
// +----------+
//
// +-----------+
// | Flight 4a |
// | Flight 6a |
// +-----------+
//
// +-----------+
// | Flight 4b |
// | Flight 6b |
// +-----------+
//
// +-----------+
// | Flight 4c |
// +-----------+

//nolint:cyclop,gocognit,gocyclo,unused
func flight13_0Parse(
	_ context.Context,
	_ flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal13, *alert.Alert, error) {
	return 0, nil, errFlightUnimplemented13
}

//nolint:unused
func flight13_2Parse(
	ctx context.Context,
	c flightConn,
	state *State,
	cache *handshakeCache,
	cfg *handshakeConfig,
) (flightVal13, *alert.Alert, error) {
	return 0, nil, errFlightUnimplemented13
}

//nolint:unused,unparam
func flight13_2Generate(
	_ flightConn,
	state *State,
	_ *handshakeCache,
	cfg *handshakeConfig,
) ([]*packet, *alert.Alert, error) {
	state.handshakeSendSequence = 0

	random := handshake.Random{}
	random.UnmarshalFixed([32]byte(handshake.HelloRetryRequestRandom()))

	exts := []extension.Extension{}

	exts = append(exts, &extension.SupportedVersions{
		Versions: supportedVersionsRange(cfg.minVersion, cfg.maxVersion),
	})

	if state.namedCurve != 0 {
		exts = append(exts, &extension.KeyShare{
			SelectedGroup: &state.namedCurve,
		})
	}

	if len(state.cookie) > 0 {
		exts = append(exts, &extension.CookieExt{
			Cookie: state.cookie,
		})
	}

	return []*packet{
		{
			record: &recordlayer.RecordLayer{
				Header: recordlayer.Header{
					Version: protocol.Version1_2,
				},
				Content: &handshake.Handshake{
					Message: &handshake.MessageServerHello{
						Version:    protocol.Version1_2,
						Random:     random,
						Extensions: exts,
					},
				},
			},
		},
	}, nil, nil
}
