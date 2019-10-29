package got

import "github.com/seerx/got/pkg/cache"

//CacheManager 缓存管理
var CacheManager *cache.Manager

// InitCacheManager 初始化缓存管理器
// 不用自己维护，使用全局的 CacheManager
// @param expiredTime 秒
func InitCacheManager(provider string, expiredTime int64) {
	CacheManager = cache.NewCacheManager(provider, expiredTime)
}

//CacheGet 获取缓存对象
// 调用前确保已经初始化 got.InitCacheManager(...)
func CacheGet(name string) cache.Entity {
	return CacheManager.GetEntity(name)
}

//CacheNew 创建缓存对象
func CacheNew(name string) cache.Entity {
	return CacheManager.NewEntity(name)
}

//CacheDestroy 销毁缓存对象
func CacheDestroy(name string) {
	CacheManager.DestroyEntity(name)
}
