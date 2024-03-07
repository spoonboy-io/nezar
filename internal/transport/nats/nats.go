package natsTransport

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spoonboy-io/nezar/internal/transport"
)

type NATS struct {
	Conn *nats.Conn
	Conf Config
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
	n.Conf = conf

	natsUrl := fmt.Sprintf("nats://%s:%s", n.Conf.Server, n.Conf.Port)
	conn, err := nats.Connect(natsUrl)
	if err != nil {
		return err
	}

	// store the connection value
	n.Conn = conn

	return nil
}

func (n *NATS) Compose(raw any) ([]byte, error) {
	// TODO do we need an interface for raw??
	var input string
	var ok bool
	if input, ok = raw.(string); !ok {
		return nil, transport.ERR_BAD_COMPOSE
	}
	return []byte(input), nil
}

func (n *NATS) Publish(target string, message []byte) error {
	err := n.Conn.Publish(target, message)
	if err != nil {
		return err
	}
	return nil
}

func (n *NATS) CloseConn() error {
	n.Conn.Close()
	// error not implemented in NATS conn close
	return nil
}
