package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Nhanderu/gorduchinha/src/domain"
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/Nhanderu/gorduchinha/src/infra/config"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type RedisCache struct {
	cfg   config.Config
	redis *redis.Client
}

func New(cfg config.Config) contract.CacheManager {
	return &RedisCache{
		cfg: cfg,
		redis: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port),
			Password: cfg.Cache.Pass,
			DB:       cfg.Cache.DB,
		}),
	}
}

func (r *RedisCache) buildKey(key string) string {
	return r.cfg.Cache.Prefix + "-" + key
}

func (r *RedisCache) CleanAll() error {

	keys, err := r.redis.Keys(r.buildKey("*")).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if len(keys) > 0 {
		err = r.redis.Del(keys...).Err()
	}
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *RedisCache) Invalidate(key string) error {

	err := r.redis.Del(r.buildKey(key)).Err()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *RedisCache) Get(key string) ([]byte, error) {

	val, err := r.redis.Get(r.buildKey(key)).Bytes()
	if err == redis.Nil {
		return val, errors.WithStack(domain.ErrCacheMiss)
	}
	if err != nil {
		return val, errors.WithStack(err)
	}

	return val, nil
}

func (r *RedisCache) Set(key string, data []byte) error {

	err := r.redis.Set(r.buildKey(key), data, r.cfg.Cache.DefaultExpiration).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *RedisCache) GetJSON(key string, data interface{}) error {

	val, err := r.Get(key)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(val, &data)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *RedisCache) SetJSON(key string, data interface{}) error {

	dataString, err := json.Marshal(data)
	if err != nil {
		return errors.WithStack(err)
	}

	err = r.Set(key, dataString)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *RedisCache) GetExpiration(key string) (time.Duration, error) {

	expiration, err := r.redis.TTL(r.buildKey(key)).Result()
	if err != nil {
		return expiration, errors.WithStack(err)
	}

	return expiration, nil
}

func (r *RedisCache) SetExpiration(key string, expiration time.Duration) error {

	err := r.redis.Expire(r.buildKey(key), expiration).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
