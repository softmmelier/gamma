package dependencies

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	Context() context.Context
	Handler() *redis.Client
}

type RedisClientConfig struct {
	Host     string
	Port     int
	Password string
}

type RedisClient struct {
	handler *redis.Client
	ctx     context.Context
}

func NewRedisClient(c RedisClientConfig) Redis {
	opt := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
		DB:   0,
	}

	if c.Password != "" {
		opt.Password = c.Password
	}

	redisClient := redis.NewClient(opt)

	// Generate context background
	ctx := context.Background()

	// Test connection
	err := redisClient.Ping(ctx).Err()
	if err != nil {

		// In case of error wait for 5 sec and try again
		time.Sleep(5 * time.Second)
		err := redisClient.Ping(ctx).Err()
		if err != nil {
			panic(err)
		}
	}

	return &RedisClient{
		handler: redisClient,
		ctx:     ctx,
	}
}

// Handler redis requested action
func (c *RedisClient) Handler() *redis.Client {
	return c.handler
}

// Context for current client instance
func (c *RedisClient) Context() context.Context {
	return c.ctx
}
