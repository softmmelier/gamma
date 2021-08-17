package pubsub

import (
	"github.com/alicebob/miniredis"
	"github.com/softmmelier/gamma/v0/dependencies"
	"github.com/softmmelier/gamma/v0/dependencies/mock"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var redisServerSub *miniredis.Miniredis

func TestNewSubscriber(t *testing.T) {
	redisServerSub = mock.RedisServer()
	port, _ := strconv.Atoi(redisServerSub.Port())

	cfg := SubscriberConfig{
		dependencies.RedisClientConfig{
			Host:     redisServerSub.Host(),
			Port:     port,
			Password: "",
		},
	}

	s := &Sub{
		config: cfg,
		h:      mock.RedisClient(redisServerSub),
	}
	defer teardownSubDeps()

	sub := NewSubscriber(cfg)

	require.Equal(t, sub.Name(), s.Name())
}

func TestSubscriber(t *testing.T) {

}

func teardownSubDeps() {
	redisServerSub.Close()
}