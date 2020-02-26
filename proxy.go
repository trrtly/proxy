package proxy

import (
	"github.com/trrtly/proxy/cache"
	"fmt"
	"github.com/trrtly/proxy/provider"
)

// Proxy struct
type Proxy struct {
	// Provider 代理提供商
	Provider provider.Provider
	// Cache 缓存组件
	Cache cache.Cache
}

// NewProxy init
func NewProxy(p provider.Provider, c cache.Cache) *Proxy {
	return &Proxy{p, c}
}

// ProduceHandler 添加数据至 ip 池中
func (p *Proxy) ProduceHandler() error {
	proxyList, err := p.Provider.GetProxys()
	if err != nil {
		return err
	}
	for _, ipPort := range proxyList {
		go p.checkProxy(ipPort)
	}
	return nil
}

func (p *Proxy) checkProxy(ipPort string) {
	fmt.Println(ipPort)
}
