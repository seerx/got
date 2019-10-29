package memcache

import (
	"fmt"
	"time"

	"github.com/seerx/got/pkg/cache"
)

//PROVIDER 提供者名称
const PROVIDER = "memory"

//memoryProvider 内存session 提供者
type memProvider struct {
	cacheMap map[string]*memCache
}

func init() {
	RegisterProvider()
}

// RegisterProvider 注册
func RegisterProvider() {
	provider := memProvider{cacheMap: make(map[string]*memCache)}
	cache.RegisterProvider(PROVIDER, &provider)
}

func (p *memProvider) CacheInit(name string) (cache.Entity, error) {
	entity := newMemCache(name)
	p.cacheMap[name] = entity
	return entity, nil
}

func (p *memProvider) CacheGet(name string) (cache.Entity, error) {
	entity, ok := p.cacheMap[name]
	if ok {
		entity.time = time.Now()
		return entity, nil
	}

	return nil, fmt.Errorf("No session found by id: %s", name)
}

func (p *memProvider) CacheDestroy(name string) error {
	delete(p.cacheMap, name)
	return nil
}

func (p *memProvider) CacheGC(maxLifeTime int64) {
	now := time.Now()
	for _, entity := range p.cacheMap {
		if now.Unix()-entity.time.Unix() > maxLifeTime {
			// 销毁超时的 Session
			p.CacheDestroy(entity.Name())
		}
	}
}
