package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"gorm2.0/model"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

// NewRedisCache as a constructor function
func NewRedisCache(host string, db int, exp time.Duration) AuthorCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

// Create redis client
func (c *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: "",
		DB:       c.db,
	})
}

// Set into redis cache
func (c *redisCache) Set(key string, value *model.Author) {
	client := c.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}


	client.Set(key, json, c.expires*time.Second)
}

// Get from Redis Cache
func (c *redisCache) Get(key string) *model.Author {
	client := c.getClient()

	value, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	author := model.Author{}
	err = json.Unmarshal([]byte(value), &author)
	if err != nil {
		panic(err)
	}

	return &author
}
