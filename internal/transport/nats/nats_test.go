package natsTransport_test

import (
	"fmt"
	"github.com/nats-io/nats-server/v2/server"
	natsServer "github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
	natsTransport "github.com/spoonboy-io/nezar/internal/transport/nats"
	"strconv"
	"testing"
	"time"
)

func StartEchoServer(conf natsTransport.Config) *server.Server {
	fmt.Println("starting test NATS server..")
	opts := natsServer.DefaultTestOptions
	opts.Port, _ = strconv.Atoi(conf.Port)
	return natsServer.RunServer(&opts)
}

func TestNATSTransport(t *testing.T) {
	conf := natsTransport.Config{
		Server: "127.0.0.1",
		Port:   "4222",
	}

	subject := "example.server"
	messageText := "hello world"

	// we'll user a NATS server to test against, no need to mock
	s := StartEchoServer(conf)

	// tests
	client := natsTransport.NATS{}

	// testing OpenConn
	if err := client.OpenConn(conf); err != nil {
		t.Fatal(err)
	}

	var gotSubject, gotMessage string

	// set up subscription to receive echo
	_, err := client.Conn.Subscribe(subject, func(msg *nats.Msg) {
		// Echo back the received message
		gotSubject = msg.Subject
		gotMessage = string(msg.Data)
	})

	if err != nil {
		t.Fatal(err)
	}

	// testing Compose
	message, err := client.Compose(messageText)
	if err != nil {
		t.Fatal(err)
	}

	// testing Publish
	if err := client.Publish(subject, message); err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	// testing CloseConn
	_ = client.CloseConn()

	// test assertions
	if gotSubject != subject {
		t.Errorf("unexpected subject, want '%s', got '%s'", subject, gotSubject)
	}
	if gotSubject != subject {
		t.Errorf("unexpected message, want '%s', got '%s'", messageText, gotMessage)
	}

	// cleanup
	fmt.Println("shutting down test NATS server..")
	s.Shutdown()

}
