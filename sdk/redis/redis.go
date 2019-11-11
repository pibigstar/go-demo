package redis

import (
	"fmt"
	goRedis "github.com/go-redis/redis"
	"time"
)

var Redis *RedisClient

const expiration = 7 * 24 * time.Hour

// RedisClient extend client and have itself func
type RedisClient struct {
	*goRedis.Client
}

// Init the redis client
func NewRedisClient() error {
	if Redis != nil {
		return nil
	}
	client := goRedis.NewClient(&goRedis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
		//the pool config
		PoolSize:    10,
		PoolTimeout: 3000,
		IdleTimeout: 50000,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	Redis = &RedisClient{client}
	return nil
}

// init the redis  client
func init() {
	err := NewRedisClient()
	if err != nil {
		fmt.Println("failed to connect redis client")
	}
}

// set the string to redis，the expire default is seven days
func (redis *RedisClient) SSet(key string, value interface{}) *goRedis.StatusCmd {
	return redis.Set(key, value, expiration)
}

// get the string value by key
func (redis *RedisClient) SGet(key string) string {
	return redis.Get(key).String()
}

// close the redis client
func (redis *RedisClient) Close() {
	redis.Close()
}

// get the redis client，if client not initialization
// and create the redis client
func GetRedisClient() (*RedisClient, error) {
	if Redis == nil {
		err := NewRedisClient()
		if err != nil {
			return nil, err
		}
		return Redis, nil
	}
	return Redis, nil
}
