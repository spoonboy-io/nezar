package publish

import (
	"github.com/spoonboy-io/nezar/internal/transport"
	natsTransport "github.com/spoonboy-io/nezar/internal/transport/nats"
	"log"
)

const (
	NATS = iota + 1
	MQTT
)

// TODO we will need to break this out so connection independent of message sending
func Publish(publishType int) error {
	var transport transport.Transport
	var config interface{}

	switch publishType {
	case NATS:
		transport = &natsTransport.NATS{}
		config = natsTransport.Config{
			// TODO these values need to come from env vars
			Server: "localhost",
			Port:   "4222",
		}
	}

	if err := transport.OpenConn(config); err != nil {
		// TODO temporary
		log.Fatalln(err)
	}

	message, err := transport.Compose("hello world")
	if err != nil {
		// TODO temporary
		log.Fatalln(err)
	}

	subject := "test.subject"

	if err := transport.Publish(subject, message); err != nil {
		// TODO temporary
		log.Fatalln(err)
	}

	transport.CloseConn()

	return nil
}
