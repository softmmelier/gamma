package pubsub

import (
	"errors"
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/softmmelier/gamma/v0/dependencies"
	"github.com/softmmelier/gamma/v0/dependencies/mock"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var redisServerPub *miniredis.Miniredis

func TestNewPublisher(t *testing.T) {
	redisServerPub = mock.RedisServer()
	port, _ := strconv.Atoi(redisServerPub.Port())

	cfg := PublisherConfig{
		dependencies.RedisClientConfig{
			Host:     redisServerPub.Host(),
			Port:     port,
			Password: "",
		},
	}

	p := &Pub{
		config: cfg,
		h:      mock.RedisClient(redisServerPub),
	}
	defer teardownPubDeps()

	pub := NewPublisher(cfg)

	require.Equal(t, pub.Name(), p.Name())
}

func TestPublisher(t *testing.T) {

	type testData struct {
		 topic string
		 msg []byte
	}

	tests := []struct {
		name string
		buildData func() *testData
		checkResponse func(t *testing.T, err error)
	}{
		{
			name:          "OK",
			buildData: func() *testData {
				return &testData{
					topic: "test-topic",
					msg:   []byte("Test message"),
				}
			},
			checkResponse: func(t *testing.T, err error) {
				require.Equal(t, err, nil)
			},
		},
		{
			name:          "NoTopic",
			buildData: func() *testData {
				return &testData{
					topic: "",
					msg:   []byte("Test message"),
				}
			},
			checkResponse: func(t *testing.T, err error) {
				require.Equal(t, err, fmt.Errorf(errTopicNotValid, ""))
			},
		},
		{
			name:          "NoMessage",
			buildData: func() *testData {
				return &testData{
					topic: "test-topic",
					msg:   []byte(""),
				}
			},
			checkResponse: func(t *testing.T, err error) {
				require.Equal(t, err, errors.New(errMsgIsEmpty))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisServerPub = mock.RedisServer()
			s := &Pub{
				config: PublisherConfig{},
				h: mock.RedisClient(redisServerPub),
			}
			defer teardownPubDeps()

			data := tt.buildData()
			err := s.Publish(data.topic, data.msg)

			tt.checkResponse(t, err)
		})
	}
}

func TestPublisherName(t *testing.T) {
	redisServerPub = mock.RedisServer()
	s := &Pub{
		config: PublisherConfig{},
		h: mock.RedisClient(redisServerPub),
	}
	defer teardownPubDeps()

	require.Equal(t, s.Name(), pkgPubName)
}

func teardownPubDeps() {
	redisServerPub.Close()
}