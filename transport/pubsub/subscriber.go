package pubsub

import (
	"fmt"
	"github.com/softmmelier/gamma/app"
	"log"
)

const (
	pkgSubName = "Transport/PubSub/Subscriber"
)

type Subscribe interface {
	Topics() []string
	Handler(string, []byte)
}

type Subscriber interface {
	app.Service
	Subscribe(Subscribe)
}

type SubscriberConfig struct {
	RedisClientConfig
}

type Sub struct {
	config SubscriberConfig
}

//NewSubscriber new redis subscriber
func NewSubscriber(c SubscriberConfig) Subscriber {
	return &Sub{
		config: c,
	}
}

func (sub Sub) Name() string {
	return pkgSubName
}

func (sub Sub) Run() {
	fmt.Println(pkgSubName, "== No runner")
}

func (sub Sub) Subscribe(subscribe Subscribe) {
	for _, topic := range subscribe.Topics() {
		log.Println("Subscribe ->", topic)
	}
}
