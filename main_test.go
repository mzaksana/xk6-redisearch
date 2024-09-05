package main

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// NewClient initializes a new Redis client
func NewClient(addr string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		 Addr:     addr,
		 Password: "", // no password set
		 DB:       0,  // use default DB
	})

	return &Redis{client: rdb}
}

// TestNewClient tests the initialization of the Redis client
func TestNewClient(t *testing.T) {
    // Initialize Redis client
    client := NewClient("127.0.0.1:6379") // Make sure Redis is running locally
    assert.NotNil(t, client, "Redis client should be initialized")

    // Ping Redis to verify connection
    res, err := client.Ping()
	 fmt.Println("err", err)
    assert.Nil(t, err, "Ping should not return an error")
    assert.Equal(t, "PONG", res, "Ping response should be 'PONG'")
}

// TestSetGet tests setting and getting values from Redis
func TestSetGet(t *testing.T) {
    // Initialize Redis client
    client := NewClient("127.0.0.1:6379") // Ensure the Redis server is running locally

    // Set a key-value pair in Redis
    err := client.Set("test-key", "test-value")
    assert.Nil(t, err, "Set should not return an error")

    // Get the value from Redis by key
    value, err := client.Get("test-key")
    assert.Nil(t, err, "Get should not return an error")
    assert.Equal(t, "test-value", value, "The value returned should be 'test-value'")
}

// TestGetNonExistentKey tests retrieving a non-existent key from Redis
func TestGetNonExistentKey(t *testing.T) {
    // Initialize Redis client
    client := NewClient("127.0.0.1:6379") // Ensure the Redis server is running locally

    // Try to get a non-existent key
    value, err := client.Get("non-existent-key")
    assert.NotNil(t, err, "Get should return an error for a non-existent key")
    assert.Empty(t, value, "The value should be empty for a non-existent key")
}
