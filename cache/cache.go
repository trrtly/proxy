package cache

import "time"

//Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	SAdd(key string) bool
	SMembers(key string) interface{}
}
