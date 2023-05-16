package database

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}


func RedisInit() {
	RedisPool = newPool("redis:6379")
}