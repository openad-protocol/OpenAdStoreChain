package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

//Subscribe 订阅
//@return *redis.PubSub, error
func (r *redisView) Subscribe(ctx context.Context,channels ...string) (*redis.PubSub, error) {
	switch v := r.cmd.(type) {
	case *redis.Client:
		return v.Subscribe(ctx,channels...),nil
	default:
		return nil, errors.New("UnSupported")
	}
}

//Subscribe 订阅
//@return *redis.PubSub, error
func (r *redisView) PSubscribe(ctx context.Context,channels ...string) (*redis.PubSub, error) {
	switch v := r.cmd.(type) {
	case *redis.Client:
		return v.PSubscribe(ctx,channels...),nil
	default:
		return nil, errors.New("UnSupported")
	}
}
