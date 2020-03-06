package proxy

import (
	"errors"
	"net"
	"time"

	"github.com/trrtly/proxy/cache"
	"github.com/trrtly/proxy/provider"
)

// Proxy struct
type Proxy struct {
	*Config
}

// Config struct
type Config struct {
	// Provider 代理提供商
	Provider provider.Provider
	// Cache 缓存组件
	Cache cache.Cache
	// Timeout 超时时间
	Timeout time.Duration
}

var (
	ErrNotFoundAddr = errors.New("未找到可用代理ip")
)

const (
	ProxyCacheKey       = "trrtly_proxy_pool_set"
	RecursiveFetchCount = 5
)

// NewProxy init
func NewProxy(c *Config) *Proxy {
	return &Proxy{c}
}

// ProduceHandler 添加数据至 ip 池中
func (p *Proxy) ProduceHandler() {
	proxyList, err := p.Provider.GetProxys()
	if err != nil {
		return
	}
	for _, addr := range proxyList {
		go p.process(addr)
	}
	return
}

// CronCheckAddrs 定时检查代理 ip 是否可用
func (p *Proxy) CronCheckAddrs() {
	p.recuCheckAddrs(1)
}

// GetRandomOne 从代理池中随机获取一个代理
func (p *Proxy) GetRandomOne() (string, error) {
	addr, err := p.Cache.SRandMember(ProxyCacheKey)
	if err != nil {
		return "", err
	}
	if err := p.checkAndRemoveAddr(addr); err != nil {
		return "", err
	}
	return addr, err
}

// MustGetRandomOne 强制从代理池中随机获取一个代理
func (p *Proxy) MustGetRandomOne() (string, error) {
	return p.getRandomOneRecursive(1)
}

// CheckAddr 检测代理是否可用
func (p *Proxy) CheckAddr(addr string) error {
	d := net.Dialer{Timeout: p.Timeout}
	conn, err := d.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (p *Proxy) recuCheckAddrs(count int8) {
	addrs, err := p.Cache.SMembers(ProxyCacheKey)
	if err != nil {
		return
	}
	for _, addr := range addrs {
		go p.checkAndRemoveAddr(addr)
	}
	count ++
	if count > 3 {
		return
	}
	time.Sleep(time.Millisecond*30)

	p.recuCheckAddrs(count)
}

// checkAndRemoveAddr 检测并移除过期的代理
func (p *Proxy) checkAndRemoveAddr(addr string) error {
	err := p.CheckAddr(addr)
	if err != nil {
		p.Cache.SRem(ProxyCacheKey, addr)
	}
	return err
}

func (p *Proxy) process(addr string) error {
	if err := p.CheckAddr(addr); err != nil {
		return err
	}
	p.Cache.SAdd(ProxyCacheKey, addr)
	return nil
}

// getRandomOneRecursive 递归获取随机代理 ip，直至代理池中无可用代理
func (p *Proxy) getRandomOneRecursive(count int) (string, error) {
	addr, err := p.GetRandomOne()
	if err != nil {
		if count > RecursiveFetchCount {
			return "", ErrNotFoundAddr
		}
		count++
		p.getRandomOneRecursive(count)
	}
	return addr, err
}
