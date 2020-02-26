package proxy

import (
	"testing"

	"github.com/trrtly/proxy/cache"
)

func TestNewProxy(t *testing.T) {
	redisOpt := &cache.RedisOpts{
		Host:        "",
		Password:    "",
		Database:    0,
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: 0,
	}
	cacheIns := cache.NewRedis(redisOpt)
}
