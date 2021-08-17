package pubsub

import (
	"errors"
	"fmt"

	"github.com/softmmelier/gamma/v0/app"
	"github.com/softmmelier/gamma/v0/dependencies"
	"github.com/softmmelier/gamma/v0/validate"
)

const (
	pkgPubName = "Transport/PubSub/Publisher"
)

const (
	errTopicNotValid = "topic is not valid. Received %v"
	errMsgIsEmpty = "message is empty"
)

type Publisher interface {
	app.Service
	Publish(topic string, msg []byte) error
}

type PublisherConfig struct {
	dependencies.RedisClientConfig
}

type Pub struct {
	config PublisherConfig
	h      dependencies.Redis
}

//NewPublisher new redis publisher
func NewPublisher(c PublisherConfig) Publisher {
	return &Pub{
		config: c,
		h:      dependencies.NewRedisClient(c.RedisClientConfig),
	}
}

func (pub Pub) Name() string {
	return pkgPubName
}


func (pub Pub) Publish(topic string, msg []byte) error {

	if topic == "" {
		return fmt.Errorf(errTopicNotValid, topic)
	}

	if validate.IsEmptyByte(msg) {
		return errors.New(errMsgIsEmpty)
	}

	pub.h.Handler().Publish(pub.h.Context(), topic, msg)
	return nil
}
