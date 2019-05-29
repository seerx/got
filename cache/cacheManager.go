package cache

import (
	"fmt"
	"sync"
	"time"
)

//Manager 缓存管理
type Manager struct {
	lock        sync.RWMutex
	provider    Provider
	MaxLifeTime int64
	gcDuration  time.Duration
}

//CacheManager 缓存管理
var CacheManager *Manager

// InitCacheManager 初始化缓存管理器
// 不用自己维护，使用全局的 CacheManager
// @param expiredTime 秒
func InitCacheManager(provider string, expiredTime int64) {
	CacheManager = NewCacheManager(provider, expiredTime)
}

// NewCacheManager 创建缓存管理器
// 需要自行维护 Manager
// @param expiredTime 秒
func NewCacheManager(provider string, expiredTime int64) *Manager {
	prv, ok := providers[provider]
	if !ok {
		panic(fmt.Errorf("session: unknown provide %q (forgotten import?)", provider))
	}

	// 秒转换为 nanoSecond
	gcd := time.Second * time.Duration(expiredTime)

	manager := &Manager{
		gcDuration:  gcd,
		MaxLifeTime: expiredTime,
		provider:    prv,
	}

	go manager.GC()

	return manager
}

// GC Cache 垃圾回收
func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.CacheGC(m.MaxLifeTime)

	time.AfterFunc(m.gcDuration, func() {
		m.GC()
	})
}

// GetEntity 获取缓存实例
func (m *Manager) GetEntity(name string) Entity {
	m.lock.RLock()
	defer m.lock.RUnlock()

	session, e := m.provider.CacheGet(name)
	if e == nil {
		return session
	}

	return nil
}

//NewEntity 创建缓存实例
func (m *Manager) NewEntity(name string) Entity {
	m.lock.Lock()
	defer m.lock.Unlock()

	entity, _ := m.provider.CacheInit(name)
	return entity
}

// GetOrNewEntiry 获取缓存实例
// 如果没有则创建一个
func (m *Manager) GetOrNewEntiry(name string) Entity {
	entity := m.GetEntity(name)
	if entity == nil {
		entity = m.NewEntity(name)
	}
	return entity
}

// DestroyEntity 释放缓存实例
func (m *Manager) DestroyEntity(name string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.CacheInit(name)
	return true
}
