package pubsub

import (
	"fmt"
	"github.com/softmmelier/gamma/v0/app"
	"github.com/softmmelier/gamma/v0/dependencies"
)

const (
	pkgSubName = "Transport/PubSub/Subscriber"
)

type Subscription interface {
	Topics() []string
	Handler(string, []byte)
}

type Subscriber interface {
	app.Runner
	Subscribe(Subscription) error
}

type SubscriberConfig struct {
	dependencies.RedisClientConfig
}

type Sub struct {
	config SubscriberConfig
	h      dependencies.Redis
}

//NewSubscriber new redis subscriber
func NewSubscriber(c SubscriberConfig) Subscriber {
	return &Sub{
		config: c,
		h:      dependencies.NewRedisClient(c.RedisClientConfig),
	}
}

func (sub Sub) Name() string {
	return pkgSubName
}

func (sub Sub) Run()  {
	fmt.Println("Run subscriber")
}

func (sub Sub) Subscribe(s Subscription) error {
	for _, topic := range s.Topics() {
		t := sub.h.Handler().Subscribe(sub.h.Context(), topic)
		channel := t.Channel()

		go func() {
			for msg := range channel {
				s.Handler(msg.Channel, []byte(msg.Payload))
			}
		}()
	}

	return nil
}
