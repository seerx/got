package cache

//Provider Session 提供者接口
type Provider interface {
	CacheInit(name string) (Entity, error)
	CacheGet(name string) (Entity, error)
	CacheDestroy(name string) error
	CacheGC(maxLifeTime int64)
}

// Session 提供者注册表
var providers = make(map[string]Provider)

// RegisterProvider 注册Session 寄存器
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("cache: Register provider is nil")
	}

	if _, p := providers[name]; p {
		panic("cache: Register provider is existed")
	}

	providers[name] = provider
}
