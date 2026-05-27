// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package dtls

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net"
	"time"

	"github.com/pion/dtls/v3/pkg/crypto/elliptic"
	"github.com/pion/dtls/v3/pkg/protocol"
	"github.com/pion/dtls/v3/pkg/protocol/handshake"
	"github.com/pion/logging"
)

// ServerOption configures a DTLS server.
type ServerOption interface {
	applyServer(*dtlsConfig) error
}

// ClientOption configures a DTLS client.
type ClientOption interface {
	applyClient(*dtlsConfig) error
}

// Option is an option that can be used with both client and server.
// This is used for options that apply to both sides of a connection,
// such as in the Resume function where the side is determined at runtime.
type Option interface {
	ServerOption
	ClientOption
}

// defensiveCopy copies a slice. This prevents the caller from mutating
// the config after construction. Returns empty slice if input is empty.
func defensiveCopy[T any](t ...T) []T { _ = "STUB: not implemented"; return nil }

// dtlsConfig is the internal configuration structure.
// This will eventually replace the exported Config struct.
type dtlsConfig struct { //nolint:dupl
	certificates                  []tls.Certificate
	cipherSuites                  []CipherSuiteID
	customCipherSuites            func() []CipherSuite
	signatureSchemes              []tls.SignatureScheme
	certificateSignatureSchemes   []tls.SignatureScheme
	srtpProtectionProfiles        []SRTPProtectionProfile
	srtpMasterKeyIdentifier       []byte
	clientAuth                    ClientAuthType
	extendedMasterSecret          ExtendedMasterSecretType
	flightInterval                time.Duration
	disableRetransmitBackoff      bool
	psk                           PSKCallback
	pskIdentityHint               []byte
	insecureSkipVerify            bool
	insecureHashes                bool
	verifyPeerCertificate         func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error
	verifyConnection              func(*State) error
	rootCAs                       *x509.CertPool
	clientCAs                     *x509.CertPool
	serverName                    string
	loggerFactory                 logging.LoggerFactory
	mtu                           int
	replayProtectionWindow        int
	keyLogWriter                  io.Writer
	sessionStore                  SessionStore
	supportedProtocols            []string
	ellipticCurves                []elliptic.Curve
	getCertificate                func(*ClientHelloInfo) (*tls.Certificate, error)
	getClientCertificate          func(*CertificateRequestInfo) (*tls.Certificate, error)
	insecureSkipVerifyHello       bool
	connectionIDGenerator         func() []byte
	paddingLengthGenerator        func(uint) uint
	helloRandomBytesGenerator     func() [handshake.RandomBytesLength]byte
	clientHelloMessageHook        func(handshake.MessageClientHello) handshake.Message
	serverHelloMessageHook        func(handshake.MessageServerHello) handshake.Message
	certificateRequestMessageHook func(handshake.MessageCertificateRequest) handshake.Message
	onConnectionAttempt           func(net.Addr) error
	listenConfig                  net.ListenConfig
	minVersion                    protocol.Version
	maxVersion                    protocol.Version
}

// applyDefaults applies default values to the config.
func (c *dtlsConfig) applyDefaults() { _ = "STUB: not implemented"; return }

// toConfig converts internal dtlsConfig to the exported Config struct.
// This is for backward compatibility and will be removed when Config is deprecated.
// All slice fields are copied to ensure immutability.
func (c *dtlsConfig) toConfig() *Config { _ = "STUB: not implemented"; return nil }

// buildConfig builds a Config from the provided options, for mixed client/server cases.
func buildConfig(opts ...Option) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// buildServerConfig builds a Config for server from the provided options.
func buildServerConfig(opts ...ServerOption) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildClientConfig builds a Config for client from the provided options.
func buildClientConfig(opts ...ClientOption) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sharedOption wraps an apply function that works for both client and server.
// This eliminates code duplication for options that behave identically on both sides.
type sharedOption func(*dtlsConfig) error

func (o sharedOption) applyServer(c *dtlsConfig) error { _ = "STUB: not implemented"; return nil }
func (o sharedOption) applyClient(c *dtlsConfig) error {
	_ = "STUB: not implemented"

	// WithCertificates sets the certificate chain to present to the other side of the connection.
	// For functional options, an explicitly empty slice is not allowed.
	return nil
}

