package proxy

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/robfig/cron"

	"github.com/trrtly/proxy/cache"
)

func TestNewProxy(t *testing.T) {
	predisOpt := &cache.RedisOpts{
		Host:        "127.0.0.1:6379",
		Password:    "xxxxxx",
		Database:    15,
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 200,
	}
	providerIns := &TestProvider{
		cache: cache.NewRedis(predisOpt),
	}

	redisOpt := &cache.RedisOpts{
		Host:        "127.0.0.1:6379",
		Password:    "xxxxxx",
		Database:    15,
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 200,
	}
	config := &Config{
		Provider: providerIns,
		Cache:    cache.NewRedis(redisOpt),
		Timeout:  time.Millisecond * 200,
	}
	proxyIns := NewProxy(config)
	c := cron.New()
	c.AddFunc("0 * * * * *", proxyIns.ProduceHandler)
	c.AddFunc("10 * * * * *", proxyIns.ProduceHandler)
	c.AddFunc("20 * * * * *", proxyIns.ProduceHandler)
	c.AddFunc("30 * * * * *", proxyIns.ProduceHandler)
	c.AddFunc("40 * * * * *", proxyIns.ProduceHandler)
	c.AddFunc("50 * * * * *", proxyIns.ProduceHandler)
	c.AddFunc("* * * * * *", proxyIns.CronCheckAddrs)
	c.AddFunc("* * * * * *", func(){
		addr, err := proxyIns.GetRandomOne()
		fmt.Println(addr, err)
	})
	go c.Start()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
}
// TestProvider struct
type TestProvider struct {
	cache cache.Cache
}
// HoroIPCacheKey test key
const HoroIPCacheKey = "HoroIPCacheKey"

// IPCacheData test cache data
type IPCacheData struct {
	Data []string `json:"data"`
}

// GetProxys interface
func (p *TestProvider) GetProxys() ([]string, error) {
	cacheDatas, err := p.cache.Get(HoroIPCacheKey)
	if err != nil {
		return nil, err
	}
	res := &IPCacheData{}
	err = json.Unmarshal(cacheDatas, res)
	if err != nil {
		return nil, err
	}
	results := []string{}
	for _, v := range res.Data {
		results = append(results, v)
	}
	return results, nil
}
