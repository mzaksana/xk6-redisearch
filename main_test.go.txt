package main

import (
	"testing"

	"github.com/redis/rueidis"
	"github.com/stretchr/testify/assert"
)

// NewClient initializes a new Redis client
func NewClient(redisHost string, redisUsername string, redisPassword string) *Redis {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{redisHost},
		Username:     redisUsername,
		Password:     redisPassword,
		DisableCache: true,
	})
	if err != nil {
		panic(err)
	}

	return &Redis{client: client}
}

// TestNewClient tests the initialization of the Redis client
func TestNewClient(t *testing.T) {
	// Replace with your actual Redis host and credentials
	redisHost := ""
	redisUsername := ""
	redisPassword := ""

	// Initialize Redis client
	redis := NewClient(redisHost, redisUsername, redisPassword)
	assert.NotNil(t, redis, "Redis client should be initialized")

	// Ping Redis to verify connection
	res := redis.client.Do(ctx, redis.client.B().Ping().Build())

	// Extract the message and error
	msg, err := res.ToMessage()
	if err != nil {
		t.Fatalf("Failed to get PING response: %v", err)
	}

	pong := msg.String()

	// Assert that the response is "PONG"
	assert.Equal(t, "{\"Value\":\"PONG\",\"Type\":\"simple string\"}", pong, "Ping response should be 'PONG'")
}