func WithCertificates(certs ...tls.Certificate) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCipherSuites sets the supported cipher suites.
// For functional options, an explicitly empty slice is not allowed.
func WithCipherSuites(suites ...CipherSuiteID) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCustomCipherSuites sets the custom cipher suites provider.
// Returns an error if the provider is nil.
func WithCustomCipherSuites(fn func() []CipherSuite) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSignatureSchemes sets the signature schemes.
// For functional options, an explicitly empty slice is not allowed.
func WithSignatureSchemes(schemes ...tls.SignatureScheme) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCertificateSignatureSchemes sets the signature and hash schemes that may be used
// in digital signatures for X.509 certificates. If not set, the signature_algorithms_cert
// extension is not sent, and SignatureSchemes is used for both handshake signatures and
// certificate chain validation, as specified in RFC 8446 Section 4.2.3.
// For functional options, an explicitly empty slice is not allowed.
func WithCertificateSignatureSchemes(schemes ...tls.SignatureScheme) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSRTPProtectionProfiles sets the SRTP protection profiles.
// For functional options, an explicitly empty slice is not allowed.
func WithSRTPProtectionProfiles(profiles ...SRTPProtectionProfile) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSRTPMasterKeyIdentifier sets the SRTP master key identifier.
func WithSRTPMasterKeyIdentifier(identifier []byte) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExtendedMasterSecret sets the extended master secret policy.
// Returns an error if the type is invalid.
func WithExtendedMasterSecret(ems ExtendedMasterSecretType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFlightInterval sets the flight interval for handshake messages.
// Returns an error if the interval is not positive.
func WithFlightInterval(interval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDisableRetransmitBackoff disables retransmit backoff.
func WithDisableRetransmitBackoff(disable bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPSK sets the pre-shared key callback.
// Returns an error if the callback is nil.
func WithPSK(callback PSKCallback) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPSKIdentityHint sets the PSK identity hint.
func WithPSKIdentityHint(hint []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInsecureSkipVerify skips certificate verification.
// This should only be used for testing.
func WithInsecureSkipVerify(skip bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInsecureHashes allows the use of insecure hash algorithms.
func WithInsecureHashes(allow bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVerifyPeerCertificate sets the peer certificate verification callback.
// Returns an error if the callback is nil.
func WithVerifyPeerCertificate(fn func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithVerifyConnection sets the connection verification callback.
// Returns an error if the callback is nil.
func WithVerifyConnection(fn func(*State) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRootCAs sets the root certificate authorities.
func WithRootCAs(pool *x509.CertPool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithServerName sets the server name for certificate verification.
func WithServerName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLoggerFactory sets the logger factory for creating loggers.
func WithLoggerFactory(factory logging.LoggerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMTU sets the maximum transmission unit.
// Returns an error if the MTU is not positive.
func WithMTU(mtu int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReplayProtectionWindow sets the replay protection window size.
// Returns an error if the window size is negative.
func WithReplayProtectionWindow(window int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKeyLogWriter sets the key log writer for debugging.
// Use of KeyLogWriter compromises security and should only be used for debugging.
func WithKeyLogWriter(writer io.Writer) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSessionStore sets the session store for resumption.
func WithSessionStore(store SessionStore) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSupportedProtocols sets the supported application protocols for ALPN.
// For functional options, an explicitly empty slice is not allowed.
func WithSupportedProtocols(protocols ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEllipticCurves sets the elliptic curves.
// For functional options, an explicitly empty slice is not allowed.
func WithEllipticCurves(curves ...elliptic.Curve) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGetClientCertificate sets the client certificate getter callback.
// Returns an error if the callback is nil.
func WithGetClientCertificate(fn func(*CertificateRequestInfo) (*tls.Certificate, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithConnectionIDGenerator sets the connection ID generator.
// Returns an error if the generator is nil.
func WithConnectionIDGenerator(fn func() []byte) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPaddingLengthGenerator sets the padding length generator.
// Returns an error if the generator is nil.
func WithPaddingLengthGenerator(fn func(uint) uint) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHelloRandomBytesGenerator sets the hello random bytes generator.
// Returns an error if the generator is nil.
func WithHelloRandomBytesGenerator(fn func() [handshake.RandomBytesLength]byte) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithClientHelloMessageHook sets the client hello message hook.
// Returns an error if the hook is nil.
func WithClientHelloMessageHook(fn func(handshake.MessageClientHello) handshake.Message) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// MinVersion sets the minimum TLS version that is acceptable.
// By default, DTLS 1.2 is currently used as the minimum as it's the only supported version.
func withMinVersion(version protocol.Version) Option {
	_ = "STUB: not implemented" // nolint:unused
	return *new(Option)
}

// MaxVersion sets the maxiumum TLS version that is acceptable.
// By default, DTLS 1.2 is currently used as the minimum as it's the only supported version.
func withMaxVersion(version protocol.Version) Option {
	_ = "STUB: not implemented" // nolint:unused
	return *new(Option)
}

// serverOnlyOption wraps an apply function for server-only options.
type serverOnlyOption func(*dtlsConfig) error

func (o serverOnlyOption) applyServer(c *dtlsConfig) error {
	_ = "STUB: not implemented"

	// WithClientAuth sets the client authentication policy.
	// Returns an error if the type is invalid.
	// This option is only applicable to servers.
	return nil
}

func WithClientAuth(auth ClientAuthType) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithClientCAs sets the client certificate authorities.
// This option is only applicable to servers.
func WithClientCAs(pool *x509.CertPool) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithGetCertificate sets the certificate getter callback.
// Returns an error if the callback is nil.
// This option is only applicable to servers.
func WithGetCertificate(fn func(*ClientHelloInfo) (*tls.Certificate, error)) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithInsecureSkipVerifyHello skips hello verify phase on the server.
// This has implication on DoS attack resistance.
// This option is only applicable to servers.
func WithInsecureSkipVerifyHello(skip bool) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithServerHelloMessageHook sets the server hello message hook.
// Returns an error if the hook is nil.
// This option is only applicable to servers.
func WithServerHelloMessageHook(fn func(handshake.MessageServerHello) handshake.Message) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithCertificateRequestMessageHook sets the certificate request message hook.
// Returns an error if the hook is nil.
// This option is only applicable to servers.
func WithCertificateRequestMessageHook(fn func(handshake.MessageCertificateRequest) handshake.Message) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithOnConnectionAttempt sets the connection attempt callback.
// Returns an error if the callback is nil.
// This option is only applicable to servers.
func WithOnConnectionAttempt(fn func(net.Addr) error) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}

// WithListenConfig sets the underlying listener config.
// This option is only applicable to servers.
func WithListenConfig(listenConfig net.ListenConfig) ServerOption {
	_ = "STUB: not implemented"
	return *new(ServerOption)
}
