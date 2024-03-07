package test

import (
	"fmt"
	"github.com/nats-io/nats-server/v2/server"
	natsServer "github.com/nats-io/nats-server/v2/test"
	natsTransport "github.com/spoonboy-io/nezar/internal/transport/nats"
	"strconv"
	"testing"
)

func StartEchoServer(conf natsTransport.Config) *server.Server {
	fmt.Println("starting test NATS server..")
	opts := natsServer.DefaultTestOptions
	opts.Port, _ = strconv.Atoi(conf.Port)
	return natsServer.RunServer(&opts)
}

func TestNATSTransport(t *testing.T) {
	conf := natsTransport.Config{
		Server: "localhost",
		Port:   "4222",
	}

	// we'll user a NATS server to test against, no need to mock
	s := StartEchoServer(conf)
	defer func() {
		fmt.Println("shutting down test NATS server..")
		s.Shutdown()
	}()

	// tests transport
	client := natsTransport.NATS{}
	if err := client.OpenConn(conf); err != nil {
		t.Fatal(err)
	}

	message, err := client.Compose("hello world")
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Publish("example.server", message); err != nil {
		t.Fatal(err)
	}

	_ := client.CloseConn()
}
