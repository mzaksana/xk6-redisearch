package main

import (
	"context"
	"log"

	"github.com/redis/rueidis"
	"go.k6.io/k6/js/modules"
)

var ctx = context.Background()

// Register the module to be used from JavaScript
func init() {
	modules.Register("k6/x/redisearch", new(Redis))
}

// Redis is the main struct that we will attach the RediSearch client to
type Redis struct {
	client rueidis.Client
}

// NewClient initializes a new Redis client
func (*Redis) NewClient(redisHost string, redisUsername string, redisPassword string) *Redis {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:  []string{redisHost},
		Username:     redisUsername,
		Password:     redisPassword,
		DisableCache: true,
	})
	if err != nil {
		log.Fatalf("Failed to create Redis client: %v", err)
	}

	return &Redis{client: client}
}

// Search performs a RediSearch query using FT.SEARCH
func (r *Redis) Search(index, query string, limit int64) (string, error) {
	// Construct the search query
	searchCmd := r.client.B().FtSearch().Index(index).Query(query).Limit().OffsetNum(0, limit).Build()

	// Execute the search
	res, err := r.client.Do(ctx, searchCmd).ToString()
	if err != nil {
		return "", err
	}

	return res, nil
}
