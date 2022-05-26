package cache

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

var globalCacheService *CacheService

type CachedItem struct {
	Key         string
	Item        CachedValue
	RefreshedAt time.Time
}

type CachedValue interface {
	Expired() bool
}

type CacheServiceConfiguration struct {
	CacheItemExpiry time.Duration
}

type CacheService struct {
	Items         []CachedItem
	Configuration CacheServiceConfiguration
}

func New() *CacheService {
	globalCacheService = &CacheService{
		Items: make([]CachedItem, 0),
		Configuration: CacheServiceConfiguration{
			CacheItemExpiry: time.Minute * 5,
		},
	}

	return globalCacheService
}

func GetService() *CacheService {
	if globalCacheService != nil {
		return globalCacheService
	}

	return New()
}

func Get[T any](name string) T {
	var item T
	if globalCacheService == nil {
		return item
	}
	itemTypeString := strings.TrimLeft(fmt.Sprintf("%v", reflect.TypeOf(item)), "*")

	for _, cachedItem := range globalCacheService.Items {
		cachedItemTypeString := strings.TrimLeft(fmt.Sprintf("%v", reflect.TypeOf(cachedItem.Item)), "*")
		if strings.EqualFold(name, cachedItem.Key) {
			if strings.EqualFold(itemTypeString, cachedItemTypeString) && !cachedItem.Item.Expired() {
				return cachedItem.Item.(T)
			} else {
				return item
			}
		}
	}

	return item
}

func Set(name string, item CachedValue) error {
	found := false
	if globalCacheService == nil {
		New()
	}

	for _, cachedItem := range globalCacheService.Items {
		if strings.EqualFold(name, cachedItem.Key) {
			found = true
			// updated item
			cachedItem.Item = item
			cachedItem.RefreshedAt = time.Now()
			break
		}
	}

	if !found {
		globalCacheService.Items = append(globalCacheService.Items, CachedItem{
			Key:         name,
			Item:        item,
			RefreshedAt: time.Now(),
		})
	}
	return nil
}

func (c *CacheService) AddConfiguration(config CacheServiceConfiguration) *CacheService {
	c.Configuration = config

	return c
}
