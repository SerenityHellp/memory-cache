package cache

import (
	"github.com/SerenityHellp/memory-cache/data"
	"reflect"
	"sync"
	"time"
)

type NoevictionCache struct {
	defaultDuration time.Duration
	mu              sync.RWMutex
	dataBlock       data.DataBlock
}

func (cache *NoevictionCache) Set(key string, value interface{}) (b bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	r := cache.dataBlock.Get(key)
	_, ok := r[key]
	//assert:  has enough space for set if found it from cache
	if !ok && cache.dataBlock.IsFull() {
		//delete expire item if full
		cache.dataBlock.DeleteExpire()
		if cache.dataBlock.IsFull() {
			//double check
			return false
		}
	}
	cache.dataBlock.Set(map[string]interface{}{key: value}, cache.defaultDuration)
	return true
}

func (cache *NoevictionCache) Get(key string) (value interface{}) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	r := cache.dataBlock.Get(key)
	value, b := r[key]
	if !b {
		//return nil ,if not found item
		return nil
	}
	if value.(data.Item).Expire() {
		//delete it if expire
		cache.dataBlock.Del(key)
		//return nil ,if expire
		return nil
	}
	return value.(data.Item).GetValue()
}

func (cache *NoevictionCache) GetSet(key string, value interface{}) (oValue interface{}, b bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	r := cache.dataBlock.Get(key)
	oValue, ok := r[key]
	//assert:  has enough space for set if found it from cache
	if !ok && cache.dataBlock.IsFull() {
		//delete expire item if full
		cache.dataBlock.DeleteExpire()
		if cache.dataBlock.IsFull() {
			//double check
			return nil, false
		}
	}
	cache.dataBlock.Set(map[string]interface{}{key: value}, cache.defaultDuration)
	return true
}

func (cache *NoevictionCache) MGet(key ...string) map[string]interface{} {
	panic("implement me")
}

func (cache *NoevictionCache) SetEx(key string, duration time.Duration, value interface{}) {
	panic("implement me")
}

func (cache *NoevictionCache) SetNx(key string, value interface{}) error {
	panic("implement me")
}

func (cache *NoevictionCache) Len(key string) int64 {
	panic("implement me")
}

func (cache *NoevictionCache) MSet(map[string]interface{}) {
	panic("implement me")
}

func (cache *NoevictionCache) MSetNx(map[string]interface{}) error {
	panic("implement me")
}

func (cache *NoevictionCache) Incr(key string) (int64, e error) {
	panic("implement me")
}

func (cache *NoevictionCache) IncrBy(key string, incr int64) error {
	panic("implement me")
}

func (cache *NoevictionCache) IncrByFloat(key string, incr float64) error {
	panic("implement me")
}

func (cache *NoevictionCache) Decr(key string) error {
	panic("implement me")
}

func (cache *NoevictionCache) Del(key ...string) int {
	panic("implement me")
}

func (cache *NoevictionCache) Exists(key ...string) int {
	panic("implement me")
}

func (cache *NoevictionCache) Expire(key string, duration time.Duration) bool {
	panic("implement me")
}

func (cache *NoevictionCache) ExpireAt(key string, timestamp int64) {
	panic("implement me")
}

func (cache *NoevictionCache) RandomKey() string {
	panic("implement me")
}

func (cache *NoevictionCache) Rename(key, newkey string) error {
	panic("implement me")
}

func (cache *NoevictionCache) RenameNx(key, newkey string) error {
	panic("implement me")
}

func (cache *NoevictionCache) Ttl(k string) int64 {
	panic("implement me")
}

func (cache *NoevictionCache) Type(key string) reflect.Type {
	panic("implement me")
}
