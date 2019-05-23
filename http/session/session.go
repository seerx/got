package session

//Session 接口
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Remove(key interface{}) error
	SID() string
}
