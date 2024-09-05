package redisearch

import (
	"fmt"

	"go.k6.io/k6/js/modules"
)

//go:generate go run github.com/szkiba/tygor@latest --package redisearch --skeleton index.d.ts
//go:generate go run github.com/szkiba/tygor@latest doc --inject README.md index.d.ts

func init() {
	register(newModule)
}

func newModule(_ modules.VU) goModule {
	return &goModuleImpl{goRedisearch: &goRedisearchImpl{greeting: "Hello, World!"}}
}

type goModuleImpl struct {
	goRedisearch goRedisearch
}

var _ goModule = (*goModuleImpl)(nil)

func (mod *goModuleImpl) newRedisearch(nameArg string) (goRedisearch, error) {
	msg := fmt.Sprintf("Hello, %s!", nameArg)

	return &goRedisearchImpl{greeting: msg}, nil
}

func (mod *goModuleImpl) defaultRedisearchGetter() (goRedisearch, error) {
	return mod.goRedisearch, nil
}

type goRedisearchImpl struct {
	greeting string
}

var _ goRedisearch = (*goRedisearchImpl)(nil)

func (impl *goRedisearchImpl) greetingGetter() (string, error) {
	return impl.greeting, nil
}
