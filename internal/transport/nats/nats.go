package natsTransport

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spoonboy-io/nezar/internal/transport"
)

type NATS struct {
	conn *nats.Conn
	conf Config
}

type Config struct {
	Server string
	Port   string
	Auth   string //TODO unlikely to be a string, but no auth for now
}

func (n *NATS) OpenConn(config any) error {
	var conf Config
	var ok bool

	if conf, ok = config.(Config); !ok {
		return transport.ERR_BAD_TYPE
	}

	if conf.Server == "" || conf.Port == "" {
		return transport.ERR_BAD_CONFIG
	}

	// store config just in case
	n.conf = conf

	natsUrl := fmt.Sprintf("nats://%s:%s", n.conf.Server, n.conf.Port)
	conn, err := nats.Connect(natsUrl)
	if err != nil {
		return err
	}

	// store the connection value
	n.conn = conn

	return nil
}

func (n *NATS) Compose(raw any) ([]byte, error) {
	// TODO do we need an interface for raw??
	return []byte{}, nil
}

func (n *NATS) Publish(target string, message []byte) error {
	err := n.conn.Publish(target, message)
	if err != nil {
		return err
	}
	return nil
}

func (n *NATS) CloseConn() error {
	n.conn.Close()
	// error not implemented in NATS conn close
	return nil
}
