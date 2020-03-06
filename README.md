# Esign SDK for Go

使用Golang开发的代理ip池。

## 快速开始

以下是请求[个人2要素信息比对](http://open.esign.cn/docs/identity/信息比对/个人2要素信息比对.html)的例子：

```go
//使用memcache保存access_token，也可选择redis或自定义cache
memcacheHandler := cache.NewMemcache("127.0.0.1:11211")

//配置参数
config := &esign.Config{
	Appid:          "xxxx",
	Secret:         "xxxx",
	Cache:          memcacheHandler,
}
es := esign.NewEsign(config)


```

**Cache 设置**

Cache 主要用来缓存代理 ip 地址：
默认采用 redis 存储。也可以直接实现`cache/cache.go`中的接口

## License

Apache License, Version 2.0
