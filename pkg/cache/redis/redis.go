package redis

import (
	fibcache "github.com/andrewrynhard/fibonacci/pkg/cache"
	"github.com/go-redis/redis"
)

// Cache is the concreate implementation of the fibacache.Cache interface.
type Cache struct {
	client *redis.Client
}

// NewRedisCache initializes and returns a Redis cache configured with an
// address and a codec.
func NewRedisCache(addr string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return &Cache{client: client}
}

// Set implements the fibacache.Cache interface.
func (c *Cache) Set(kv fibcache.KeyValuePair) (err error) {
	if err = c.client.Set(string(kv.Key), kv.Value, 0).Err(); err != nil {
		return
	}

	return nil
}

// Get implements the fibacache.Cache interface.
func (c *Cache) Get(k fibcache.Key) (kv *fibcache.KeyValuePair, err error) {
	val, err := c.client.Get(string(k)).Result()
	if err != nil {
		return
	}
	return &fibcache.KeyValuePair{Key: k, Value: val}, nil
}
