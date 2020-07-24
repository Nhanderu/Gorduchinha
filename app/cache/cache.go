package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type redisCache struct {
	redis             *redis.Client
	prefix            string
	defaultExpiration time.Duration
}

func New(
	host string,
	port int,
	db int,
	pass string,
	prefix string,
	defaultExpiration time.Duration,
) contract.CacheManager {

	return redisCache{
		redis: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: pass,
			DB:       db,
		}),
		prefix:            prefix,
		defaultExpiration: defaultExpiration,
	}
}

func (r redisCache) buildKey(key string) string {
	return r.prefix + "-" + key
}

func (r redisCache) CleanAll() error {

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

func (r redisCache) Invalidate(key string) error {

	err := r.redis.Del(r.buildKey(key)).Err()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r redisCache) Get(key string) ([]byte, error) {

	val, err := r.redis.Get(r.buildKey(key)).Bytes()
	if err == redis.Nil {
		return val, errors.WithStack(constant.ErrCacheMiss)
	}
	if err != nil {
		return val, errors.WithStack(err)
	}

	return val, nil
}

func (r redisCache) Set(key string, data []byte) error {

	err := r.redis.Set(r.buildKey(key), data, r.defaultExpiration).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r redisCache) GetJSON(key string, data interface{}) error {

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

func (r redisCache) SetJSON(key string, data interface{}) error {

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

func (r redisCache) GetExpiration(key string) (time.Duration, error) {

	expiration, err := r.redis.TTL(r.buildKey(key)).Result()
	if err != nil {
		return expiration, errors.WithStack(err)
	}

	return expiration, nil
}

func (r redisCache) SetExpiration(key string, expiration time.Duration) error {

	err := r.redis.Expire(r.buildKey(key), expiration).Err()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
