package memcache

import (
	"sync"
	"time"
)

// memCache 内存缓存
type memCache struct {
	name string
	// data sync.Map
	time time.Time
	lock *sync.RWMutex
	data map[interface{}]interface{}
}

func newMemCache(name string) *memCache {
	ss := memCache{
		name: name,
		time: time.Now(),
		data: make(map[interface{}]interface{}),
		lock: new(sync.RWMutex)}
	return &ss
}

//Name 获取缓存名称
func (mc *memCache) Name() string {
	return mc.name
}

// Get 获取 缓存
func (mc *memCache) Get(key interface{}) interface{} {
	mc.lock.RLock()
	defer mc.lock.RUnlock()
	return mc.data[key]
}

// Remove 删除 缓存
func (mc *memCache) Remove(key interface{}) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	delete(mc.data, key)
	return nil
}

// Set 设置 缓存
func (mc *memCache) Set(key, value interface{}) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.data[key] = value

	return nil
}
