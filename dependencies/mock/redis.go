package mock

import (
	"context"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

func RedisServer() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

type mockRedis struct {
	ctx context.Context
	handler *redis.Client
}

func (r mockRedis) Context() context.Context {
	return r.ctx
}

func (r mockRedis) Handler() *redis.Client {
	return r.handler
}

type Redis interface {
	Context() context.Context
	Handler() *redis.Client
}

func RedisClient(srv *miniredis.Miniredis) Redis {
	cli := redis.NewClient(&redis.Options{
		Addr: srv.Addr(),
	})

	return &mockRedis{
		ctx:     context.Background(),
		handler: cli,
	}
}
