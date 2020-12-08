package cache

import (
	"gorm2.0/model"
)

// AuthorCache interface
type AuthorCache interface {
	Set(key string, value *model.Author)
	Get(key string) *model.Author
}

//
//
//type authorCache struct {
//	host    string
//	db      int
//	expires time.Duration
//}
//
//// NewRedisCache as a constructor function
//func NewAuthorCache(host string, db int, exp time.Duration) AuthorCache {
//	return &redisCache{
//		host:    host,
//		db:      db,
//		expires: exp,
//	}
//}
//
//// Create redis client
//func (a *authorCache) getClient() *redis.Client {
//	return redis.NewClient(&redis.Options{
//		Addr:     a.host,
//		Password: "",
//		DB:       a.db,
//	})
//}
//
//// Set into redis cache
//func (a *authorCache) Set(key string, value *model.Author) {
//	client := a.getClient()
//
//	json, err := json.Marshal(value)
//	if err != nil {
//		panic(err)
//	}
//
//	client.Set(key, json, a.expires*time.Second)
//}
//
//// Get from Redis Cache
//func (a *authorCache) Get(key string) *model.Author {
//	client := a.getClient()
//
//	value, err := client.Get(key).Result()
//	if err != nil {
//		return nil
//	}
//
//	author := model.Author{}
//	err = json.Unmarshal([]byte(value), &author)
//	if err != nil {
//		panic(err)
//	}
//
//	return &author
//}
