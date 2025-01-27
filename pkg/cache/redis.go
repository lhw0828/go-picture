package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

// 设置缓存
func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, bytes, expiration).Err()
}

// 获取缓存
func (c *RedisCache) Get(ctx context.Context, key string, value interface{}) error {
	bytes, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, value)
}

// 删除缓存
func (c *RedisCache) Del(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// 设置过期时间
func (c *RedisCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

// 检查key是否存在
func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// 批量获取
func (c *RedisCache) MGet(ctx context.Context, keys []string) ([]interface{}, error) {
	return c.client.MGet(ctx, keys...).Result()
}

// 批量设置
func (c *RedisCache) MSet(ctx context.Context, values map[string]interface{}) error {
	pairs := make([]interface{}, 0, len(values)*2)
	for k, v := range values {
		bytes, err := json.Marshal(v)
		if err != nil {
			return err
		}
		pairs = append(pairs, k, bytes)
	}
	return c.client.MSet(ctx, pairs...).Err()
}

// 自增
func (c *RedisCache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

// 自减
func (c *RedisCache) Decr(ctx context.Context, key string) (int64, error) {
	return c.client.Decr(ctx, key).Result()
}
