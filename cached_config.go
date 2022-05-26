package cache

import (
	"time"
)

type CachedConfig struct {
	Value     string
	ExpiresIn time.Time
}

func (c *CacheService) AddCachedConfig(key, value string) {
	svc := GetService()
	item := CachedConfig{
		Value:     value,
		ExpiresIn: time.Now().Add(svc.Configuration.CacheItemExpiry),
	}

	Set(key, item)
}

func (c CachedConfig) Expired() bool {
	if c.ExpiresIn.IsZero() {
		return false
	}

	return time.Now().After(c.ExpiresIn)
}
