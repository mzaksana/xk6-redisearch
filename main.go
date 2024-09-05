package main

import (
	redisearch "github.com/mzaksana/xk6-redisearch/redisearch"
	"go.k6.io/k6/js/modules"
)

// Register the module to be used from JavaScript
func init() {
	modules.Register("k6/x/redisearch", new(redisearch.RootModule))
}
