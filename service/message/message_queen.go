package message

import (
	"github.com/nats-io/nats.go"
	"sync"
)

type IMessage interface {
	Start(wg *sync.WaitGroup) bool
	Stop() bool
	Publish(subject string, msg []byte) bool
	AddSubscription(subject string) bool
	Subscribe(subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error)
	Unsubscribe(sub *nats.Subscription) bool
	UnsubscribeAll() bool
}
