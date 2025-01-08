package message

import (
	"github.com/nats-io/nats.go"
	"sync"
	"testing"
)

func TestNats(t *testing.T) {
	t.Log("NatsTest")
	q := NewNatsQueen("nats://localhost:4222")
	q.Subscribe("test.subject", func(msg *nats.Msg) {
		t.Log("Received message: ", msg.Data)
	})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go q.Start(wg)
	wg.Wait()
}

func CreateTable(t *testing.T) {

}
