package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIfServiceDoesNotExistsOnGetServiceIsCreated(t *testing.T) {
	getSvc := GetService()
	getSvc.AddCachedConfig("config", "key")
	configValue := Get[CachedConfig]("config")
	assert.Equal(t, 1, len(getSvc.Items))
	assert.Equal(t, "key", configValue.Value)
}

func TestIfServiceExistsReturnService(t *testing.T) {
	svc := New()
	svc.AddCachedConfig("config", "key")
	assert.Equal(t, 1, len(svc.Items))

	getSvc := GetService()
	configValue := Get[CachedConfig]("config")
	assert.Equal(t, 1, len(getSvc.Items))
	assert.Equal(t, "key", configValue.Value)
}

func TestIfServiceIsNilWhenGetValueThenEmptyValueIsReturned(t *testing.T) {
	globalCacheService = nil
	configValue := Get[CachedConfig]("config")
	assert.Equal(t, configValue, CachedConfig{})
}

func TestMainGeneric(t *testing.T) {
	cacheService := New()
	tokenItem := CachedJwtToken{
		TokenType:   "ddd",
		ExpiresIn:   20,
		AccessToken: "dddd",
	}
	configItem := CachedConfig{
		Value: "something",
	}

	Set("token1", tokenItem)
	Set("config1", configItem)

	assert.Equal(t, 2, len(cacheService.Items))
	s := Get[CachedConfig]("config1")
	assert.Equal(t, "something", s.Value)
}

func TestMainExpiredCached(t *testing.T) {
	cacheService := New()
	cacheService.Configuration = CacheServiceConfiguration{
		CacheItemExpiry: (time.Millisecond * 100),
	}
	tokenItem := CachedJwtToken{
		TokenType:   "ddd",
		ExpiresIn:   20,
		AccessToken: "dddd",
	}

	cacheService.AddCachedConfig("config1", "something")

	Set("token1", tokenItem)

	time.Sleep(time.Millisecond * 200)
	assert.Equal(t, 2, len(cacheService.Items))
	s := Get[CachedConfig]("config1")
	assert.Equal(t, "", s.Value)
}
