package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//Redis redis cache
type Redis struct {
	conn *redis.Pool
}

//RedisOpts redis 连接属性
type RedisOpts struct {
	Host        string `yml:"host" json:"host"`
	Password    string `yml:"password" json:"password"`
	Database    int    `yml:"database" json:"database"`
	MaxIdle     int    `yml:"max_idle" json:"max_idle"`
	MaxActive   int    `yml:"max_active" json:"max_active"`
	IdleTimeout int32  `yml:"idle_timeout" json:"idle_timeout"` //second
}

//NewRedis init
func NewRedis(opts *RedisOpts) *Redis {
	pool := &redis.Pool{
		MaxActive:   opts.MaxActive,
		MaxIdle:     opts.MaxIdle,
		IdleTimeout: time.Second * time.Duration(opts.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", opts.Host,
				redis.DialDatabase(opts.Database),
				redis.DialPassword(opts.Password),
			)
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
	return &Redis{pool}
}

//SetConn 设置conn
func (r *Redis) SetConn(conn *redis.Pool) {
	r.conn = conn
}

//Get 获取一个值
func (r *Redis) Get(key string) ([]byte, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.Bytes(conn.Do("GET", key))
}

//SAdd 向集合添加一个或多个值
func (r *Redis) SAdd(key string, values ...interface{}) (int, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.Int(conn.Do("SADD", append([]interface{}{key}, values...)...))
}

//SRem 移除集合中的一个或多个元素
func (r *Redis) SRem(key string, values ...interface{}) (int, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.Int(conn.Do("SREM", append([]interface{}{key}, values...)...))
}

//SMembers 获取集合中的所有元素值
func (r *Redis) SMembers(key string) ([]string, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.Strings(conn.Do("SMEMBERS", key))
}

//SRandMember 随机获取集合中的一个元素
func (r *Redis) SRandMember(key string) (string, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.String(conn.Do("SRANDMEMBER", key))
}

//SRandMembers 随机获取集合中的多个元素
func (r *Redis) SRandMembers(key string, count int) ([]string, error) {
	conn := r.conn.Get()
	defer conn.Close()

	return redis.Strings(conn.Do("SRANDMEMBER", key, count))
}
