package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.k6.io/k6/js/modules"
)

var ctx = context.Background()

// Register the module to be used from JavaScript
func init() {
    modules.Register("k6/x/redis", new(Redis))
}

// Redis is the main struct that we will attach the Redis client to
type Redis struct {
    client *redis.Client
}

// NewClient initializes a new Redis client
func (*Redis) NewClient(addr string) *Redis {
    rdb := redis.NewClient(&redis.Options{
        Addr: addr,
        Password: "", // no password set
        DB: 0,  // use default DB
    })

    return &Redis{client: rdb}
}

// Ping sends a PING command to Redis
func (r *Redis) Ping() (string, error) {
    return r.client.Ping(ctx).Result()
}

// Set sets a key-value pair in Redis
func (r *Redis) Set(key, value string) error {
    return r.client.Set(ctx, key, value, 0).Err()
}

// Get retrieves a value from Redis by key
func (r *Redis) Get(key string) (string, error) {
    return r.client.Get(ctx, key).Result()
}

