package pubsub

import (
	"fmt"
	"github.com/softmmelier/gamma/app"
)

const (
	pkgPubName = "Transport/PubSub/Publisher"
)

type Publisher interface {
	app.Service
	Publish(topic string, msg []byte)
}

type PublisherConfig struct {
	RedisClientConfig
}

type Pub struct {
	config PublisherConfig
	h      Redis
}

//NewPublisher new redis publisher
func NewPublisher(c PublisherConfig) Publisher {
	return &Pub{
		config: c,
	}
}

func (pub Pub) Name() string {
	return pkgPubName
}

func (pub Pub) Run() {
	fmt.Println(pkgPubName, "== No runner")
}

func (pub Pub) Publish(topic string, msg []byte) {

}
