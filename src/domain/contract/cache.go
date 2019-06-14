package contract

import (
	"time"
)

type CacheManager interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error

	GetJSON(key string, data interface{}) error
	SetJSON(key string, data interface{}) error

	GetExpiration(key string) (time.Duration, error)
	SetExpiration(key string, expiration time.Duration) error

	Invalidate(key string) error
	CleanAll() error
}
