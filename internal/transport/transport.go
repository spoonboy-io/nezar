package transport

import "errors"

var (
	// helpers to be used by transport packages
	ERR_BAD_TYPE    = errors.New("config is not of the correct type")
	ERR_BAD_CONFIG  = errors.New("required config is missing, please check")
	ERR_BAD_COMPOSE = errors.New("raw message is not of expected type")
)

// Transport is an interface to be implemented by any transport
// package that will send messages using its messaging protocol
type Transport interface {
	// OpenConn should pass a config so that the connection to
	// transport and can be opened
	OpenConn(config any) error

	// Compose should transform a message payload into a format
	// acceptable to the transport
	Compose(raw any) ([]byte, error)

	// Publish sends the message, target is the topic/channel
	// or subject used by the transport
	Publish(target string, message []byte) error

	// CloseConn is a helper for closing connection and clean up
	CloseConn() error
}
