package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis redis cache
type Redis struct {
	conn *redis.Client
}

// RedisOpts redis 连接属性
type RedisOpts struct {
	Addr         string
	DB           int
	DialTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	PoolSize     int
	MinIdleConns int
}

// NewRedis 实例化
func NewRedis(opts *RedisOpts) *Redis {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         opts.Addr,
		DB:           opts.DB,
		DialTimeout:  opts.DialTimeout,
		WriteTimeout: opts.WriteTimeout,
		ReadTimeout:  opts.ReadTimeout,
		PoolSize:     opts.PoolSize,
		MinIdleConns: opts.MinIdleConns,
	})

	return &Redis{conn: redisClient}
}

// Get 获取一个值
func (r *Redis) Get(key string) interface{} {
	return r.conn.Get(context.Background(), key).Val()
}

// Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	return r.conn.Set(context.Background(), key, val, timeout).Err()
}

// IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	if r.conn.Exists(context.Background(), key).Val() > 0 {
		return true
	}
	return false
}

// Delete 删除
func (r *Redis) Delete(key string) error {
	return r.conn.Del(context.Background(), key).Err()
}
