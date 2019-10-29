package cache

//Entity 缓存接口
type Entity interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Remove(key interface{}) error
	Name() string
}
