// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

// RandomCIDGenerator is a random Connection ID generator where CID is the
// specified size. Specifying a size of 0 will indicate to peers that sending a
// Connection ID is not necessary.
func RandomCIDGenerator(size int) func() []byte { _ = "STUB: not implemented"; return nil }

//nolint -- nonrecoverable

// OnlySendCIDGenerator enables sending Connection IDs negotiated with a peer,
// but indicates to the peer that sending Connection IDs in return is not
// necessary.
func OnlySendCIDGenerator() func() []byte { _ = "STUB: not implemented"; return nil }

// cidDatagramRouter extracts connection IDs from incoming datagram payloads and
// uses them to route to the proper connection.
// NOTE: properly routing datagrams based on connection IDs requires using
// constant size connection IDs.
func cidDatagramRouter(size int) func([]byte) (string, bool) { _ = "STUB: not implemented"; return nil }

// cidConnIdentifier extracts connection IDs from outgoing ServerHello records
// and associates them with the associated connection.
// NOTE: a ServerHello should always be the first record in a datagram if
// multiple are present, so we avoid iterating through all packets if the first
// is not a ServerHello.
func cidConnIdentifier() func([]byte) (string, bool) {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}
